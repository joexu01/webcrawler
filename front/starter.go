package main

import (
	"net/http"
	"webcrawler/front/controller"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("front/view")))
	http.Handle("/search", controller.CreateSearchResultHandler(
		"front/view/template.html"))
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		panic(err)
	}
}
