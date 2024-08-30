package main

import (
	d "app/pkg/di"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	app := d.Di()
	defer app.DB.Close()

	dep := app.Valid.User

	_, _, _, pp, _ := Data()

	err := dep.AddAll(&pp)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(pp)
}
