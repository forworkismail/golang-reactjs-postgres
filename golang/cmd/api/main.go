package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

const version = "1.0.0"

type config struct {
	port int
	env  string
}

type application struct {
	config config
	logger *log.Logger
}

func main() {
	var cfg config
	portNumber := 4000

	flag.IntVar(&cfg.port, "port", portNumber, "Server port to listen on")
	flag.StringVar(&cfg.env, "env", "development", "Application environment (development|production)")
	flag.Parse()

	// logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	// app := &application{
	// 	config: cfg,
	// 	logger: logger,
	// }

	fmt.Println("Running on port", portNumber)

	err := http.ListenAndServe(fmt.Sprintf(":%d", cfg.port), nil)
	if err != nil {
		log.Println(err)
	}

}
