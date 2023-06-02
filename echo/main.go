package main

import (
	"github.com/joaosoft/web-servers/echo/server"
)

func main() {
	if err := server.New(8081).Start(); err != nil {
		panic(err)
	}
}
