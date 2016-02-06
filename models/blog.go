package models

import (
	"strings"
	"time"

	"github.com/albert-wang/tracederror"
	"github.com/jmoiron/sqlx"
)

type BlogPost struct {
	ID         int32      `db:"id"`
	CategoryID int32      `db:"category"`
	Title      string     `db:"title"`
	Content    string     `db:"content"`
	Publish    *time.Time `db:"publish"`
	Hero       string     `db:"hero"`

	Category string `db:"category_value"`
}

func clean(v string) string {
	v = strings.Replace(v, "{:truncate}", "", -1)
	v = strings.Replace(v, "{:longtruncate}", "", -1)
	return v
}

func (b *BlogPost) Short() string {
	ind := strings.Index(b.Content, "{:truncate}")
	if ind > 0 {
		return clean(b.Content[0:ind])
	}

	return b.Full()
}

func (b *BlogPost) Long() string {
	ind := strings.Index(b.Content, "{:longtruncate}")
	if ind > 0 {
		return b.Content[0:ind]
	}
	return b.Full()
}

func (b *BlogPost) Full() string {
	return clean(b.Content)
}

func (b *BlogPost) Save(db sqlx.Ext) error {
	if b.ID == 0 {
		q := `INSERT INTO blog_posts (category, title, content, publish, hero) 
		                      VALUES(:category,:title,:content,:publish,:hero)
		      RETURNING id
		`

		rows, err := sqlx.NamedQuery(db, q, b)
		if err != nil {
			return err
		}

		defer rows.Close()
		rows.Next()
		err = rows.Scan(&b.ID)
		return err
	} else {
		q := `UPDATE blog_posts SET 
			category = :category, title = :title, content = :content, publish = :publish,
			hero = :hero
			WHERE id = :id
		`
		_, err := sqlx.NamedExec(db, q, b)
		return err
	}
}

func GetBlogPostByID(db sqlx.Ext, id int32) (*BlogPost, error) {
	q := `
		SELECT p.*, c.category AS category_value 
		FROM blog_posts p JOIN blog_categories c ON c.id = p.category
		WHERE p.id = $1
	`

	res := BlogPost{}
	err := sqlx.Get(db, &res, q, id)
	return &res, tracederror.New(err)
}

func GetRecentBlogPosts(db sqlx.Ext, count int32) ([]BlogPost, error) {
	q := `
		SELECT p.*, c.category AS category_value 
		FROM blog_posts p JOIN blog_categories c ON c.id = p.category
		WHERE publish IS NOT NULL
		ORDER BY publish DESC
		LIMIT $1
	`

	res := []BlogPost{}
	err := sqlx.Select(db, &res, q, count)
	return res, tracederror.New(err)
}

func GetRecentPostsInCategory(db sqlx.Ext, category string, count int32) ([]BlogPost, error) {
	q := `
		SELECT p.*, c.category AS category_value 
		FROM blog_posts p JOIN blog_categories c ON c.id = p.category
		WHERE publish IS NOT NULL AND c.category = $1
		ORDER BY publish DESC
		LIMIT $2
	`

	res := []BlogPost{}
	err := sqlx.Select(db, &res, q, category, count)
	return res, tracederror.New(err)
}
