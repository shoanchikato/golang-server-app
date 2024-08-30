package main

import (
	d "app/pkg/di"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	app := d.Di()
	defer app.DB.Close()

	dep := app.Valid.Permission
	dp := app.Valid.PermissionManagement

	// _, _, _, pp, _ := Data()

	pp, err := dep.GetAll(0, 50)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, p := range *pp {
		err := dp.AddPermissionToRole(p.Id, 1)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	fmt.Println(pp)
}
