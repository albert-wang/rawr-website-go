package debug

import (
	"github.com/albert-wang/rawr-website-go/routes"

	"github.com/garyburd/redigo/redis"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

func RegisterRoutes(router *mux.Router, db *sqlx.DB, pool *redis.Pool) {
	debugRouter := router.PathPrefix("/debug").Subrouter()
	debugRouter.HandleFunc("/hello", routes.Wrap(getHello, db, pool)).Methods("GET")

}
