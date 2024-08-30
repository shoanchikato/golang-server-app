package main

import (
	d "app/pkg/di"
	"encoding/json"
	"fmt"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	app := d.Di()
	defer app.DB.Close()

	dep := app.Jwt

	// _, _, _, pp, _ := Data()

	token, err := dep.GetRefreshToken(2342)
	if err != nil {
		fmt.Println(err)
		return
	}

	details, err := dep.ParseToken(&token)
	if err != nil {
		fmt.Println(err)
		return
	}

	s := strings.Builder{}
	json.NewEncoder(&s).Encode(details)

	fmt.Println(s.String(), token)
}
