package controllers

import (
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/mariliamorais/goStarterProjects/store/models"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	allProducts := models.GetAllProducts()
	temp.ExecuteTemplate(w, "Index", allProducts)
}

func New(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name := r.FormValue("name")
		description := r.FormValue("description")
		price := r.FormValue("price")
		quantity := r.FormValue("quantity")

		priceToFloat, err := strconv.ParseFloat(price, 64)
		if err != nil {
			log.Println("Error converting price:", err)
		}

		quantityToInt, err := strconv.Atoi(quantity)
		if err != nil {
			log.Println("Error converting quantity:", err)
		}

		models.InsertNewProduct(name, description, priceToFloat, quantityToInt)
	}
	http.Redirect(w, r, "/", 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	productID := r.URL.Query().Get("id")
	models.DeleteProduct(productID)
	http.Redirect(w, r, "/", 301)
}

func UpdateById(w http.ResponseWriter, r *http.Request) {
	idDoProduto := r.URL.Query().Get("id")
	produto := models.UpdateProductById(idDoProduto)
	temp.ExecuteTemplate(w, "Edit", produto)
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		name := r.FormValue("nome")
		description := r.FormValue("descricao")
		price := r.FormValue("preco")
		quantity := r.FormValue("quantidade")

		idToInt, err := strconv.Atoi(id)
		if err != nil {
			log.Println("Error converting ID to int:", err)
		}

		priceToFloat, err := strconv.ParseFloat(price, 64)
		if err != nil {
			log.Println("Error converting price to float64:", err)
		}

		quantityToInt, err := strconv.Atoi(quantity)
		if err != nil {
			log.Println("Error converting quantity to int:", err)
		}

		models.UpdateProduct(idToInt, name, description, priceToFloat, quantityToInt)
	}
	http.Redirect(w, r, "/", 301)
}