package main

import (
	"flag"
	"fmt"
	"os"

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

type application struct {
	config config
	logger *jsonlog.Logger
}

func main() {

	var cfg config

	flag.IntVar(&cfg.port, "port", 4000, "API Server Port")
	flag.StringVar(&cfg.env, "env", "development", "Environment")

	appVersion := flag.Bool("version", false, "Display version and exit")

	// Initilate a logger
	logger := jsonlog.New(os.Stdout, jsonlog.LevelInfo)

	flag.Parse()

	// check the version
	if *appVersion {
		fmt.Printf("Version:\t%s\n", version)
		fmt.Printf("Build time:\t%s\n", buildTime)
		os.Exit(0)
	}

	// instantiate the class(application) with the class parameters
	app := &application{
		config: cfg,
	}

	// call the serve methods on the instantiated class
	err := app.serve()
	if err != nil {
		logger.Printfatal(err, nil)
	}

}
