package cli

import (
	"os"
	"strings"
	"text/template"

	"github.com/albert-wang/rawr-website-go/models"
	"github.com/albert-wang/rawr-website-go/routes"
)

func exportPosts(args []string, context *routes.Context) error {
	res := []models.BlogPost{}
	err := context.DB.Select(&res, "SELECT * FROM blog_posts")
	if err != nil {
		return err
	}

	tpl := template.New("assets/templates/export.tpl")
	tpl.ParseFiles("assets/templates/export.tpl")
	for _, v := range res {
		cleanTitle := v.Title
		cleanTitle = strings.ToLower(cleanTitle)
		cleanTitle = strings.Replace(cleanTitle, " ", "-", -1)

		file, err := os.Create(args[0] + "/" + cleanTitle + ".md")
		if err != nil {
			return err
		}

		err = tpl.Execute(file, v)
		if err != nil {
			return err
		}

		file.Close()
	}

	return nil
}
