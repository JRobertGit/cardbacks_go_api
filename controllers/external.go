package controllers

import (
	"encoding/json"
	"fmt"
	"go-api/app"
	"go-api/models"
	"go-api/repositories"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

type ExternalAPIController struct {
	*app.App
	repositories.CardBackRepository
}

func NewExternalAPIController(app *app.App, cbr repositories.CardBackRepository) *ExternalAPIController {
	return &ExternalAPIController{app, cbr}
}

func (apic *ExternalAPIController) GetById(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Access-Control-Allow-Origin", "*")
	params := mux.Vars(req)
	extReq, err := http.NewRequest("GET", apic.Config.ExternalAPI.BaseURL+"/"+params["id"], nil)
	if err != nil {
		fmt.Printf(err.Error())
		http.Error(res, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	q := extReq.URL.Query()
	q.Add("locale", "en_US")
	extReq.URL.RawQuery = q.Encode()

	resp, err := apic.ExternalClient.Do(extReq)
	if err != nil {
		fmt.Printf(err.Error())
		http.Error(res, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf(err.Error())
		http.Error(res, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	res.Header().Set("Content-Type", "application/json; charset=utf-8")
	res.WriteHeader(http.StatusOK)
	fmt.Fprintf(res, string(body))

	cardBack := models.CardBack{}
	err = json.Unmarshal(body, &cardBack)
	if err != nil {
		fmt.Printf(err.Error())
		return
	}
	if cardBack != (models.CardBack{}) {
		err = apic.Create(&cardBack)
		if err != nil {
			fmt.Printf(err.Error())
			return
		}
	}
}
