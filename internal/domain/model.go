package domain

//Se declara la estructura modelo de nuestro producto
type Product struct {
    ID           int     `json:"id"`
    Name         string  `json:"name" binding:"required"`
    Quantity     int     `json:"quantity" binding:"required"`
    CodeValue    string  `json:"code_value" binding:"required"`
    IsPublished  bool    `json:"is_published"`
    Expiration   string  `json:"expiration" binding:"required"`
    Price        float64 `json:"price" binding:"required"`
}

type PatchRequest struct {
	Name         *string  `json:"name"`
	Quantity     *int     `json:"quantity"`
	Code_value   *string  `json:"code_value"`
	Is_published *bool    `json:"is_published"`
	Expiration   *string  `json:"expiration"`
	Price        *float64 `json:"price"`
}
