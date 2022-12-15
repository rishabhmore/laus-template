package server

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/rishabhmore/go-hustle-template/tools/logger"
)

func StartEchoServer() {
	// Temp setup
	// Setup
	e := echo.New()

	// Start server
	go func() {
		logger.Logger.Info("Server starting up... ")
		if err := e.Start(":1323"); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("shutting down the server")
		} else {
			e.Logger.Info("Shutting down gracefully...")
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 10 seconds.
	// Use a buffered channel to avoid missing signals as recommended for signal.Notify
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	// With a timeout of 10 seconds, we will wait for the server to shutdown gracefully
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	// Finally we will cancel the timeout at the end of execution
	defer cancel()

	// Shut downs the server gracefully.
	// We will log any errors as fatal exit
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	} else {
		e.Logger.Info("Shutdowned: Server offline")
	}
}
