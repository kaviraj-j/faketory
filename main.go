package main

import (
	"os"

	"github.com/kaviraj-j/faketory/internal/app"
	"github.com/kaviraj-j/faketory/internal/types"
)

func main() {
	port := os.Getenv("PORT")
	app := app.NewApp(types.Config{
		AppUrl:  "",
		AppPort: port,
	})
	app.Run()
}
