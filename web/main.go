package main

import (
	"web-servers/web/server"
)

func main() {
	if err := server.Run(8081); err != nil {
		panic(err)
	}
}
