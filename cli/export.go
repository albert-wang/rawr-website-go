package cli

import (
	"os"
	"fmt"
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

	for i, v := range res {
		cleanTitle := v.Title
		cleanTitle = strings.ToLower(cleanTitle)
		cleanTitle = strings.Replace(cleanTitle, " ", "-", -1)

		fname := fmt.Sprintf("%s/%04d-%s.md", args[0], i, cleanTitle)

		file, err := os.Create(fname)
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
