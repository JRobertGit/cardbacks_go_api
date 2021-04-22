package routes

import (
	"fmt"
	"go-api/app"
	"go-api/controllers"
	"go-api/repositories"
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouter(a *app.App) *mux.Router {
	r := mux.NewRouter().StrictSlash(true)

	// Repositories
	cbr := repositories.NewCardBackRepository(a.DataSource)

	// Controllers
	cbc := controllers.NewCardBackController(cbr)
	apic := controllers.NewExternalAPIController(a, cbr)

	// Routes
	r.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		res.Header().Set("Content-Type", "application/json; charset=utf-8")
		res.WriteHeader(http.StatusOK)
		fmt.Fprintf(res, "Hello, Go API is Working!")
	}).Methods(http.MethodGet)
	r.HandleFunc("/cardbacks", cbc.FilterAllAsync).Queries(
		"type", "{type}",
		"items", "{items}",
		"items_per_workers", "{items_per_workers}",
	).Methods(http.MethodGet)
	r.HandleFunc("/cardbacks", cbc.GetAll).Methods(http.MethodGet)
	r.HandleFunc("/cardbacks/{id}", cbc.GetById).Methods(http.MethodGet)
	r.HandleFunc("/external/cardbacks/{id}", apic.GetById).Methods(http.MethodGet)

	return r
}
