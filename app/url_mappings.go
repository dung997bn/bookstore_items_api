package app

import (
	"net/http"

	"github.com/dung997bn/bookstore_items_api/controllers"
)

func mapUrls() {
	router.HandleFunc("/ping", controllers.PingController.Ping).Methods(http.MethodGet)
	router.HandleFunc("/items/{id}", controllers.ItemController.GetByID).Methods(http.MethodGet)
	router.HandleFunc("/items", controllers.ItemController.Create).Methods(http.MethodPost)

	router.HandleFunc("/items/search", controllers.ItemController.Search).Methods(http.MethodPost)

}
