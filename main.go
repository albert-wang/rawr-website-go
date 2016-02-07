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
	"github.com/mitchellh/goamz/aws"
	"github.com/mitchellh/goamz/s3"
	"github.com/quipo/statsd"

	"github.com/albert-wang/rawr-website-go/admin"
	"github.com/albert-wang/rawr-website-go/blog"
	"github.com/albert-wang/rawr-website-go/cli"
	"github.com/albert-wang/rawr-website-go/config"
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

	cfg := config.Config{}
	err := config.LoadConfigurationFromFileAndEnvironment(os.Args[1], &cfg)
	if err != nil {
		log.Fatal("Could not load configuration file=", os.Args[1], " due to error=", err)
	}

	// Open up the DB and Redis connections.
	db, err := sqlx.Open("postgres", cfg.PostgresConnectionURL)
	if err != nil {
		log.Fatal("Could not open postgres connection with url=", cfg.PostgresConnectionURL, " error=", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("Could not ping postgres connection with url=", cfg.PostgresConnectionURL, " error=", err)
	}

	pool := createRedisPool(cfg.RedisHost, cfg.RedisPassword)
	conn := pool.Get()
	_, err = conn.Do("PING")
	if err != nil {
		log.Fatal("Could not ping redis DB with host=", cfg.RedisHost, " error=", err)
	}

	client := statsd.NewStatsdClient(cfg.StatsDHost, cfg.StatsDPrefix)
	err = client.CreateSocket()
	if err != nil {
		log.Fatal("Could not create statsd socket with host=", cfg.RedisHost, " error=", err)
	}

	auth, err := aws.SharedAuth()
	if err != nil {
		log.Fatal("Could not load aws authentication")
	}

	ctx := routes.CreateContext(db, pool, auth, &cfg)

	//Check for arguments.
	if len(os.Args) > 2 {
		cli.Dispatch(os.Args[2:], ctx)
		return
	}

	router := mux.NewRouter()
	router.NotFoundHandler = http.HandlerFunc(routes.NotFound)
	router.HandleFunc("/favicon{suffix}", serveFavicon)
	router.HandleFunc("/mstile{suffix}", serveFavicon)

	blog.RegisterRoutes(router, ctx)
	admin.RegisterRoutes(router, ctx)
	if cfg.Debug {
		debug.RegisterRoutes(router, ctx)
	}

	go collectStats(db, pool, client)

	listeningAddress := fmt.Sprintf("localhost:%d", cfg.Port)

	log.Print("Listening on addr=", listeningAddress)
	http.ListenAndServe(listeningAddress, router)
}
