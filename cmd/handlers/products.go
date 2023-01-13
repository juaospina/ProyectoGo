package handler

import (
	"fmt"
	"errors"
	"strconv"
	"net/http"
	"ProyectoGo/internal/product"
	"github.com/gin-gonic/gin"
	"ProyectoGo/internal/domain"
	"ProyectoGo/pkg/response"
)

var (
	ErrToken     = errors.New("error: invalid token")
	InvalidId         = errors.New("error: invalid Id")
	InvalidCode  = errors.New("invalid expiration date, (format: DD/MM/YYYY)")
	InvalidExpiration = errors.New("there is already a product with that code")
)

func Ping(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}

func GetAllProducts(c *gin.Context) {
	services, _ := product.NewServices("products.json")
	products, err := services.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}
	c.JSON(http.StatusOK, products)
}

func GetProductByID(c *gin.Context) {
	services, _ := product.NewServices("products.json")
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, response.RequestError(InvalidId))
		return
	}
	product, _ := services.GetByID(id)
	c.JSON(http.StatusOK, product)
}

func CreateProduct(c *gin.Context) {
	services, _ := product.NewServices("products.json")
	var newProduct domain.Product
	c.BindJSON(&newProduct)
	createdProduct, _ := services.Create(newProduct)
	if createdProduct == nil{
		c.JSON(http.StatusBadRequest, response.RequestError(InvalidCode))
	} else {
		c.JSON(http.StatusCreated, "El producto fue creado con exito")
	}
	
}

func UpdateProduct(c *gin.Context) {
	services, _ := product.NewServices("products.json")
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, response.RequestError(InvalidId))
		return
	}
	var updatedProduct domain.Product
	c.BindJSON(&updatedProduct)
	updated, _ := services.Update(id, updatedProduct)
	c.JSON(http.StatusOK, updated)
}

func UpdateProductPatch(c *gin.Context) {
	services, _ := product.NewServices("products.json")
	id, _ := strconv.Atoi(c.Param("id"))
	var patchProduct domain.PatchRequest
	c.BindJSON(&patchProduct)
	updated, _ := services.UpdatePATCHservice(id, patchProduct)
	c.JSON(http.StatusOK, updated)
}

func DeleteProduct(c *gin.Context) {
	services, _ := product.NewServices("products.json")
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, response.RequestError(InvalidId))
		return
	}
	services.Delete(id)
	c.JSON(http.StatusOK, response.OKRequest("El producto ha sido eliminado con exito",fmt.Sprintf("id: %v",id)))
}

func SearchProduct(c *gin.Context) {
	services, _ := product.NewServices("products.json")
	priceGt := c.Query("priceGt")
	priceGtfloat, _ := strconv.ParseFloat(priceGt, 64)

	products, _ := services.Search(priceGtfloat)
	c.JSON(http.StatusOK, response.OKRequest(fmt.Sprintf("Se han obtenido los productos con precio mayor a %v, con exito", priceGt), products))
}
