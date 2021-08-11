package middlewares

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

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

func (r *RouterWrapper) Middleware(m ...MiddlewareFunc) {
	r.middlewares = append(r.middlewares, m...)
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
