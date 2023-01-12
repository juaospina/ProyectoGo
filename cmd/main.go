package main

import (
	"ProyectoGo/cmd/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	routes.Setup(router)
	router.Run(":8080")
}