package services

import (
	"fmt"
	. "github.com/katalabut/linker/router"
	. "github.com/katalabut/linker/services/linker"
	"github.com/katalabut/linker/tool/db"
	"net/http"
)

type Service interface {
	CreateRoutes() Routes
}
type Services []Service

func Run() Services {
	return Services{
		&Linker{
			Db: db.NewDBConnection(),
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
