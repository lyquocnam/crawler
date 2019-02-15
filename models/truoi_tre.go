package models

import (
	"errors"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/gocarina/gocsv"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

type TuoiTreDocument struct {
	Document
}

func (doc *TuoiTreDocument) Fetch(url string) error {
	if len(url) < 1 {
		return errors.New("Bạn chưa nhập địa chỉ trang web (URL)")
	}

	resp, err := http.Get("https://dulich.tuoitre.vn/ngam-tuyet-trang-anh-dao-o-nhat-ban-thang-2-20190212121811143.htm")
	if err != nil {
		return err
	}

	html, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return err
	}

	title := strings.TrimSpace(html.Find("title").First().Text())
	author := strings.TrimSpace(html.Find(".author").First().Text())
	dateString := strings.TrimSpace(html.Find(".date-time").First().Text())
	t, err := doc.ParseDateTime(dateString)
	if err != nil {
		return err
	}

	doc.Title = title
	doc.Author = author
	doc.Date = DateTime{Time: *t}
	doc.URL = url

	return nil
}

func (doc *TuoiTreDocument) SaveCSV() error {
	filename := "data.csv"
	var file *os.File
	data := make([]TuoiTreDocument, 0)
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		file, err = os.Create(filename)
	} else {
		file, err = os.OpenFile(filename, os.O_APPEND|os.O_RDWR, os.ModeAppend)
		if err != nil {
			return err
		}
		err = gocsv.UnmarshalFile(file, &data)
		if err != nil {
			return err
		}
	}
	defer file.Close()

	data = append(data, *doc)
	err := gocsv.MarshalWithoutHeaders(data, file)
	if err != nil {
		return err
	}

	return nil
}

func (doc *TuoiTreDocument) ParseDateTime(d string) (*time.Time, error) {
	if len(d) < 1 {
		return nil, errors.New("DateTime string không được rỗng")
	}

	s1 := strings.Split(d, " ")
	if len(s1) < 2 {
		return nil, errors.New("DateTime string không đúng định dạng")
	}

	dateArr := strings.Split(s1[0], "/")
	if len(dateArr) != 3 {
		return nil, errors.New("DateTime string không đúng định dạng")
	}

	day, err := strconv.Atoi(dateArr[0])
	if err != nil {
		return nil, errors.New("DateTime string không đúng định dạng")
	}

	month, err := strconv.Atoi(dateArr[1])
	if err != nil {
		return nil, errors.New("DateTime string không đúng định dạng")
	}
	year, err := strconv.Atoi(dateArr[2])
	if err != nil {
		return nil, errors.New("DateTime string không đúng định dạng")
	}

	hourArr := strings.Split(s1[1], ":")
	if len(hourArr) < 2 {
		return nil, errors.New("DateTime string không đúng định dạng")
	}

	hour, err := strconv.Atoi(hourArr[0])
	if err != nil {
		return nil, errors.New("DateTime string không đúng định dạng")
	}

	minute, err := strconv.Atoi(hourArr[1])
	if err != nil {
		return nil, errors.New("DateTime string không đúng định dạng")
	}

	result := time.Date(year, time.Month(month), day, hour, minute, 0, 0, time.FixedZone("", 0))

	return &result, nil
}

func (doc *TuoiTreDocument) Print() {
	fmt.Printf(""+
		"title: %s\n"+
		"author: %s\n"+
		"date: %s\n"+
		"url: %s\n", doc.Title, doc.Author, doc.Date.Format("02/01/06 03:04"), doc.URL)
}
