package main

import (
	"go_microservices/handlers"
	"log"
	"net/http"
	"os"
)

func main() {

	l := log.New(os.Stdout, "product-api", log.LstdFlags)

	hh := handlers.NewHello(l)
	gh := handlers.NewGoodBye(l)

	sm := http.NewServeMux()
	sm.Handle("/", hh)
	sm.Handle("/bye", gh)

	http.ListenAndServe(":8080", sm)
}
