package middlewares

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type MiddlewareFunc func(HandlerFunc) HandlerFunc
type HandlerFunc func(w http.ResponseWriter, req *http.Request) error

func NewRouter(middlewares ...MiddlewareFunc) *Router {
	return &Router{
		Router:      httprouter.New(),
		middlewares: middlewares,
	}
}

type Router struct {
	*httprouter.Router
	middlewares []MiddlewareFunc
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
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
