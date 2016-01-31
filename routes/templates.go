package routes

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"

	"github.com/albert-wang/tracederror"
	"github.com/russross/blackfriday"
)

func templateBlackfriday(mkd string) template.HTML {
	res := blackfriday.MarkdownCommon([]byte(mkd))
	return template.HTML(res)
}

func LoadTemplates(str ...string) (*template.Template, error) {
	pathed := []string{}
	for _, v := range str {
		pathed = append(pathed, fmt.Sprintf("data/templates/%s", v))
	}

	res := template.New(str[0])
	// Load default functions.
	mapping := template.FuncMap{
		"blackfriday": templateBlackfriday,
	}

	res = res.Funcs(mapping)

	res, err := res.ParseFiles(pathed...)
	if err != nil {
		log.Print(tracederror.New(err))
		return nil, tracederror.New(err)
	}

	return res, tracederror.New(err)
}

func RenderTemplateWithData(w io.Writer, r *http.Request, tpl *template.Template, data interface{}) error {
	err := tpl.Execute(w, data)
	if err != nil {
		return MakeHttpError(err, http.StatusInternalServerError, r)
	}

	return nil
}
