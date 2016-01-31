package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/garyburd/redigo/redis"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/quipo/statsd"

	"github.com/albert-wang/rawr-website-go/blog"
	"github.com/albert-wang/rawr-website-go/debug"
	"github.com/albert-wang/rawr-website-go/routes"
)

func createRedisPool(host, pass string) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", host)
			if err != nil {
				return nil, err
			}

			if len(pass) != 0 {
				if _, err := c.Do("AUTH", pass); err != nil {
					c.Close()
					return nil, err
				}
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}

func serveFavicon(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static"+r.RequestURI)
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	config := Config{}
	err := LoadConfigurationFromFileAndEnvironment(os.Args[1], &config)
	if err != nil {
		log.Fatal("Could not load configuration file=", os.Args[1], " due to error=", err)
	}

	// Open up the DB and Redis connections.
	db, err := sqlx.Open("postgres", config.PostgresConnectionURL)
	if err != nil {
		log.Fatal("Could not open postgres connection with url=", config.PostgresConnectionURL, " error=", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("Could not ping postgres connection with url=", config.PostgresConnectionURL, " error=", err)
	}

	pool := createRedisPool(config.RedisHost, config.RedisPassword)
	conn := pool.Get()
	_, err = conn.Do("PING")
	if err != nil {
		log.Fatal("Could not ping redis DB with host=", config.RedisHost, " error=", err)
	}

	client := statsd.NewStatsdClient(config.StatsDHost, config.StatsDPrefix)
	err = client.CreateSocket()
	if err != nil {
		log.Fatal("Could not create statsd socket with host=", config.RedisHost, " error=", err)
	}

	router := mux.NewRouter()
	router.NotFoundHandler = http.HandlerFunc(routes.NotFound)
	router.HandleFunc("/favicon{suffix}", serveFavicon)
	router.HandleFunc("/mstile{suffix}", serveFavicon)

	blog.RegisterRoutes(router, db, pool)
	if config.Debug {
		debug.RegisterRoutes(router, db, pool)
	}

	go collectStats(db, pool, client)

	listeningAddress := fmt.Sprintf("localhost:%d", config.Port)

	log.Print("Listening on addr=", listeningAddress)
	http.ListenAndServe(listeningAddress, router)
}
