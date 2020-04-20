package main

import (
	"database/sql"
	"log"
	"fmt"
	"net/http"
	"text/template"
	_ "github.com/go-sql-driver/mysql"
)

// We will create Employee struct that has following properties: Id, Name and City.

type Employee struct{
	Id int
	Name  string
	City string
}

// DB connection

func dbConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "newuser"
	dbPass := "password"
	dbName := "GOLANG"
	// Make connection using SQL OPEN
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil{
		panic(err.Error())
	}
	return db
}

// Variable for template in form folder

var tmpl = template.Must(template.ParseGlob("form/*"))

func Index(w http.ResponseWriter, r *http.Request)  {
	// Call DB connection
	db := dbConn()
	// Make select * db query
	selDB, err := db.Query("SELECT * FROM employee")
	if err != nil{
		panic(err.Error())
	}
	emp := Employee{}
	res := []Employee{}
	// Get elements and print
	for selDB.Next(){
		var id int
		var name, city string
		// Read the DB data
		err = selDB.Scan(&id, &name, &city)
		if err != nil{
			panic(err.Error())
		}
		emp.Id = id
		emp.Name = name
		emp.City = city
		res = append(res, emp)
	}
	fmt.Println(res)
	// Put data in template
	tmpl.ExecuteTemplate(w, "index", res)
	// Close DB conn
	defer db.Close()

	
}

func show(w http.ResponseWriter, r *http.Request)  {
	db := dbConn()
	nId := 1
	selDB, err := db.Query("SELECT * FROM Employee WHERE id=?", nId)
	if  err != nil{
		panic(err.Error())
	}

	emp := Employee{}
	for selDB.Next(){
		var id int
        var name, city string
        err = selDB.Scan(&id, &name, &city)
        if err != nil {
			panic(err.Error())
		}
		emp.Id = id
        emp.Name = name
        emp.City = city
	}
	tmpl.ExecuteTemplate(w, "Show", emp)
    defer db.Close()
}

func New(w http.ResponseWriter, r *http.Request)  {
	tmpl.ExecuteTemplate(w, "New", nil)
}


func Insert(w http.ResponseWriter, r *http.Request)  {
	db := dbConn()
	if r.Method == "POST"{
		name := r.FormValue("name")
		city := r.FormValue("city")
		insForm, err := db.Prepare("INSERT INTO Employee(name, city) VALUES(?,?)")
		if err != nil{
			panic(err.Error())
		}
		insForm.Exec(name, city)
		log.Println("INSERT: Name: " + name + " | City: " + city)
		defer db.Close()
		http.Redirect(w, r, "/", 301)
	}
	defer db.Close()
	http.Redirect(w, r, "/", 301)
}

func Update(w http.ResponseWriter, r *http.Request)  {
	db := dbConn()
	if r.Method == "POST"{
		name := r.FormValue("name")
		city := r.FormValue("city")
		id := r.FormValue("uid")
		insForm, err := db.Prepare("UPDATE Employee SET name=?, city=? WHERE id=?")
		if err != nil{
			panic(err.Error())
		}
		insForm.Exec(name, city, id)
		log.Println("UPDATE: Name: " + name + " | City: " + city)	
	}
	defer db.Close()
	http.Redirect(w, r, "/", 301)
}

func Delete(w http.ResponseWriter, r *http.Request)  {
	db := dbConn()
	emp := r.URL.Query().Get("id")
	delForm, err := db.Prepare("DELETE FROM Employee WHERE id=?")
    if err != nil {
        panic(err.Error())
	}
	delForm.Exec(emp)
	log.Println("DELETE")
	defer db.Close()
	http.Redirect(w, r, "/", 301)
}

func main() {
    log.Println("Server started on: http://localhost:8080")
    http.HandleFunc("/", Index)
    http.ListenAndServe(":8080", nil)
}
