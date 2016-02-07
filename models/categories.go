package models

import (
	"github.com/albert-wang/tracederror"
	"github.com/jmoiron/sqlx"
)

type BlogCategory struct {
	ID       int32  `db:"id"`
	Category string `db:"category"`
	Hidden   bool   `db:"hidden"`
}

func GetCategories(db sqlx.Ext) ([]BlogCategory, error) {
	q := `
		SELECT * FROM blog_categories ORDER BY id;
	`

	res := []BlogCategory{}
	err := sqlx.Select(db, &res, q)
	return res, tracederror.New(err)
}

func IsValidCategory(db sqlx.Ext, category string) (bool, error) {
	q := `SELECT * FROM blog_categories WHERE category = $1`

	res := []BlogCategory{}
	err := sqlx.Select(db, &res, q, category)

	if err != nil {
		return false, tracederror.New(err)
	}

	return len(res) == 1, nil
}
