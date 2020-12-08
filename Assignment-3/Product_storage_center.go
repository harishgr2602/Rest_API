package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//Product struct to add product details
type Product struct {
	ProductID    string        `json:"productid"`
	Name         string        `json:"name"`
	Price        string        `json:"price"`
	NoOfStock    string        `json:"noofstock"`
	Manufacturer *Manufacturer `json:"manufacturer"`
}

//Manufacturer struct to add manufacturer details
type Manufacturer struct {
	CompanyName  string `json:"companyname"`
	CustomerCare string `json:"customercare"`
	Location     string `json:"location"`
}

var products []Product

func main() {
	r := mux.NewRouter()
	products = append(products, Product{ProductID: "1", Name: "MarieGoldBiscuit", Price: "10", NoOfStock: "40", Manufacturer: &Manufacturer{CompanyName: "Britania", CustomerCare: "1800-9875-1838", Location: "Banaglore"}})
	products = append(products, Product{ProductID: "2", Name: "Bovonto", Price: "32", NoOfStock: "60", Manufacturer: &Manufacturer{CompanyName: "KaliMark", CustomerCare: "04761-252363", Location: "Thanjavur"}})
	products = append(products, Product{ProductID: "3", Name: "Coconut Hair Oil", Price: "45", NoOfStock: "11", Manufacturer: &Manufacturer{CompanyName: "Parachute", CustomerCare: "1800-4095-2122", Location: "Trivandrum"}})
	products = append(products, Product{ProductID: "4", Name: "kurkurre", Price: "5", NoOfStock: "90", Manufacturer: &Manufacturer{CompanyName: "ITC", CustomerCare: "1800-9875-9954", Location: "Chennai"}})
	products = append(products, Product{ProductID: "5", Name: "Achi Pickle", Price: "39", NoOfStock: "19", Manufacturer: &Manufacturer{CompanyName: "Achi", CustomerCare: "1900-9125-1747", Location: "Chennai"}})
	products = append(products, Product{ProductID: "6", Name: "KadalaiMittai", Price: "10", NoOfStock: "52", Manufacturer: &Manufacturer{CompanyName: "ManiMark", CustomerCare: "0476-251478", Location: "Vellore"}})

	r.HandleFunc("/products", getProducts).Methods("GET")
	r.HandleFunc("/products/{productid}", getProduct).Methods("GET")
	r.HandleFunc("/products/{productid}", postProduct).Methods("POST")
	r.HandleFunc("/products/{productid}", deleteProduct).Methods("DELETE")
	r.HandleFunc("/products/{productid}", updateProduct).Methods("PUT")

	log.Fatal(http.ListenAndServe(":5050", r))
}

func getProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}
func getProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	p := mux.Vars(r)
	for _, v := range products {
		if v.ProductID == p["productid"] {
			json.NewEncoder(w).Encode(v)
			return
		}
	}
	json.NewEncoder(w).Encode(&Product{})
}
func postProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var prod Product
	p := mux.Vars(r)
	_ = json.NewDecoder(r.Body).Decode(&prod)
	prod.ProductID = p["productid"]
	products = append(products, prod)
	json.NewEncoder(w).Encode(products)
}
func deleteProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	p := mux.Vars(r)
	for k, v := range products {
		if v.ProductID == p["productid"] {
			products = append(products[:k], products[(k+1):]...)
			break
		}
	}
	json.NewEncoder(w).Encode(products)
}

func updateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	p := mux.Vars(r)
	for k, v := range products {
		if v.ProductID == p["productid"] {
			products = append(products[:k], products[k+1:]...)
			var pro Product
			_ = json.NewDecoder(r.Body).Decode(&pro)
			pro.ProductID = p["productid"]
			products = append(products, pro)
			json.NewEncoder(w).Encode(products)
			return
		}
	}
	json.NewEncoder(w).Encode(products)
}
