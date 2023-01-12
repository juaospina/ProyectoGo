package handler

import (
	"strconv"
	"net/http"
	"ProyectoGo/internal/product"
	"github.com/gin-gonic/gin"
	"ProyectoGo/internal/domain"
)

func Ping(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}

func GetAllProducts(c *gin.Context) {
	services, _ := product.NewServices("products.json")
	products, _ := services.GetAll()
	c.JSON(http.StatusOK, products)
}

func GetProductByID(c *gin.Context) {
	services, _ := product.NewServices("products.json")
	id, _ := strconv.Atoi(c.Param("id"))
	product, _ := services.GetByID(id)
	c.JSON(http.StatusOK, product)
}

func CreateProduct(c *gin.Context) {
	services, _ := product.NewServices("products.json")
	var newProduct domain.Product
	c.BindJSON(&newProduct)
	createdProduct, _ := services.Create(newProduct)
	c.JSON(http.StatusCreated, createdProduct)
}

func UpdateProduct(c *gin.Context) {
	services, _ := product.NewServices("products.json")
	id, _ := strconv.Atoi(c.Param("id"))
	var updatedProduct domain.Product
	c.BindJSON(&updatedProduct)
	updated, _ := services.Update(id, updatedProduct)
	c.JSON(http.StatusOK, updated)
}

func DeleteProduct(c *gin.Context) {
	services, _ := product.NewServices("products.json")
	id, _ := strconv.Atoi(c.Param("id"))
	services.Delete(id)
	c.JSON(http.StatusNoContent, gin.H{})
}

func SearchProduct(c *gin.Context) {
	services, _ := product.NewServices("products.json")
	priceGt := c.Query("priceGt")
	priceGtfloat, _ := strconv.ParseFloat(priceGt, 64)

	products, _ := services.Search(priceGtfloat)
	c.JSON(http.StatusOK, products)
}
