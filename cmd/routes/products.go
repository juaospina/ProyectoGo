package routes

import (
	"ProyectoGo/cmd/handlers"
	"github.com/gin-gonic/gin"
)

func Setup(router *gin.Engine) {
	//Se crea un primer grupo de rutas para la Api /api
	api := router.Group("/api")

	{
		//Se crea la ruta para la función Pong, /ping
		api.GET("/ping", handler.Ping)

		//Se agrupan todas las rutas para /products
		productGroup := api.Group("/products")
		{
			productGroup.GET("", handler.GetAllProducts)
			productGroup.GET(":id", handler.GetProductByID)
			productGroup.POST("", handler.CreateProduct)
			productGroup.PUT(":id", handler.UpdateProduct)
			productGroup.PATCH(":id", handler.UpdateProductPatch)
			productGroup.DELETE(":id", handler.DeleteProduct)
			productGroup.GET("/search", handler.SearchProduct)
		}
	}
}
