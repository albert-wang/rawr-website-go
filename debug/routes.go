package debug

import (
	"github.com/albert-wang/rawr-website-go/routes"

	"github.com/gorilla/mux"
)

func RegisterRoutes(router *mux.Router, ctx *routes.Context) {
	debugRouter := router.PathPrefix("/debug").Subrouter()

	debugRouter.HandleFunc("/hello", routes.Wrap(getHello, ctx)).
		Methods("GET")
}
