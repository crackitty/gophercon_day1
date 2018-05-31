package main

import (
	"log"
	"net/http"
	"os"

	"github.com/crackitty/gophercon_day1/cmd/routing"
	"github.com/crackitty/gophercon_day1/cmd/webserver"
)

func main() {
	log.Printf("Service is starting")

	port := os.Getenv("SERVICE_PORT")

	if len(port) == 0 {
		log.Fatal("Port cannot be zero")
	}

	router := routing.BaseRouter()
	ws := webserver.New("", port, router)
	//http.HandleFunc("/", handler())
	log.Fatal(ws.Start())

	http.ListenAndServe(":8080", router)
}
