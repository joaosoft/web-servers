package server

import (
	"fmt"
	"web-servers/buffalo/routes"
)

func Run(port int) error {
	routes.Router.Options.Addr = fmt.Sprintf(":%d", port)

	return routes.Router.Serve()
}
