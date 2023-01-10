package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
	"github.com/gin-gonic/gin"
)

type Product struct {
	Id          int
	Name        string
	Quantity    int
	Code_value  string
	Is_published bool
	Expiration  string
	Price       float64 
}

type Products struct {
	list []Product
	id   int
}

func (p *Products) LoadFromJSON(file string) error {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, &p.list)
	if err != nil {
		return err
	}

	return nil
}

func (p *Products) Search(priceGt float64) []Product {
	var result []Product
	for _, product := range p.list {
		if product.Price > priceGt {
			result = append(result, product)
		}
	}

	return result
}

func (p *Products) GetProduct(id int) (*Product, error) {
	for _, product := range p.list {
		if product.Id == id {
			return &product, nil
		}
	}

	return nil, fmt.Errorf("El Producto con ID %d no ha sido encontrado", id)
}

func main() {
	products := &Products{}

	err := products.LoadFromJSON("products.json")
	if err != nil {
		fmt.Printf("Error cargando los productos desde JSON: %s", err)
	}

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	r.GET("/products", func(c *gin.Context) {
		c.JSON(http.StatusOK, products.list)
	})

	r.GET("/products/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.String(http.StatusBadRequest, "ID de producto invalido")
			return
		}

		product, err := products.GetProduct(id)
		if err != nil {
			c.String(http.StatusNotFound, err.Error())
			return
		}

		c.JSON(http.StatusOK, product)
	})

	r.GET("/products/search", func(c *gin.Context) {
		priceGt, err := strconv.ParseFloat(c.Query("priceGt"), 64)
		if err != nil {
			c.String(http.StatusBadRequest, "Parametro priceGt invalido")
			return
		}

		result := products.Search(priceGt)
		c.JSON(http.StatusOK, result)
	})

	r.POST("/products", func(c *gin.Context) {
		var newProduct Product
		err := c.BindJSON(&newProduct)
		if err != nil {
			c.String(http.StatusBadRequest, "Formato de producto invalido")
			return
		}
	
		if newProduct.Name == "" || newProduct.Quantity < 0 || newProduct.Code_value == "" || newProduct.Price < 0 || newProduct.Expiration == "" {
			c.String(http.StatusBadRequest, "Alguno de los campos obligatorios esta vacio o es invalido")
			return
		}
	
		// Check that the code_value is unique
		for _, product := range products.list {
			if product.Code_value == newProduct.Code_value {
				c.String(http.StatusBadRequest, "El valor del codigo ya ha sido utilizado")
				return
			}
		}
	
		// Check that the expiration date is in the correct format and is a valid date
		_, err = time.Parse("02/01/2006", newProduct.Expiration)
		if err != nil {
			c.String(http.StatusBadRequest, "La fecha de vencimiento no tiene el formato correcto o es una fecha invalida")
			return
		}
	
		// Assign a new unique ID to the product
		products.id = len(products.list)+1
		fmt.Println(products.id)
		newProduct.Id = products.id
		
	
		// Add the new product to the list
		products.list = append(products.list, newProduct)
	
		c.JSON(http.StatusOK, newProduct)
	})

	r.Run(":8080")
}