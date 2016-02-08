package routes

import (
	"log"
	"net/http"

	"github.com/albert-wang/tracederror"
	"github.com/garyburd/redigo/redis"
	"github.com/jmoiron/sqlx"
	"github.com/mitchellh/goamz/aws"
	"github.com/mitchellh/goamz/s3"

	"github.com/albert-wang/rawr-website-go/config"
)

type Context struct {
	Pool   *redis.Pool
	DB     *sqlx.DB
	Config *config.Config

	// AWS configuration
	Auth   aws.Auth
	S3     *s3.S3
	Bucket *s3.Bucket
}

func CreateContext(DB *sqlx.DB, p *redis.Pool, auth aws.Auth, cfg *config.Config) *Context {
	s3 := s3.New(auth, aws.USEast)
	bucket := s3.Bucket(cfg.GalleryBucket)

	return &Context{
		Pool:   p,
		DB:     DB,
		Config: cfg,
		Auth:   auth,
		S3:     s3,
		Bucket: bucket,
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
