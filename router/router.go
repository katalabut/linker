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
	router := mux.NewRouter().StrictSlash(true)
	router.NotFoundHandler = NotFoundHandler
	router.MethodNotAllowedHandler = NotFoundHandler

	for _, route := range routes {
		fmt.Println(route.Pattern)
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	return router
}
