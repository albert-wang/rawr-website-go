package routes

import (
	"net/http"

	"github.com/albert-wang/tracederror"
	"github.com/garyburd/redigo/redis"
	"github.com/jmoiron/sqlx"
)

type Context struct {
	Pool *redis.Pool
	DB   *sqlx.DB
}

func Wrap(fn func(http.ResponseWriter, *http.Request, Context) error, DB *sqlx.DB, p *redis.Pool) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		err := fn(w, r, Context{
			Pool: p,
			DB:   DB,
		})

		if err != nil {
			inner := tracederror.Inner(err)
			maybeHttpErr, ok := inner.(*HttpError)
			if ok {
				//Render the appropriate error page.
				RenderErrorPage(w, r, maybeHttpErr)
			}
		}
	}
}
