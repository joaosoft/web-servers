package main

import "web-servers/gin/server"

func main() {
	if err := server.New(8081).Start(); err != nil {
		panic(err)
	}
}
