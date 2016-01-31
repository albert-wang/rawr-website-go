package blog

import (
	"github.com/garyburd/redigo/redis"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

func RegisterRoutes(router *mux.Router, db *sqlx.DB, pool *redis.Pool) {

}
