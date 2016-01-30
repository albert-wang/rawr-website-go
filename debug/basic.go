package debug

import (
	"net/http"

	"github.com/albert-wang/rawr-website-go/routes"
)

func getHello(w http.ResponseWriter, r *http.Request, ctx routes.Context) error {
	tpl, err := routes.LoadTemplates("echo.tpl")
	if err != nil {
		return routes.MakeHttpError(err, http.StatusInternalServerError, r)
	}

	return routes.RenderTemplateWithData(w, r, tpl, nil)
}
