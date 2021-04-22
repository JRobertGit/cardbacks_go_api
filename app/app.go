package app

import (
	"context"
	"fmt"
	"go-api/api"
	"go-api/config"
	"go-api/data"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type App struct {
	Config         *config.Config
	DataSource     *data.FileSource
	ExternalClient *http.Client
}

func New(config *config.Config) *App {
	cxt := context.Background()

	dataSource, err := data.NewFileSource(config.DataSources.CSV)
	if err != nil {
		log.Fatal(err)
	}

	client, err := api.GetAuthenticatedClient(cxt, config.ExternalAPI)
	if err != nil {
		log.Fatal(err)
	}

	return &App{config, dataSource, client}
}

func (a *App) Run(router *mux.Router) {
	port := a.Config.Port
	addr := fmt.Sprintf(":%v", port)
	fmt.Printf("APP is listening on port: %d\n", port)
	log.Fatal(http.ListenAndServe(addr, router))
}
