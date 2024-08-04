package main

import (
	"log"
	"os"
	"sync"
)

type application struct {
	config config
	wg     sync.WaitGroup
	logger *log.Logger
}

func main() {
	var cfg config
	loadConfig(&cfg)

	l := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	app := application{
		config: cfg,
		logger: l,
	}

	err := app.serve()

	if err != nil {
		app.logger.Fatal(err)
	}
}
