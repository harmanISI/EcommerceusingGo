package model

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"log"
)

var allProducts []Product
var db *sql.DB

type Product struct{
	Id int
	Name string
	Price float32
	ImgUrl string
}

func AddProducts(){
	allProducts = nil
	getConnection()
	res, err := db.Query("SELECT * FROM product")

	if err != nil {
		log.Fatal(err)
	}

	var tempProduct Product
	for res.Next() {
		res.Scan(&tempProduct.Id, &tempProduct.Name, &tempProduct.Price, &tempProduct.ImgUrl)
		allProducts = append(allProducts,tempProduct)
	}

}

func GetAllProducts() []Product {
	return allProducts
}

func GetProduct(id int) Product {
	for i:=0; i<len(allProducts); i++{
		if id == allProducts[i].Id{
			return allProducts[i]
		}
	}
	return Product{}
}

func getConnection() sql.DB{
	cfg := mysql.Config{
		User:   "root",
		Passwd: "passpasspass@123",
		Net:    "tcp",
		Addr:   "localhost:3306",
		DBName: "ecommerce",
	}
	// Get a database handle.
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}

	return *db
}
func CloseDatabase(db *sql.DB) {
	defer db.Close()
}