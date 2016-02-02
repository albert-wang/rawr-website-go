package routes

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/albert-wang/tracederror"
	"github.com/russross/blackfriday"
)

func templateBlackfriday(mkd string) template.HTML {
	res := blackfriday.MarkdownCommon([]byte(mkd))
	return template.HTML(res)
}

func s3img(path string) string {
	return "//s3.amazonaws.com/img.rawrrawr.com/" + path
}

func add(a, b int) int {
	return a + b
}

func slug(s string) string {
	s = strings.TrimSpace(s)
	s = strings.ToLower(s)
	fields := strings.Fields(s)

	res := strings.Join(fields, "-")
	return url.QueryEscape(res)
}

func timef(f string) string {
	return time.Now().Format(f)
}

func LoadTemplates(str ...string) (*template.Template, error) {
	pathed := []string{}
	for _, v := range str {
		pathed = append(pathed, fmt.Sprintf("assets/templates/%s", v))
	}

	res := template.New(str[0])
	// Load default functions.
	mapping := template.FuncMap{
		"blackfriday": templateBlackfriday,
		"s3img":       s3img,
		"add":         add,
		"slug":        slug,
		"timef":       timef,
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
