package router

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func InitRouter(routes Routes, NotFoundHandler http.HandlerFunc) *mux.Router {
	r := mux.NewRouter().StrictSlash(true)

	for _, route := range routes {
		fmt.Println(route.Pattern)
		r.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	r.NotFoundHandler = NotFoundHandler
	r.MethodNotAllowedHandler = NotFoundHandler
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))

	return r
}
