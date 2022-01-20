package main

import (
	"html/template"
	"net/http"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	Users := AllUsers(db)
	files := []string{"./templates/home.html", "./templates/base.html"}
	tmpl, err := template.ParseFiles(files...)
	CheckErr(err)
	tmpl.Execute(w, Users)
}

func ContactsPage(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "<h2>ContactsPage</h2>")
	files := []string{"./templates/contacts.html", "./templates/base.html"}
	t, err := template.ParseFiles(files...)
	CheckErr(err)
	Title := "contacts"
	t.Execute(w, Title)
}
