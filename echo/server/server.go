package server

import (
	"fmt"
	"web-servers/echo/routes"
)

func Run(port int) error {
	return routes.Router.Start(fmt.Sprintf(":%d", port))
}
