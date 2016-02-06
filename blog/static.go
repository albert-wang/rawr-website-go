package blog

import (
	"net/http"

	"github.com/albert-wang/rawr-website-go/routes"
)

func getAbout(w http.ResponseWriter, r *http.Request, ctx routes.Context) error {
	tpl, err := routes.LoadTemplates("base.tpl", "about.tpl")
	if err != nil {
		return routes.MakeHttpError(err, http.StatusNotFound, r)
	}

	return routes.RenderTemplateWithData(w, r, tpl, nil)
}

func getProjects(w http.ResponseWriter, r *http.Request, ctx routes.Context) error {
	tpl, err := routes.LoadTemplates("base.tpl", "projects.tpl")
	if err != nil {
		return routes.MakeHttpError(err, http.StatusNotFound, r)
	}

	return routes.RenderTemplateWithData(w, r, tpl, nil)
}
