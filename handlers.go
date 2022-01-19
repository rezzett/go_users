package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	Users := AllUsers(db)
	tmpl, err := template.ParseFiles("./templates/home.html")
	if err != nil {
		panic(err)
	}

	tmpl.Execute(w, Users)
}

func ContactsPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h2>ContactsPage</h2>")
}
