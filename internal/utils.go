package internal

import (
	"os"
	"io/ioutil"
	"log"
	"encoding/json"
	"ProyectoGo/internal/domain"
)

type ProductStorage struct {
	Products []domain.Product `json:"products"`
}


func ReadProductStorage(filePath string) (ProductStorage, error) {
	var productStorage ProductStorage
	raw, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
		return productStorage, err
	}
	
	json.Unmarshal(raw, &(productStorage.Products))
	return productStorage, nil
}


func WriteProductStorage(filePath string, productStorage ProductStorage) error {
	raw, err := json.Marshal(&productStorage.Products)
	if err != nil {
		log.Fatal(err)
		return err
	}
	
	f, err := os.Create("products.json")
    if err != nil {
        panic(err)
    }
    defer f.Close()

    // Escribir el contenido en formato de array en el archivo
    json.NewEncoder(f).Encode(&productStorage.Products)

	err = ioutil.WriteFile(filePath, raw, 0644)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil

}