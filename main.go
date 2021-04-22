package main

import (
	"go-api/app"
	"go-api/config"
	"go-api/routes"
	"log"
)

func main() {
	config, err := config.New("./config/config.json")
	if err != nil {
		log.Fatal(err)
	}

	app := app.New(config)
	router := routes.NewRouter(app)
	app.Run(router)
}
