package setup

import (
	d "app/pkg/di"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
)

func Run() d.DI {
	filePath := "../../../.env"

	err := godotenv.Load(filePath)
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}

	env, err := godotenv.Read(filePath)
	if err != nil {
		log.Fatal("error loading env map")
	}

	app := d.Di(
		env["DB_DRIVER_TEST"],
		env["DB_CONN_TEST"],
		env["SECRET_TEST"],
		os.Stdout,
	)

	d.AddAdminUser(app.Validators)

	return app
}
