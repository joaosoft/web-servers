package server

import (
	"os"
	"strconv"
	"web-servers/martini/routes"
)

func Run(port int) (err error) {
	if err = os.Setenv("PORT", strconv.Itoa(port)); err != nil {
		return err
	}

	routes.Router.Run()

	return nil
}