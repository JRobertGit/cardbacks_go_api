package data

import (
	"bytes"
	"go-api/models"
	"reflect"
	"testing"
)

var r = bytes.NewBufferString("id,sortCategory,text,name,image,slug\n0,0,Text,Name,Image,Slug\n")

var cb = models.CardBack{
	ID:           0,
	SortCategory: 0,
	Text:         "Text",
	Name:         "Name",
	Image:        "Image",
	Slug:         "Slug",
}

var data = []models.CardBack{cb}

func TestWriteCSVData(t *testing.T) {
	t.Run("should write data in cvs format", func(t *testing.T) {
		expected := "id,sortCategory,text,name,image,slug\n0,0,Text,Name,Image,Slug\n"
		var b bytes.Buffer

		writeCSVData(&data, &b)
		actual := b.String()

		if actual != expected {
			t.Errorf("actual %v, expected %v", actual, expected)
		}
	})
}

func TestReadCSVData(t *testing.T) {
	t.Run("should read data from csv format", func(t *testing.T) {
		expected := data
		actual, _ := readCSVData(r)

		if reflect.DeepEqual(actual, expected) {
			t.Errorf("actual %v, expected %v", actual, expected)
		}
	})
}
