package admin

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"

	"github.com/albert-wang/rawr-website-go/models"
	"github.com/albert-wang/rawr-website-go/routes"
	"github.com/albert-wang/tracederror"
)

type editPage struct {
	Post       *models.BlogPost
	Categories []models.BlogCategory
}

func getEdit(w http.ResponseWriter, r *http.Request, ctx routes.Context) error {
	err := routes.CheckAuth(r, ctx)
	if err != nil {
		return err
	}

	vars := mux.Vars(r)
	id := int64(0)
	if vars["id"] != "new" {
		id, err = strconv.ParseInt(vars["id"], 10, 64)
		if err != nil {
			return tracederror.New(err)
		}
	}

	post := &models.BlogPost{}

	if id != 0 {
		post, err = models.GetBlogPostByID(ctx.DB, int32(id))
		if err != nil {
			return routes.MakeHttpError(err, http.StatusNotFound, r)
		}
	}

	categories, err := models.GetCategories(ctx.DB)
	if err != nil {
		return tracederror.New(err)
	}

	tpl, err := routes.LoadTemplates("base.tpl", "edit.tpl")
	if err != nil {
		return tracederror.New(err)
	}

	data := editPage{
		Post:       post,
		Categories: categories,
	}

	return routes.RenderTemplateWithData(w, r, tpl, &data)
}

type editBody struct {
	ID       int32
	Category int32
	Title    string
	Content  string
	Publish  int32
	Hero     string
}

type editResponse struct {
	ID int32
}

func postEdit(w http.ResponseWriter, r *http.Request, ctx routes.Context) error {
	err := routes.CheckAuth(r, ctx)
	if err != nil {
		return err
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return routes.MakeHttpError(err, http.StatusBadRequest, r)
	}

	res := editBody{}
	err = json.Unmarshal(body, &res)
	if err != nil {
		return routes.MakeHttpError(err, http.StatusBadRequest, r)
	}

	post := &models.BlogPost{
		ID:         res.ID,
		CategoryID: 1,
	}

	if res.ID != 0 {
		post, err = models.GetBlogPostByID(ctx.DB, res.ID)
		if err != nil {
			return tracederror.New(err)
		}
	}

	post.Content = res.Content
	post.Title = res.Title
	post.CategoryID = res.Category
	post.Hero = res.Hero

	t := time.Unix(int64(res.Publish), 0)
	post.Publish = &t

	if res.Publish == 0 {
		post.Publish = nil
	}

	err = post.Save(ctx.DB)
	if err != nil {
		return tracederror.New(err)
	}

	result := editResponse{
		ID: post.ID,
	}

	bytes, err := json.Marshal(result)
	if err != nil {
		return tracederror.New(err)
	}

	w.Write(bytes)
	return nil
}

func postRender(w http.ResponseWriter, r *http.Request, ctx routes.Context) error {
	err := routes.CheckAuth(r, ctx)
	if err != nil {
		return err
	}

	tpl, err := routes.LoadTemplates("render.tpl")
	if err != nil {
		return tracederror.New(err)
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return tracederror.New(err)
	}

	bodyStr := string(body)
	bodyStr = strings.Replace(bodyStr, "{:truncate}", "", 1)
	bodyStr = strings.Replace(bodyStr, "{:longtruncate}", "", 1)

	return routes.RenderTemplateWithData(w, r, tpl, bodyStr)
}
