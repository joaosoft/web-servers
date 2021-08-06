package server

import (
	"fmt"
	"web-servers/gin/routes"
)

func Run(port int) error {
	return routes.Router.Run(fmt.Sprintf(":%d", port))
}
