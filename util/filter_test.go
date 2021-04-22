package util

import (
	"fmt"
	"go-api/models"
	"reflect"
	"sync"
	"testing"
)

var collection = []models.CardBack{
	{ID: 0, Slug: fmt.Sprintf("item: %v", 0)},
	{ID: 1, Slug: fmt.Sprintf("item: %v", 1)},
	{ID: 2, Slug: fmt.Sprintf("item: %v", 2)},
}

func TestFilterDataAsync(t *testing.T) {
	limit := -1
	workers := 5
	out := *FilterAsync("", limit, workers, &collection)

	for _, actual := range out {
		expected := collection[actual.ID]
		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("actual %v, expected %v", actual, expected)
		}
	}
}

func TestQueueData(t *testing.T) {
	input := make(chan *models.CardBack)
	t.Run("should queue all items in the data collection", func(t *testing.T) {
		go queueData(&collection, input)
		count := 0
		for actual := range input {
			expected := (collection)[actual.ID]
			if !reflect.DeepEqual(*actual, expected) {
				t.Errorf("actual %v, expected %v", *actual, expected)
			}
			count++
		}
		if count != len(collection) {
			t.Errorf("actual %v, expected %v", count, len(collection))
		}
	})
}

func TestReadFilteredData(t *testing.T) {
	t.Run("should read all data coming from a channel into the out collection ", func(t *testing.T) {
		c := make(chan *models.CardBack)
		done := make(chan bool)
		out := []models.CardBack{}

		go queueData(&collection, c)
		go readFilteredData(c, &out, done)
		<-done // Waiting for completion...

		for _, actual := range out {
			expected := collection[actual.ID]
			if !reflect.DeepEqual(actual, expected) {
				t.Errorf("actual %v, expected %v", actual, expected)
			}
		}
		if len(out) != len(collection) {
			t.Errorf("actual %v, expected %v", len(out), len(collection))
		}
	})
}

func TestFilterWorker(t *testing.T) {
	t.Run("should push all the collection on empty filter and no limit", func(t *testing.T) {
		filterWorkerTest("", -1, &collection, t)
	})

	t.Run("should push all even items on even filter and no limit", func(t *testing.T) {
		filterWorkerTest("even", -1, &collection, t)
	})

	t.Run("should push all odd items on odd filter and no limit", func(t *testing.T) {
		filterWorkerTest("odd", -1, &collection, t)
	})

	t.Run("should push 1 even item on even filter and limit: 1", func(t *testing.T) {
		filterWorkerTest("even", 1, &collection, t)
	})

	t.Run("should push nothing on invalid filter", func(t *testing.T) {
		filterWorkerTest("invalid", -1, &collection, t)
	})
}

func filterWorkerTest(filter string, limit int, collection *[]models.CardBack, t *testing.T) {
	isLimit := limit > -1
	input := make(chan *models.CardBack)
	output := make(chan *models.CardBack)
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		wg.Wait()
		close(output)
	}()
	go filterWorker(filter, limit, input, output, &wg)

	go queueData(collection, input)

	for actual := range output {
		expected := (*collection)[actual.ID]
		if !reflect.DeepEqual(*actual, expected) {
			t.Errorf("actual %v, expected %v", *actual, expected)
		}
		limit--
	}
	if isLimit && limit < 0 {
		t.Errorf("actual %v, expected %v", limit, 0)
	}
}

func TestIsFilter(t *testing.T) {
	t.Run("should return true on odd filter and odd value", func(t *testing.T) {
		if !isFilter("odd", &collection[1]) {
			t.Errorf("actual %v, expected %v", false, true)
		}
	})

	t.Run("should return true on even filter and even value", func(t *testing.T) {
		if !isFilter("even", &collection[0]) {
			t.Errorf("actual %v, expected %v", false, true)
		}
	})

	t.Run("should return true on no filter and any value", func(t *testing.T) {
		if !isFilter("", &collection[0]) {
			t.Errorf("actual %v, expected %v", false, true)
		}
	})
}
