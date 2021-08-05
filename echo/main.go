package main

import (
	"web-servers/echo/routes"
	_ "web-servers/echo/routes"
)

func main() {
	if err := routes.Router.Start(":8081"); err != nil {
		panic(err)
	}
}
