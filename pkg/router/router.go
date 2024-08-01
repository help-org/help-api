package router

import (
	"database/sql"
	"net/http"

	"directory/internal/services/directory"
)

type routes map[string]map[string]http.HandlerFunc
type methods map[string]http.HandlerFunc

type Router struct {
	routes routes
}

func New(db *sql.DB) (router *Router) {
	router = &Router{
		routes: make(routes),
	}

	service := directory.NewService(db)

	router.addRoute("GET", "/directory", service.ListLocal)

	return
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	if handlers, ok := r.routes[req.URL.Path]; ok {
		if handler, methodExists := handlers[req.Method]; methodExists {
			handler(w, req)

			return
		}
	}

	http.NotFound(w, req)
}

func (r *Router) addRoute(method, path string, handler http.HandlerFunc) {
	if r.routes[path] == nil {
		r.routes[path] = make(methods)
	}

	r.routes[path][method] = handler
}
