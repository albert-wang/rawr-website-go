package blog

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/albert-wang/tracederror"
	"github.com/gorilla/mux"

	"github.com/albert-wang/rawr-website-go/models"
	"github.com/albert-wang/rawr-website-go/routes"
)

type archivePage struct {
	Page     int32
	Posts    []models.BlogPost
	Category string
}

func getGenericArchive(w http.ResponseWriter, r *http.Request, ctx routes.Context) error {
	vars := mux.Vars(r)

	page, err := strconv.ParseInt(vars["page"], 10, 32)
	if err != nil {
		return tracederror.New(err)
	}

	if page <= 0 {
		return routes.MakeHttpError(fmt.Errorf("Page must be larger than 0"), http.StatusBadRequest, r)
	}

	tpl, err := routes.LoadTemplates("base.tpl", "archive.tpl")
	if err != nil {
		return routes.MakeHttpError(err, http.StatusNotFound, r)
	}

	posts, err := models.GetRecentBlogPosts(ctx.DB, 100, int32((page-1)*100))
	if err != nil {
		return tracederror.New(err)
	}

	data := archivePage{
		Page:     int32(page),
		Posts:    posts,
		Category: "All",
	}

	return routes.RenderTemplateWithData(w, r, tpl, data)
}

func getCategoryArchive(w http.ResponseWriter, r *http.Request, ctx routes.Context) error {
	vars := mux.Vars(r)

	page, err := strconv.ParseInt(vars["page"], 10, 32)
	if err != nil {
		return tracederror.New(err)
	}

	if page <= 0 {
		return routes.MakeHttpError(fmt.Errorf("Page must be larger than 0"), http.StatusBadRequest, r)
	}

	category := vars["category"]
	valid, err := models.IsValidCategory(ctx.DB, category)

	if err != nil {
		return tracederror.New(err)
	}

	if !valid {
		return routes.MakeHttpError(fmt.Errorf("Invalid category %s", category), http.StatusBadRequest, r)
	}

	tpl, err := routes.LoadTemplates("base.tpl", "archive.tpl")
	if err != nil {
		return routes.MakeHttpError(err, http.StatusNotFound, r)
	}

	posts, err := models.GetRecentPostsInCategory(ctx.DB, category, 100, int32((page-1)*100))
	if err != nil {
		return tracederror.New(err)
	}

	data := archivePage{
		Page:     int32(page),
		Posts:    posts,
		Category: strings.Title(category),
	}

	return routes.RenderTemplateWithData(w, r, tpl, data)
}
