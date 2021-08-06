package main

import (
	"web-servers/fasthttp/server"
)

func main() {
	if err := server.Run(8081); err != nil {
		panic(err)
	}
}
