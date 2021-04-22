package controllers

import (
	"encoding/json"
	"fmt"
	"go-api/repositories"
	"go-api/util"
	"math"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type CardBackController struct {
	repositories.CardBackRepository
}

func NewCardBackController(cbr repositories.CardBackRepository) *CardBackController {
	return &CardBackController{cbr}
}

func (cbc *CardBackController) GetAll(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Access-Control-Allow-Origin", "*")
	res.Header().Set("Content-Type", "application/json; charset=utf-8")
	res.WriteHeader(http.StatusOK)

	cb, err := cbc.CardBackRepository.GetAll()
	if err != nil {
		fmt.Printf(err.Error())
		http.Error(res, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(res).Encode(cb)
	if err != nil {
		fmt.Printf(err.Error())
		http.Error(res, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}

func (cbc *CardBackController) FilterAllAsync(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Access-Control-Allow-Origin", "*")
	res.Header().Set("Content-Type", "application/json; charset=utf-8")
	res.WriteHeader(http.StatusOK)

	params := req.URL.Query()
	filter := params.Get("type")
	items := params.Get("items")
	items_per_workers := params.Get("items_per_workers")

	cb, err := cbc.CardBackRepository.GetAll()
	if err != nil {
		fmt.Printf(err.Error())
		http.Error(res, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	i, err := strconv.Atoi(items)
	if err != nil {
		fmt.Printf(err.Error())
		http.Error(res, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	ipw, err := strconv.Atoi(items_per_workers)
	if err != nil {
		fmt.Printf(err.Error())
		http.Error(res, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	workers := int(math.Ceil(float64(i) / float64(ipw)))
	cb = util.FilterAsync(filter, ipw, workers, cb)

	limit := int(math.Min(float64(i), float64(len(*cb))))
	json.NewEncoder(res).Encode((*cb)[0:limit])
	if err != nil {
		fmt.Printf(err.Error())
		http.Error(res, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}

func (cbc *CardBackController) GetById(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Access-Control-Allow-Origin", "*")
	res.Header().Set("Content-Type", "application/json; charset=utf-8")
	params := mux.Vars(req)
	key, err := strconv.Atoi(params["id"])
	if err != nil {
		fmt.Printf(err.Error())
		http.Error(res, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	cb, err := cbc.CardBackRepository.FindByID(key)
	if err != nil {
		fmt.Printf(err.Error())
		http.Error(res, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	if cb == nil {
		http.Error(res, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	res.WriteHeader(http.StatusOK)
	err = json.NewEncoder(res).Encode(cb)
	if err != nil {
		fmt.Printf(err.Error())
		http.Error(res, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}
