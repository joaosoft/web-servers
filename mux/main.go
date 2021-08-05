package main

import (
	"fmt"
	"net/http"
	"web-servers/mux/routes"
	_ "web-servers/mux/routes"
)

func main() {
	fmt.Println("server started at http://localhost:8081/")
	if err := http.ListenAndServe(":8081", routes.Router); err != nil {
		panic(err)
	}
}
