package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Got A Visitor")
		fmt.Fprintf(w, "Hello Visitor")
	})

	http.ListenAndServe(":8080", nil)
}
