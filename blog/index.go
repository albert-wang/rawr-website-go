package blog

import (
	"net/http"

	"github.com/albert-wang/tracederror"

	"github.com/albert-wang/rawr-website-go/routes"
)

func getIndex(w http.ResponseWriter, r *http.Request, ctx routes.Context) error {
	tpl, err := routes.LoadTemplates("base.tpl", "index.tpl")
	if err != nil {
		return tracederror.New(err)
	}

	return routes.RenderTemplateWithData(w, r, tpl, nil)
}
