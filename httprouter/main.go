package main

import "web-servers/httprouter/server"

func main() {
	if err := server.New(8082).Start(); err != nil {
		panic(err)
	}
}
