package main

import (
	"web-servers/gocraft/server"
)

func main() {
	if err := server.New(8081).Start(); err != nil {
		panic(err)
	}
}
