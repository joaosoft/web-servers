package server

import (
	"web-servers/revel/app/tmp/run"

	"github.com/revel/revel"
)

func Run(port int) (err error) {
	revel.Init("dev", "web-servers", "")
	run.Run(port)

	return nil
}
