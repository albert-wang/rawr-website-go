package routes

import (
	"fmt"
	"log"
	"net/http"

	"github.com/albert-wang/tracederror"
)

type HttpError struct {
	Err        error
	StatusCode int
	Request    *http.Request
}

func (s *HttpError) Error() string {
	return fmt.Sprintf("%d - %s", s.StatusCode, s.Err)
}

func MakeHttpError(err error, status int, req *http.Request) error {
	res := &HttpError{
		Err:        err,
		StatusCode: status,
		Request:    req,
	}

	return tracederror.NewWithContext(res, req)
}

func RenderErrorPage(w http.ResponseWriter, r *http.Request, initialError *HttpError) {
	tpl, err := LoadTemplates("error.tpl")
	if err != nil {
		//Stuffs super screwed
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = RenderTemplateWithData(w, r, tpl, initialError)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func NotFound(w http.ResponseWriter, r *http.Request) {
	log.Print("Tried to load path=", r.RequestURI, " not found")

	tpl, err := LoadTemplates("404.tpl")
	if err != nil {
		//Stuffs super screwed
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	err = RenderTemplateWithData(w, r, tpl, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
}
