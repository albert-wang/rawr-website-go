package routes

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
)

func LoadTemplates(str ...string) (*template.Template, error) {
	pathed := []string{}
	for _, v := range str {
		pathed = append(pathed, fmt.Sprintf("data/templates/%s", v))
	}

	return template.ParseFiles(pathed...)
}

func RenderTemplateWithData(w io.Writer, r *http.Request, tpl *template.Template, data interface{}) error {
	err := tpl.Execute(w, data)
	if err != nil {
		return MakeHttpError(err, http.StatusInternalServerError, r)
	}

	return nil
}
