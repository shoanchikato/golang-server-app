package main

import (
	d "app/pkg/di"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	app := d.Di()
	defer app.DB.Close()

	dep := app.Auth.Permission

	// _, _, _, pp, _ := Data()

	pp, err := dep.GetAll(2, 0, 50)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(len(*pp))
}
