package models

import (
	"time"
)

type Document struct {
	Title  string   `csv:"title"`
	URL    string   `csv:"url"`
	Author string   `csv:"author"`
	Date   DateTime `csv:"date"`
}

type Fetcher interface {
	Fetch(url string) error
	ParseDateTime(date string) (*time.Time, error)
	ExportCSV() error
	Print()
}
