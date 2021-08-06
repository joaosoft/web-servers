package server

import (
	"web-servers/beego/routes"

	"github.com/astaxie/beego"
)

func Run(port int) error {
	beego.BConfig.Listen.HTTPPort = port
	routes.Router.Run()

	return nil
}
