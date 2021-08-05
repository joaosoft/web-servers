package main

import (
	"net/http"
	"web-servers/httprouter/routes"
)

func main() {
	if err := http.ListenAndServe(":8081", routes.Router); err != nil {
		panic(err)
	}
}
