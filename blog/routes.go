package blog

import (
	"github.com/albert-wang/rawr-website-go/routes"

	"github.com/gorilla/mux"
)

func RegisterRoutes(router *mux.Router, ctx *routes.Context) {
	router.HandleFunc("/", routes.Wrap(getIndex, ctx))
	router.HandleFunc("/post/{id:[0-9]+}/{slug}", routes.Wrap(getPost, ctx))
	router.HandleFunc("/post/{id:[0-9]+}", routes.Wrap(getPost, ctx))
	router.HandleFunc("/blog/{page:[0-9]+}", routes.Wrap(getGenericArchive, ctx))
	router.HandleFunc("/blog/{category:[a-z_]+}/{page:[0-9]+}", routes.Wrap(getCategoryArchive, ctx))

	//Non-blog related, but informational pages
	router.HandleFunc("/projects", routes.Wrap(getProjects, ctx))
	router.HandleFunc("/about", routes.Wrap(getAbout, ctx))
}
