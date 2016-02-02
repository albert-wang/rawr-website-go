package blog

import (
	"github.com/albert-wang/rawr-website-go/routes"

	"github.com/gorilla/mux"
)

func RegisterRoutes(router *mux.Router, ctx *routes.Context) {
	router.HandleFunc("/", routes.Wrap(getIndex, ctx))
	router.HandleFunc("/post/{id}/{slug}", routes.Wrap(getPost, ctx))
	//router.HandleFunc("/recent/{page}", routes.Wrap(getRecent, db, pool))
}
