package main

import (
	"web-servers/gin/routes"
)

func main() {
	if err := routes.Router.Run(":8081"); err != nil {
		panic(err)
	}
}
