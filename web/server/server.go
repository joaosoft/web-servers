package server

import (
	"fmt"
	"web-servers/web/routes"

	"github.com/joaosoft/web"
)

func Run(port int) error {
	routes.Router.Reconfigure(web.WithServerAddress(fmt.Sprintf(":%d", port)))
	return routes.Router.Start()
}
