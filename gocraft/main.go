package main

import (
	"web-servers/gocraft/server"
)

func main() {
	if err := server.Run(8081); err != nil {
		panic(err)
	}
}
