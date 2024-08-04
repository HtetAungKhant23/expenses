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

	db, err := openDB(&cfg.db)
	if err != nil {
		l.Fatalln(err)
	}
	l.Println("Successfully connected to the database.")
	defer db.Close()

	app := application{
		config: cfg,
		logger: l,
	}

	err = app.serve()

	if err != nil {
		app.logger.Fatalln(err)
	}
}
