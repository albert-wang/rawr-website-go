package blog

import (
	"net/http"
	"strconv"
	"time"

	"github.com/albert-wang/tracederror"
	"github.com/gorilla/mux"

	"github.com/albert-wang/rawr-website-go/models"
	"github.com/albert-wang/rawr-website-go/routes"
)

type postPage struct {
	Post   *models.BlogPost
	Recent []models.BlogPost
}

func getPost(w http.ResponseWriter, r *http.Request, ctx routes.Context) error {
	vars := mux.Vars(r)

	tpl, err := routes.LoadTemplates("base.tpl", "post.tpl")
	if err != nil {
		return tracederror.New(err)
	}

	recent, err := models.GetRecentBlogPosts(ctx.DB, 4)
	if err != nil {
		return tracederror.New(err)
	}

	id, err := strconv.ParseInt(vars["id"], 10, 32)
	if err != nil {
		return tracederror.New(err)
	}

	self, err := models.GetBlogPostByID(ctx.DB, int32(id))
	if err != nil {
		return tracederror.New(err)
	}

	authErr := routes.CheckAuth(r, ctx)
	if self.Publish == nil && authErr != nil {
		return routes.MakeHttpError(nil, http.StatusNotFound, r)
	}

	if self.Publish != nil && self.Publish.After(time.Now().UTC()) && authErr != nil {
		return routes.MakeHttpError(nil, http.StatusNotFound, r)
	}

	data := postPage{
		Post:   self,
		Recent: recent,
	}

	return routes.RenderTemplateWithData(w, r, tpl, &data)
}
