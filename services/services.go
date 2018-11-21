package services

import (
	"fmt"
	. "linker/router"
	. "linker/services/linker"
	"linker/tool/db"
	"net/http"
)

type Service interface {
	CreateRoutes() Routes
}
type Services []Service

func Run() Services {
	return Services{
		&Linker{
			db.NewDBConnection(),
		},
	}
}

func GetRoutes(s Services) (Routes, http.HandlerFunc) {
	routes := Routes{}
	for _, service := range s {
		routes = append(routes, service.CreateRoutes()...)
	}
	return routes, http.HandlerFunc(notFoundHandle)
}

func notFoundHandle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "notFoundHandle")
	return
}
