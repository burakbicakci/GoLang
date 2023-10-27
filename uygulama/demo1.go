package uygulama

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Product struct { //json verilerini almak için struct tanımladık
	Id          int     `json:"id"`
	ProductName string  `json:"productsName"`
	CategoryId  int     `json:"categoryId"`
	UnitPrice   float64 `json:"unitPrice"`
}

type Categorytype struct {
	Id           int    `json:"id"`
	CategoryName string `json:"categoryName"`
}

func GetAllProducts() {
	response, err := http.Get("http://localhost:3000/products") // get isteği
	if err != nil {                                             // hata varsa
		fmt.Println(err)
	}
	defer response.Body.Close()                    // işlem bitince kapat
	bodyBtytes, _ := ioutil.ReadAll(response.Body) // okuma işlemi

	var products []Product
	json.Unmarshal(bodyBtytes, &products) // json verisini products dizisine aktar
	fmt.Println(products)
}

func AddProducts() {
	product := Product{Id: 6, ProductName: "Laptop", CategoryId: 1, UnitPrice: 5000}
	productJson, _ := json.Marshal(product) // product nesnesini json verisine çevir

	response, err := http.Post("http://localhost:3000/products", "application/json;charset=utf-8", bytes.NewBuffer(productJson)) // post isteği

	if err != nil { // hata varsa
		fmt.Println(err)
	}

	defer response.Body.Close() // işlem bitince kapat

	fmt.Println("Kayıt eklendi")
}
