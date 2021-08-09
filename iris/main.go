package main

import "web-servers/iris/server"

func main() {
	if err := server.New(8081).Start(); err != nil {
		panic(err)
	}
}
