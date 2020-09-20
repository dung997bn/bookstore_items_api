package main

import (
	"os"

	"github.com/dung997bn/bookstore_items_api/app"
)

func main() {
	os.Setenv("LOG_LEVEL", "info")
	app.StartApplication()
}
