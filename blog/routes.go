package blog

import (
	"github.com/albert-wang/rawr-website-go/routes"

	"github.com/garyburd/redigo/redis"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

func RegisterRoutes(router *mux.Router, db *sqlx.DB, pool *redis.Pool) {
	router.HandleFunc("/", routes.Wrap(getIndex, db, pool))
}
