package main

import (
	"log"
	"web-servers/fasthttp/routes"
	_ "web-servers/fasthttp/routes"

	"github.com/valyala/fasthttp"
)

func main() {
	if err := fasthttp.ListenAndServe(":8081", routes.Router.HandleRequest); err != nil {
		log.Fatalf("Error in ListenAndServe: %s", err)
	}
}
