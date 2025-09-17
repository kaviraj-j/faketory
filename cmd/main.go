package main

import (
	"github.com/kaviraj-j/faketory/internal/app"
	"github.com/kaviraj-j/faketory/internal/types"
)

func main() {
	app := app.NewApp(types.Config{
		AppUrl:  "localhost",
		AppPort: "8080",
	})
	app.Run()
}
