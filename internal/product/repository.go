package product

import (
	"errors"
	"ProyectoGo/internal"
	"ProyectoGo/internal/domain"
	"fmt"
)

type Repository struct {
	productStorage *internal.ProductStorage
	storagePath    string
}

func NewRepository(storagePath string) (*Repository, error) {
	productStorage, err := internal.ReadProductStorage(storagePath)
	fmt.Println(productStorage)
	if err != nil {
		return nil, err
	}
	return &Repository{productStorage: &productStorage, storagePath: storagePath}, nil
}


func (r *Repository) GetAll() ([]domain.Product, error) {
	return r.productStorage.Products, nil
}

func (r *Repository) GetByID(id int) (*domain.Product, error) {
	for _, product := range r.productStorage.Products {
		if product.ID == id {
			return &product, nil
		}
	}
	return nil, fmt.Errorf("Product with ID %d not found", id)
}

func (r *Repository) Create(newProduct domain.Product) (*domain.Product, error) {
	newProduct.ID = len(r.productStorage.Products) + 1
	r.productStorage.Products = append(r.productStorage.Products, newProduct)
	//Valida code_value unico
	for _, p := range r.productStorage.Products {
		if newProduct.CodeValue == p.CodeValue && newProduct.ID != p.ID {
			return nil, errors.New("El valor del codigo ya esta en uso")
		}
	}
	
	err := internal.WriteProductStorage(r.storagePath, *r.productStorage)
	if err != nil {
		return nil, err
	}
	return &newProduct, nil
}

func (r *Repository) Update(id int, updatedProduct domain.Product) (*domain.Product, error) {
	for i, product := range r.productStorage.Products {
		if product.ID == id {
			r.productStorage.Products[i] = updatedProduct
			
			err := internal.WriteProductStorage(r.storagePath, *r.productStorage)
			if err != nil {
				return nil, err
			}
			return &updatedProduct, nil
		}
	}
	return nil, fmt.Errorf("Product with ID %d not found", id)
}

func (r *Repository) Delete(id int) error {
	for i, product := range r.productStorage.Products {
		if product.ID == id {
			r.productStorage.Products = append(r.productStorage.Products[:i], r.productStorage.Products[i+1:]...)
			err := internal.WriteProductStorage(r.storagePath, *r.productStorage)
			if err != nil {
				return err
			}
			return nil
		}
	}
	return fmt.Errorf("Product with ID %d not found", id)
}
