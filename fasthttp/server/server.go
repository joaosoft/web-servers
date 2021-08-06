package server

import (
	"fmt"
	"web-servers/fasthttp/routes"

	"github.com/valyala/fasthttp"
)

func Run(port int) error {
	return fasthttp.ListenAndServe(fmt.Sprintf(":%d", port), routes.Router.HandleRequest)
}
