package main

import (
	d "app/pkg/di"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	app := d.Di()
	defer app.DB.Close()

	if err := app.App.Listen(":3000"); err != nil {
		log.Fatal("App stopped", err)
	}
}
