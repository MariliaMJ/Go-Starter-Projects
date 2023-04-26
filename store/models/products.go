package models

import "github.com/mariliamorais/goStarterProjects/store/db"

type Product struct {
	Id         int
	Name       string
	Description  string
	Price      float64
	Quantity int
}

func GetAllProducts() []Product {
	db := db.ConnectDB()

	allproducts, err := db.Query("select * from Products")
	if err != nil {
		panic(err.Error())
	}

	p := Product{}
	products := []Product{}

	for allproducts.Next() {
		var id, quantity int
		var name, description string
		var price float64

		err = allproducts.Scan(&id, &name, &description, &price, &quantity)
		if err != nil {
			panic(err.Error())
		}

		p.Id = id
		p.Name = name
		p.Description = description
		p.Price = price
		p.Quantity = quantity

		products = append(products, p)
	}
	defer db.Close()
	return products
}
func InsertNewProduct(name, description string, price float64, quantity int) {
	db := db.ConnectDB()

	insertInDB, err := db.Prepare("insert into produtos(name, description, price, quantity) values($1, $2, $3, $4)")
	if err != nil {
		panic(err.Error())
	}

	insertInDB.Exec(name, description, price, quantity)
	defer db.Close()

}

func DeleteProduct(id string) {
	db := db.ConnectDB()

	deleteProduct, err := db.Prepare("delete from produtos where id=$1")
	if err != nil {
		panic(err.Error())
	}

	deleteProduct.Exec(id)
	defer db.Close()

}

func UpdateProductById(id string) Product {
	db := db.ConectaComBancoDeDados()

	dbProduct, err := db.Query("select * from produtos where id=$1", id)
	if err != nil {
		panic(err.Error())
	}

	productToUpdate := Product{}

	for dbProduct.Next() {
		var id, quantity int
		var name, description string
		var price float64

		err = dbProduct.Scan(&id, &name, &description, &price, &quantity)
		if err != nil {
			panic(err.Error())
		}
		productToUpdate.Id = id
		productToUpdate.Name = name
		productToUpdate.Description = description
		productToUpdate.Price = price
		productToUpdate.Quantity = quantity
	}
	defer db.Close()
	return productToUpdate
}

func UpdateProduct(id int, name, description string, price float64, quantity int) {
	db := db.ConnectDB()

	UpdateProduct, err := db.Prepare("update produtos set name=$1, description=$2, price=$3, quantity=$4 where id=$5")
	if err != nil {
		panic(err.Error())
	}
	UpdateProduct.Exec(name, description, price, quantity, id)
	defer db.Close()
}