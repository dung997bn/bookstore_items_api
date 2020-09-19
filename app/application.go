package app

import (
	"net/http"
	"time"

	"github.com/dung997bn/bookstore_items_api/client/elasticsearch"
	"github.com/gorilla/mux"
)

var (
	router = mux.NewRouter()
)

//StartApplication func
func StartApplication() {
	elasticsearch.Init()
	mapUrls()
	srv := &http.Server{
		Handler: router,
		Addr:    "127.0.0.1:8082",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	if err := srv.ListenAndServe(); err != nil {
		panic(err)
	}
}
