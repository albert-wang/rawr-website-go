package main

import (
	"github.com/garyburd/redigo/redis"
	"github.com/jmoiron/sqlx"
	"github.com/quipo/statsd"

	"time"
)

func collectStats(db *sqlx.DB, pool *redis.Pool, client *statsd.StatsdClient) {
	ticker := time.NewTicker(time.Second * 5)

	for _ = range ticker.C {
		stats := db.Stats()
		oc := stats.OpenConnections
		active := pool.ActiveCount()

		client.Gauge("dbconnections", int64(oc))
		client.Gauge("redisconnections", int64(active))
	}
}
