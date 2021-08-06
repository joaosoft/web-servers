package server

import (
	"fmt"
	"net/http"
	"web-servers/gocraft/routes"
)

func Run(port int) error {
	return http.ListenAndServe(fmt.Sprintf(":%d", port), routes.Router)
}
