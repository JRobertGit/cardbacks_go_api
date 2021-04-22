package util

import (
	"go-api/models"
	"sync"
)

const (
	odd  = "odd"
	even = "even"
	none = ""
)

func FilterAsync(filter string, limit int, workers int, data *[]models.CardBack) *[]models.CardBack {
	fData := []models.CardBack{}

	ci := make(chan *models.CardBack)
	go queueData(data, ci)

	done := make(chan bool)
	co := make(chan *models.CardBack)
	go readFilteredData(co, &fData, done)

	var wg sync.WaitGroup
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go filterWorker(filter, limit, ci, co, &wg)
	}
	wg.Wait()
	close(co)

	<-done
	return &fData
}

func queueData(data *[]models.CardBack, co chan<- *models.CardBack) {
	for _, i := range *data {
		cp := i
		co <- &cp
	}
	close(co)
}

func readFilteredData(ci <-chan *models.CardBack, out *[]models.CardBack, done chan<- bool) {
	for item := range ci {
		*out = append(*out, *item)
	}
	done <- true
}

func filterWorker(filter string, limit int, ci <-chan *models.CardBack, co chan<- *models.CardBack, wg *sync.WaitGroup) {
	for cb := range ci {
		if limit == 0 {
			break
		}
		if isFilter(filter, cb) {
			co <- cb
			limit--
		}
	}
	wg.Done()
}

func isFilter(filter string, cb *models.CardBack) bool {
	return (filter == even && cb.ID%2 == 0) ||
		(filter == odd && cb.ID%2 != 0) ||
		filter == none
}
