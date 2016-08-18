package admin

import (
	"github.com/albert-wang/rawr-website-go/routes"

	"github.com/gorilla/mux"
)

func RegisterRoutes(router *mux.Router, ctx *routes.Context) {
	adminRouter := router.PathPrefix("/admin").Subrouter()

	adminRouter.HandleFunc("/edit/{id}", routes.Wrap(getEdit, ctx)).
		Methods("GET")

	adminRouter.HandleFunc("/edit", routes.Wrap(postEdit, ctx)).
		Methods("POST")

	adminRouter.HandleFunc("/render", routes.Wrap(postRender, ctx)).
		Methods("POST")

	adminRouter.HandleFunc("/gallery_reset/{gallery}", routes.Wrap(getGalleryReset, ctx)).
		Methods("GET")
}
