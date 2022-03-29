package main

import (
	"flag"
	"fmt"
	"os"
	"sync"

	"multi-cloud-architecture-converter/internal/jsonlog"
)

// Creae the entire application as Object and add all other functionalities to it as methods
// initiate the object in each function to use it

var (
	buildTime string
	version   string
)

// application configuration struct
type config struct {
	port int
	env  string
}

// application object with object properties
type application struct {
	config config
	logger *jsonlog.Logger
	wg     sync.WaitGroup
}

func main() {

	var cfg config

	flag.IntVar(&cfg.port, "port", 4000, "API Server Port")
	flag.StringVar(&cfg.env, "env", "development", "Environment")

	appVersion := flag.Bool("version", false, "Display version and exit")

	flag.Parse()

	// Initilate a logger
	logger := jsonlog.New(os.Stdout, jsonlog.LevelInfo)

	// check the version
	if *appVersion {
		fmt.Printf("Version:\t%s\n", version)
		fmt.Printf("Build time:\t%s\n", buildTime)
		os.Exit(0)
	}

	// instantiate the class(application) with the class parameters
	app := &application{
		config: cfg,
		logger: logger,
	}

	// call the serve methods on the instantiated class
	err := app.serve()
	if err != nil {
		logger.PrintFatal(err, nil)

	}

}
