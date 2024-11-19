package main

import (
	"fmt"
	"log"
	"net/http"
)

const webPort = "8086"

type Config struct{}

func main() {
	app := Config{}
	log.Printf("Starting Broker Service On Port %s\n", webPort)

	// define server
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	// start the server
	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}
