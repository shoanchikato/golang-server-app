package main

import (
	d "app/pkg/di"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	app := d.Di()
	defer app.DB.Close()

	dep := app.Valid.Author

	// _, _, _, _, _ = Data()

	pp, err := dep.GetAll(0, 0)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(pp)
}
