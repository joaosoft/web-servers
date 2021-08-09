package server

import (
	"context"
	"fmt"
	"net/http"
	"web-servers/domain/server"
	"web-servers/httprouter/routes"

	"github.com/julienschmidt/httprouter"
)

type Server struct {
	App    *http.Server
	Router *RouterWrapper
	Port   int
}

func New(port int) server.IServer {
	router := NewRouterWrapper()
	server := &Server{
		App: &http.Server{
			Addr:    fmt.Sprintf(":%d", port),
			Handler: router,
		},
		Router: router,
		Port:   port,
	}

	return server
}

func (s *Server) Start() (err error) {
	routes.Init(s.Router.Router)
	return s.App.ListenAndServe()
}

func (s *Server) Stop() (err error) {
	return s.App.Shutdown(context.Background())
}

type MiddlewareFunc func(HandlerFunc) HandlerFunc
type HandlerFunc func(w http.ResponseWriter, req *http.Request) error

func NewRouterWrapper(middlewares ...MiddlewareFunc) *RouterWrapper {
	return &RouterWrapper{
		Router:      httprouter.New(),
		middlewares: middlewares,
	}
}

type RouterWrapper struct {
	*httprouter.Router
	middlewares []MiddlewareFunc
}

func (r *RouterWrapper) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	handlerRoute := func(w http.ResponseWriter, req *http.Request) error {
		// empty
		return nil
	}

	length := len(r.middlewares)
	for i, _ := range r.middlewares {
		if r.middlewares[length-1-i] != nil {
			handlerRoute = r.middlewares[length-1-i](handlerRoute)
		}
	}

	r.Router.ServeHTTP(w, req)
}
