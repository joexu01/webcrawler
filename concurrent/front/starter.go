package main

import (
	"fmt"
	"net/http"
	"webcrawler/concurrent/front/controller"
)

func main() {
	fmt.Println("localhost:8000")
	http.Handle("/", http.FileServer(http.Dir("concurrent/front/view")))
	http.Handle("/search", controller.CreateSearchResultHandler(
		"concurrent/front/view/template.html"))
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		panic(err)
	}
}
