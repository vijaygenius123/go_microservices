package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		log.Println("Request")
		d, err := ioutil.ReadAll(r.Body)

		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			rw.Write([]byte("Oops"))
			return
		}

		fmt.Fprintf(rw, "Hello %s", d)

	})

	http.ListenAndServe(":8080", nil)
}
