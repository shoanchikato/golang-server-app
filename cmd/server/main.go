package main

import (
	d "app/pkg/di"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	app := d.Di()
	defer app.DB.Close()

	go func() {
		if err := app.App.Listen(":3000"); err != nil {
			log.Fatal("App stopped", err)
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	sig := <-sigChan
	fmt.Println("\nReceived signal:", sig)

	fmt.Println("Application shutting down...")
}
