package main

import (
	"fmt"
	"github.com/katalabut/linker/router"
	"github.com/katalabut/linker/services"
	"net/http"
)

func main() {
	s := services.Run()
	routes, notFound := services.GetRoutes(s)
	router := router.InitRouter(routes, notFound)

	fmt.Println("open: http://localhost:8000")
	http.ListenAndServe(":8000", router)
}
