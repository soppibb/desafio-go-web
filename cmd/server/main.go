package main

import (
	"github.com/bootcamp-go/desafio-go-web/cmd/server/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	routes := routes.NewRouter(r)
	if err := routes.SetTicketsRoutes(); err != nil {
		panic(err)
	}

	if err := r.Run(); err != nil {
		panic(err)
	}
	//soppibb sos genial

}
