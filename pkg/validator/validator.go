package validator

import (
    "errors"
    "regexp"
	"ProyectoGo/internal/domain"
)

// ValidateProduct valida los campos de un producto
func ValidateProduct(product *domain.Product) error {
    if product.Name == "" {
        return errors.New("El nombre del producto no puede estar vacío")
    }

    if product.Quantity < 0 {
        return errors.New("La cantidad del producto no puede ser menor a 0")
    }

    if product.Price <= 0 {
        return errors.New("El precio del producto no puede ser menor o igual a 0")
    }

    if match, _ := regexp.MatchString("^\\d{2}/\\d{2}/\\d{4}$", product.Expiration); !match {
        return errors.New("La fecha de vencimiento debe tener el formato: dd/mm/yyyy")
    }

    return nil
}