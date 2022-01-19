package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type User struct {
	Name             string
	Age              uint16
	Money            int16
	Grades, Happines float64
	Hobbies          []string
}

func (u *User) Info() string {
	return fmt.Sprintf("Name: %s\nAge: %d\nMoney: %d", u.Name, u.Age, u.Money)
}

func HomePage(w http.ResponseWriter, r *http.Request) {
	alyx := User{"Alyx", 25, 280, 5.0, 0.9, []string{"Travel", "Sience", "Sport"}}
	tmpl, _ := template.ParseFiles("./templates/home.html")
	tmpl.Execute(w, alyx)
}

func ContactsPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h2>ContactsPage</h2>")
}
