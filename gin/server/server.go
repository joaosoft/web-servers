package server

import (
	"context"
	"fmt"
	"net/http"
	"web-servers/domain/server"
	"web-servers/gin/routes"

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
