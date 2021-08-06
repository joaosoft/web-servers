package server

import (
	"fmt"
	"net/http"
	"web-servers/goji/routes"
)

func Run(port int) error {
	return http.ListenAndServe(fmt.Sprintf(":%d", port), routes.Router)
}
