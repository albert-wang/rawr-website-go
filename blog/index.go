package blog

import (
	"fmt"
	"net/http"

	"github.com/albert-wang/tracederror"

	"github.com/albert-wang/rawr-website-go/models"
	"github.com/albert-wang/rawr-website-go/routes"
)

type indexPage struct {
	Posts    []models.BlogPost
	Featured models.BlogPost
}

func getIndex(w http.ResponseWriter, r *http.Request, ctx routes.Context) error {
	tpl, err := routes.LoadTemplates("base.tpl", "index.tpl")
	if err != nil {
		return tracederror.New(err)
	}

	recent, err := models.GetRecentBlogPosts(ctx.DB, 4, 0)
	if err != nil {
		return tracederror.New(err)
	}

	featured, err := models.GetRecentPostsInCategory(ctx.DB, "featured", 1, 0)
	if err != nil {
		return tracederror.New(err)
	}

	if len(featured) == 0 {
		return tracederror.New(fmt.Errorf("No featured posts?"))
	}

	data := indexPage{
		Posts:    recent,
		Featured: featured[0],
	}

	return routes.RenderTemplateWithData(w, r, tpl, &data)
}
