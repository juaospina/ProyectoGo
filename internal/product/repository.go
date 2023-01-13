package product

import (
	"errors"
	"ProyectoGo/internal"
	"ProyectoGo/internal/domain"
	"fmt"
)

var (
	ProductNotFound = errors.New("Producto no encontrado")
)
type Repository struct {
	productStorage *internal.ProductStorage
	storagePath    string
}

func NewRepository(storagePath string) (*Repository, error) {
	productStorage, err := internal.ReadProductStorage(storagePath)
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
			updatedProduct.ID = id
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

func (r *Repository) UpdatePATCH(id int, request domain.PatchRequest) (*domain.Product, error) {
	var productPatch domain.Product

	PatchOK := false

	for i := range r.productStorage.Products {
		if (r.productStorage.Products)[i].ID == id {
			if request.Name != nil {
				(r.productStorage.Products)[i].Name = *request.Name
			}
			if request.Quantity != nil {
				(r.productStorage.Products)[i].Quantity = *request.Quantity
			}
			if request.Code_value != nil {
				(r.productStorage.Products)[i].CodeValue = *request.Code_value
			}
			if request.Is_published != nil {
				(r.productStorage.Products)[i].IsPublished = *request.Is_published
			}
			if request.Expiration != nil {
				(r.productStorage.Products)[i].Expiration = *request.Expiration
			}
			if request.Price != nil {
				(r.productStorage.Products)[i].Price = *request.Price
			}

			PatchOK = true
			productPatch = (r.productStorage.Products)[i]
			break
		}
	}

	if !PatchOK {
		return nil, fmt.Errorf("%w. %s", ProductNotFound, "El producto no existe")
	}

	err := internal.WriteProductStorage(r.storagePath, *r.productStorage)
			if err != nil {
				return nil, err
			}
			return &productPatch, nil
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
