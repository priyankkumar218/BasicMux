package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Product struct {
	Id       string
	Name     string
	Quantity int
	Price    float64
}

var Products []Product

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: homePage")
	fmt.Fprint(w, "Welcome to the HomePage!")
}

func returnAllProducts(w http.ResponseWriter, r *http.Request) {
	log.Println("Endpoint Hit: returnAllProducts")
	json.NewEncoder(w).Encode(Products)
}

func getProduct(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL.Path)
	vars := mux.Vars(r)
	key := vars["id"]
	//key := r.URL.Path[len("/product/"):]
	for _, product := range Products {
		if string(product.Id) == key {
			json.NewEncoder(w).Encode(product)
		}
	}
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/products", returnAllProducts)
	myRouter.HandleFunc("/product/{id}", getProduct)
	myRouter.HandleFunc("/", homePage)
	err := http.ListenAndServe("", myRouter)
	if err != nil {
		panic(err)
	}
}

func main() {
	Products = []Product{
		Product{"1", "Chair", 100, 100.00},
		Product{"2", "Desk", 200, 200.00},
	}
	handleRequests()
}
