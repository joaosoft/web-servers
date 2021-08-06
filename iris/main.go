package main

import (
	"web-servers/iris/routes"

	"github.com/kataras/iris"
)

func main() {
	if err := routes.Router.Run(iris.Addr(":8081")); err != nil {
		panic(err)
	}
}
