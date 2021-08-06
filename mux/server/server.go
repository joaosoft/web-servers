package server

import (
	"fmt"
	"net/http"
	"web-servers/mux/routes"
)

func Run(port int) (err error) {
	return http.ListenAndServe(fmt.Sprintf(":%d", port), routes.Router)
}
