package tests

import (
	"crawler/models"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var url = "https://dulich.tuoitre.vn/nguoi-nhat-du-bao-hoa-anh-dao-no-bang-cach-nao-20190212104808155.htm"

func TestTuoiTreFetch(t *testing.T) {
	doc := models.TuoiTreDocument{}
	err := doc.Fetch(url)
	if err != nil {
		t.Error(err)
	}
}

func TestTuoiTreSaveFile(t *testing.T) {
	doc := models.TuoiTreDocument{}
	err := doc.Fetch(url)
	if err != nil {
		t.Error(err)
	}

	err = doc.SaveCSV()
	if err != nil {
		t.Error(err)
	}
}

func TestTuoiTreParseDateTime(t *testing.T) {
	doc := models.TuoiTreDocument{}
	d, err := doc.ParseDateTime("13/02/2019 10:18 GMT+7")
	if err != nil {
		t.Error(err)
		return
	}

	if d.IsZero() {
		t.Errorf("time must be not zero value")
	}

	assert.Equal(t, 13, d.Day())
	assert.Equal(t, time.Month(2), d.Month())
	assert.Equal(t, 2019, d.Year())
	assert.Equal(t, 10, d.Hour())
	assert.Equal(t, 18, d.Minute())
}
