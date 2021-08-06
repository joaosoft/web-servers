package main

import (
	"web-servers/buffalo/routes"
)

func main() {
	if err := routes.Router.Serve(); err != nil {
		panic(err)
	}
}
