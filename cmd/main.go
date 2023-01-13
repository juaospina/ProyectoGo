package main

import (
	"ProyectoGo/cmd/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	//Se inicializa el servidor haciendo uso de Gin
	router := gin.Default()

	//Se inicializan las rutas a traves de el SetUp de la carpeta routes
	routes.Setup(router)

	//Se corre el servidor en el puerto de prueba 8080
	router.Run(":8080")
}