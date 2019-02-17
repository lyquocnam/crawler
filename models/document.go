package models

import (
	"github.com/PuerkitoBio/goquery"
	"time"
)

type Document struct {
	Title        string   `csv:"title"`
	URL          string   `csv:"url"`
	Author       string   `csv:"author"`
	Date         DateTime `csv:"date"`
	Content      *goquery.Document
	RelatedLinks []string
}

type Fetcher interface {
	Fetch(url string) error
	FindRelated() error
	ParseDateTime(date string) (*time.Time, error)
	ExportCSV() error
	Print()
}
