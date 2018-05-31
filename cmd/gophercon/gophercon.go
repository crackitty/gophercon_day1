package main

import (
	"log"
	"net/http"

	"github.com/crackitty/gophercon_day1/cmd/routing"
)

func main() {
	log.Printf("Service is starting")

	r := routing.BaseRouter()

	http.ListenAndServe(":8080", r)
}
