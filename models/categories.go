package models

import (
	"github.com/albert-wang/tracederror"
	"github.com/jmoiron/sqlx"
)

type BlogCategory struct {
	ID       int32  `db:"id"`
	Category string `db:"category"`
}

func GetCategories(db sqlx.Ext) ([]BlogCategory, error) {
	q := `
		SELECT * FROM blog_categories ORDER BY id;
	`

	res := []BlogCategory{}
	err := sqlx.Select(db, &res, q)
	return res, tracederror.New(err)
}
