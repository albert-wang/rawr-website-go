package routes

import (
	"log"
	"net/http"

	"github.com/albert-wang/tracederror"
	"github.com/garyburd/redigo/redis"
	"github.com/jmoiron/sqlx"

	"github.com/albert-wang/rawr-website-go/config"
)

type Context struct {
	Pool   *redis.Pool
	DB     *sqlx.DB
	Config *config.Config
}

func CreateContext(DB *sqlx.DB, p *redis.Pool, cfg *config.Config) *Context {
	return &Context{
		Pool:   p,
		DB:     DB,
		Config: cfg,
	}
}

func Wrap(fn func(http.ResponseWriter, *http.Request, Context) error, ctx *Context) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		err := fn(w, r, *ctx)

		if err != nil {
			log.Print(err)

			inner := tracederror.Inner(err)
			maybeHttpErr, ok := inner.(*HttpError)
			if ok {
				//Render the appropriate error page.
				RenderErrorPage(w, r, maybeHttpErr)
			}
		}
	}
}
