package routes

import (
	"ProyectoGo/cmd/handlers"
	"github.com/gin-gonic/gin"
)

func Setup(router *gin.Engine) {
	api := router.Group("/api")

	{
		api.GET("/ping", handler.Ping)

		productGroup := api.Group("/products")
		{
			productGroup.GET("", handler.GetAllProducts)
			productGroup.GET(":id", handler.GetProductByID)
			productGroup.POST("", handler.CreateProduct)
			productGroup.PUT(":id", handler.UpdateProduct)
			productGroup.PATCH(":id", handler.UpdateProduct)
			productGroup.DELETE(":id", handler.DeleteProduct)
			productGroup.GET("/search", handler.SearchProduct)
		}
	}
}
