package main

import (
	"os"
	"web-servers/martini/routes"
)

func main() {
	if err := os.Setenv("PORT", "8081"); err != nil {
		panic(err)
	}

	routes.Router.Run()
}
