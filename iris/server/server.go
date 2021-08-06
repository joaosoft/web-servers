package server

import (
	"fmt"
	"web-servers/iris/routes"

	"github.com/kataras/iris/v12"
)

func Run(port int) error {
	return routes.Router.Run(
		iris.Addr(fmt.Sprintf(":%d", port)),
		iris.WithOptimizations,
	)
}
