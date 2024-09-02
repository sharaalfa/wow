package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"wow/server/app"
)

func main() {
	newApp, err := app.NewApp()
	if err != nil {
		log.Fatalf("Failed to create newApp: %v", err)
	}

	go newApp.Run()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	if err := newApp.Close(); err != nil {
		log.Fatalf("Failed to close newApp: %v", err)
	}
	log.Println("Server stopped")
}
