package main

import (
	"fmt"
	"net/http"

	"github.com/katalabut/linker/router"
	"github.com/katalabut/linker/services"
)

func main() {
	s := services.Run()
	routes, notFound := services.GetRoutes(s)
	router := router.InitRouter(routes, notFound)

	fmt.Println("open: http://localhost:8080")
	http.ListenAndServe(":8080", router)
}
