package main

import (
	d "app/pkg/di"
	"encoding/json"
	"fmt"
	"os"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	app := d.Di()
	defer app.DB.Close()

	dep := app.Jwt

	// _, _, _, pp, _ := Data()

	tokenStr, err := dep.GetAccessToken(2342)
	if err != nil {
		fmt.Println(err)
		return
	}

	newToken := ""
	if len(os.Args) < 2 {
		newToken = tokenStr
	} else {
		newToken = os.Args[1]
	}

	token, err := dep.ParseToken(&newToken)
	if err != nil {
		fmt.Println(err)
		return
	}

	json.NewEncoder(os.Stdout).Encode(token)

	expires, _ := token.GetExpires()
	issued, _ := token.GetIssued()

	hasExpired, _ := token.HasExpired()

	fmt.Println(tokenStr, time.Now().After(expires), issued, hasExpired)
}
