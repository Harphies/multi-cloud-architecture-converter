package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// create a serve method on the application class

func (app *application) serve() error {

	// declare the HTTP server
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", app.config.port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	// Create a shutdownError channel--used to receive any error returned by the graceful shutdown
	shutdownError := make(chan error) // normal channel

	go func() {
		// Intercept the signal into a quit channel
		quit := make(chan os.Signal, 1) // buffered channel is used to avoid never missed a signal

		// use the signal.Notify to listen to incoming SIGINT and SIGTERM signals and relay them to the quit channel
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

		// Read the signal from quit channel into s variable.
		// This code will block until a signal is received
		// it can not go down to the next stage
		s := <-quit

		// Log a message when a signal is received and stringify the signal put into s
		app.logger.PrintInfo("caught signal", map[string]string{
			"signal": s.String(),
		})

		// Create a context with 5-second timeout
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		// only if there is an error
		err := srv.Shutdown(ctx)
		if err != nil {
			shutdownError <- err
		}

		// Log a message to notify we are waiting for background goroutine to complete their tasks
		app.logger.PrintInfo("completing background tasks", map[string]string{
			"add": srv.Addr,
		})

		// call Wait to block until our WaitGroup counter is zero--blocking until the background groutines have finished
		// return nil on the shutdownError channel to indicate that the shutdown completed without any issues
		app.wg.Wait()
		shutdownError <- nil
	}()

	// start the http server
	app.logger.PrintInfo("starting %s server on %s", map[string]string{
		"addr": srv.Addr,
		"env":  app.config.env,
	})

	// check for indication the server begin graceful shutdown
	// ErrServerClosed is an indication, check for that
	err := srv.ListenAndServe()
	if !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	// Otherwise wait to receive the return from Shutdown() method on the shutdownError channel
	// if there is an error, we know there is a problem with the shutdown
	err = <-shutdownError
	if err != nil {
		return err
	}

	// At this point, we know the graceful shutdown is successful
	app.logger.PrintInfo("stopped server", map[string]string{
		"add": srv.Addr,
	})

	// start the server and return any error
	return nil
}
