package server

import (
	"fmt"
	"github.com/joaosoft/web-servers/domain/server"
	"github.com/joaosoft/web-servers/fasthttp/routes"
	"time"

	routing "github.com/qiangxue/fasthttp-routing"

	"github.com/valyala/fasthttp"
)

type Server struct {
	App    *fasthttp.Server
	Router *routing.Router
	Port   int
}

func New(port int) server.IServer {
	router := routing.New()
	server := &Server{
		App: &fasthttp.Server{
			Handler:     router.HandleRequest,
			IdleTimeout: time.Minute,
		},
		Router: router,
		Port:   port,
	}

	return server
}

func (s *Server) Start() (err error) {
	s.App.Handler = s.Router.HandleRequest
	routes.Init(s.Router)
	return s.App.ListenAndServe(fmt.Sprintf(":%d", s.Port))
}

func (s *Server) Stop() (err error) {
	return s.App.Shutdown()
}
