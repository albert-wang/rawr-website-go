package cli

import (
	"os"
	"log"
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

	tpl := template.New("export.tpl")
	tpl, err = tpl.ParseFiles("assets/templates/export.tpl")
	if err != nil {
		log.Print(err)
	}

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
