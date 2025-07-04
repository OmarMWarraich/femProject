package main

import (
	"github.com/OmarMWarraich/femProject/internal/app"
)

func main() {
	app, err := app.NewApplication()
	if err != nil {
		panic(err)
	}

	app.Logger.Println(("We are running our app"))
}