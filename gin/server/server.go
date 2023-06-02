package server

import (
	"context"
	"fmt"
	"github.com/joaosoft/web-servers/domain/server"
	"github.com/joaosoft/web-servers/gin/routes"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Server struct {
	App    *http.Server
	Router *gin.Engine
	Port   int
}

func New(port int) server.IServer {
	gin.SetMode(gin.ReleaseMode)

	router := gin.New()
	server := &Server{
		App: &http.Server{
			Addr:    fmt.Sprintf(":%d", port),
			Handler: router,
		},
		Router: router,
		Port:   port,
	}

	router.Use(gin.Recovery())

	return server
}

func (s *Server) Start() (err error) {
	routes.Init(s.Router)
	return s.App.ListenAndServe()
}

func (s *Server) Stop() (err error) {
	return s.App.Shutdown(context.Background())
}
