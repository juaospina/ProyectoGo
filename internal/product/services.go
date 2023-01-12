package product

import (
	"ProyectoGo/internal/domain"
	"ProyectoGo/pkg/validator"
)

type Services struct {
	productRepo *Repository
}

func NewServices(storagePath string) (*Services, error) {
	repo, err := NewRepository(storagePath)
	if err != nil {
		return nil, err
	}
	return &Services{productRepo: repo}, nil
}

func (s *Services) GetAll() ([]domain.Product, error) {
	return s.productRepo.GetAll()
}

func (s *Services) GetByID(id int) (*domain.Product, error) {
	return s.productRepo.GetByID(id)
}

func (s *Services) Create(newProduct domain.Product) (*domain.Product, error) {
	if err := validator.ValidateProduct(&newProduct); err != nil {
		return nil, err
	}
	return s.productRepo.Create(newProduct)
}

func (s *Services) Update(id int, updatedProduct domain.Product) (*domain.Product, error) {
	if err := validator.ValidateProduct(&updatedProduct); err != nil {
		return nil, err
	}
	return s.productRepo.Update(id, updatedProduct)
}

func (s *Services) Delete(id int) error {
	return s.productRepo.Delete(id)
}

func (s *Services) Search(priceGt float64) ([]domain.Product, error) {
    products, err := s.productRepo.GetAll()
    if err != nil {
        return nil, err
    }
    filteredProducts := make([]domain.Product, 0)
    for _, product := range products {
        if product.Price > priceGt {
            filteredProducts = append(filteredProducts, product)
        }
    }
    return filteredProducts, nil
}
