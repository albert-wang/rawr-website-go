package models

import (
	"fmt"
	"log"
	"sort"
	"strings"
	"time"

	"github.com/albert-wang/tracederror"
	"github.com/jmoiron/sqlx"

	"github.com/mitchellh/goamz/s3"
)

type Gallery struct {
	ID          int    `db:"id"`
	Name        string `db:"name"`
	Description string `db:"description"`
	Displayed   bool   `db:"displayed"`
	Hero        string `db:"hero"`
	S3Prefix    string `db:"s3prefix"`
}

func GetGalleries(db sqlx.Ext) ([]Gallery, error) {
	q := `SELECT * FROM galleries WHERE displayed`

	res := []Gallery{}
	err := sqlx.Select(db, &res, q)
	return res, tracederror.New(err)
}

func GetHiddenGalleries(db sqlx.Ext) ([]Gallery, error) {
	q := `SELECT * FROM galleries`

	res := []Gallery{}
	err := sqlx.Select(db, &res, q)
	return res, tracederror.New(err)
}

func GetGalleryByID(db sqlx.Ext, id int32) (*Gallery, error) {
	q := `SELECT * FROM galleries WHERE id=$1`

	res := Gallery{}
	err := sqlx.Get(db, &res, q, id)

	return &res, tracederror.New(err)
}

type Image struct {
	Thumb string
	Orig  string
	Hero  string

	LastModified time.Time
	ETag         string
}

type ByDate []Image

func (r ByDate) Len() int {
	return len(r)
}

func (r ByDate) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}

func (r ByDate) Less(i, j int) bool {
	return r[i].LastModified.Before(r[j].LastModified)
}

func GetImagesInGallery(bucket *s3.Bucket, g *Gallery) ([]Image, error) {
	pre := fmt.Sprintf("%s%s", g.S3Prefix, "orig-")
	res, err := bucket.List(pre, "/", "", 1000)
	if err != nil {
		return nil, tracederror.New(err)
	}

	paths := []Image{}
	for _, v := range res.Contents {
		if v.Key[len(v.Key)-1] != '/' {
			raw := strings.TrimPrefix(v.Key, pre)

			t, err := time.Parse("2006-01-02T15:04:05.000Z", v.LastModified)
			if err != nil {
				log.Print("Couldn't parse time from amazon, assuming 'now' for upload date.")
				log.Print(err)

				t = time.Now()
			}

			paths = append(paths, Image{
				Thumb: fmt.Sprintf("%sthumb-%s", g.S3Prefix, raw),
				Orig:  fmt.Sprintf("%sorig-%s", g.S3Prefix, raw),
				Hero:  fmt.Sprintf("%shero-%s", g.S3Prefix, raw),

				ETag:         v.ETag,
				LastModified: t,
			})
		}
	}

	sort.Sort(ByDate(paths))
	return paths, nil
}
