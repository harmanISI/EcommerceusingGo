package model

import (
	"fmt"
	"log"
)

type User struct {
	Uid int;
	Name string;
	Username string;
	Password string;
}

func CheckUser(uName string, pass string) string{
	getConnection()
	res, err := db.Query("SELECT * FROM users")

	if err != nil {
		log.Fatal(err)
	}
	var tempUser User
	for res.Next() {
		res.Scan(&tempUser.Uid,&tempUser.Name,&tempUser.Username,&tempUser.Password)
		if tempUser.Username == uName && tempUser.Password == pass{
			return tempUser.Name
		}
	}

	return ""
}

func RegisterUser(email string,pass string,name string ) bool{
	db :=getConnection()

	res, err := db.Query("SELECT * FROM users")

	if err != nil {
		log.Fatal(err)
	}
	var tempUser User
	for res.Next() {
		res.Scan(&tempUser.Uid,&tempUser.Name,&tempUser.Username,&tempUser.Password)
		if tempUser.Username == email {
			fmt.Println(" user already exist")
			return false
		}
	}

	_, err = db.Query("INSERT INTO ecommerce.users(fname,email,pass) VALUES (?,?,?)", name, email, pass)

	if err != nil {
		panic(err.Error())
	}
	CloseDatabase(&db)
	return true
}