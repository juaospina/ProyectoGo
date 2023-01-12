package internal

import (
	"fmt"
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
	
	json.Unmarshal(raw, &productStorage)
	fmt.Println(productStorage)
	return productStorage, nil
}



func WriteProductStorage(filePath string, productStorage ProductStorage) error {
	raw, err := json.MarshalIndent(productStorage, "", " ")
	if err != nil {
		log.Fatal(err)
		return err
	}
	err = ioutil.WriteFile(filePath, raw, 0644)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}
