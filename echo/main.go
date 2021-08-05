package main

import (
	"fmt"
	"net/http"
	"web-servers/echo/routes"
	_ "web-servers/echo/routes"
)

func main() {
	fmt.Println("server started at http://localhost:8081/")
	if err := http.ListenAndServe(":8081", routes.Router); err != nil {
		panic(err)
	}
}
