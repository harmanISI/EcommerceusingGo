package main

import (
	"ecommerceProject/model"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/smtp"
	"strconv"
)

var tmpl *template.Template

type DataStruct struct {
	AllProducts []model.Product
	PageTitle string
	User string
}
var data DataStruct

func homePage(w http.ResponseWriter, r *http.Request) {
	tmpl = template.Must(template.ParseFiles("homepage.html"))

	model.AddProducts()
	allProducts := model.GetAllProducts()

	data.PageTitle = "HomePage"
	data.AllProducts = allProducts
	tmpl.Execute(w,data)
}

func prodPage(w http.ResponseWriter, r *http.Request){

	tmpl = template.Must(template.ParseFiles("product.html"))
	id, _ := strconv.Atoi(r.URL.Query().Get("id"))
	tempProduct := model.GetProduct(id)
	tmpl.Execute(w,tempProduct)
}

func logPage(w http.ResponseWriter,r * http.Request) {
	tmpl = template.Must(template.ParseFiles("login.html"))
	fmt.Println(r.Method)
	if r.Method == "POST"{
		username := r.FormValue("loginemail")
		password := r.FormValue("loginpass")
		fmt.Println(username,password)
		if model.CheckUser(username,password) != ""{
			/*tmpl = template.Must(template.ParseFiles("homepage.html"))*/
			data.User = model.CheckUser(username,password)
			homePage(w,r)
			/*tmpl.Execute(w,data)*/
		}
	}
	tmpl.Execute(w," ")
}

func payPage(w http.ResponseWriter,r * http.Request){
	tmpl = template.Must(template.ParseFiles("paypage.html"))
	id, _ := strconv.Atoi(r.URL.Query().Get("id"))
	tempProduct := model.GetProduct(id)
	tmpl.Execute(w,tempProduct)
}

func signup(w http.ResponseWriter, r *http.Request) {
	var nm string = ""
	var email string = ""
	var pass string = ""
	var str string
	tmpl = template.Must(template.ParseFiles("login.html"))
	if r.Method == "POST"{
		email = r.FormValue("email")
		pass = r.FormValue("pass")
		nm = r.FormValue("name")
		fmt.Println(email,pass,nm)
		if model.RegisterUser(email,pass,nm) {
			str = "User Created"
		} else {
			str = "User already Exists"
		}
	}
	tmpl.Execute(w,str)
}

func aboutus(w http.ResponseWriter, r *http.Request){
	tmpl = template.Must(template.ParseFiles("aboutus.html"))
	tmpl.Execute(w,"")
}

func orderplaced(w http.ResponseWriter, r *http.Request){
	id, _ := strconv.Atoi(r.URL.Query().Get("id"))
	tempProduct := model.GetProduct(id)
	str1 := "You bought "+ tempProduct.Name + " for $ "+ fmt.Sprintf("%f", tempProduct.Price)
	send(str1)
	homePage(w,r)
}

func send(body string) {
	fmt.Println("nfejvhbejfhcfj")
	from := "golangatecoomerce@gmail.com"
	pass := "Golang@123"
	to := "hjarewal7984@gmail.com"

	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: Hello there  Ecoomerce is there \n\n" +
		body

	err := smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", from, pass, "smtp.gmail.com"),
		from, []string{to}, []byte(msg))

	if err != nil {
		log.Printf("smtp error: %s", err)
		return
	}

	log.Print("sent, visit http://foobarbazz.mailinator.com")
}

func main() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/prod", prodPage)
	http.HandleFunc("/login", logPage)
	http.HandleFunc("/pay",payPage)
	http.HandleFunc("/signup",signup)
	http.HandleFunc("/about",aboutus)
	http.HandleFunc("/order",orderplaced)
	http.Handle("/img/", http.StripPrefix("/img", http.FileServer(http.Dir("./img"))))
	log.Println("Starting web server on port 8080")
	http.ListenAndServe(":8080", nil)
}