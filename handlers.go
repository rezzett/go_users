package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	users := AllUsers(db)
	tmpl, _ := template.ParseFiles("./templates/home.html")

	tmpl.Execute(w, users)
}

func ContactsPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h2>ContactsPage</h2>")
}
