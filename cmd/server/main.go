package main

import (
	d "app/pkg/di"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	env, err := godotenv.Read()
	if err != nil {
		log.Fatal("error loading env map")
	}

	app := d.Di(
		env["DB_DRIVER"],
		env["DB_CONN"],
		env["SECRET"],
		os.Stdout,
	)
	defer app.DB.Close()

	go func() {
		if err := app.App.Listen(":3000"); err != nil {
			log.Fatal("App stopped with an error: ", err)
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	sig := <-sigChan
	fmt.Println("\nReceived signal:", sig)

	fmt.Println("Application shutting down...")
}
