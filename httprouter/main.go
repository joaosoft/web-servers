package main

import "github.com/joaosoft/web-servers/httprouter/server"

func main() {
	if err := server.New(8082).Start(); err != nil {
		panic(err)
	}
}
