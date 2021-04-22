package repositories

import (
	"go-api/data"
	"go-api/models"
	"reflect"
	"testing"
)

func TestCardBacksRepository(t *testing.T) {
	mockCB := models.CardBack{ID: 0, Slug: "Mock"}
	source := data.FileSource{
		Data: &[]models.CardBack{mockCB},
	}
	cbr := NewCardBackRepository(&source)

	t.Run("should get all cardbacks", func(t *testing.T) {
		expected := source.Data
		actual, _ := cbr.GetAll()
		if actual != expected {
			t.Errorf("actual %v, expected %v", actual, expected)
		}
	})

	t.Run("should add new cardback to repository", func(t *testing.T) {
		mock := models.CardBack{ID: 1, Slug: "Mock2"}
		expected := &[]models.CardBack{mockCB, mock}
		cbr.Create(&mock)
		actual, _ := cbr.GetAll()
		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("actual %v, expected %v", actual, expected)
		}
	})

	t.Run("should get cardback by ID", func(t *testing.T) {
		expected := mockCB
		actual, _ := cbr.FindByID(0)
		if *actual != expected {
			t.Errorf("actual %v, expected %v", actual, expected)
		}
	})

	t.Run("should return true if cardback exists", func(t *testing.T) {
		expected := true
		actual := cbr.Exists("Mock")
		if actual != expected {
			t.Errorf("actual %v, expected %v", actual, expected)
		}
	})

	t.Run("should return false if cardback doesn't exists", func(t *testing.T) {
		expected := false
		actual := cbr.Exists("Other")
		if actual != expected {
			t.Errorf("actual %v, expected %v", actual, expected)
		}
	})
}
