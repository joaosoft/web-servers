package main

import (
	"web-servers/iris/routes"

	"github.com/kataras/iris/v12"
)

func main() {
	if err := routes.Router.Run(iris.Addr(":8081")); err != nil {
		panic(err)
	}
}
