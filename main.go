package main

import (
	"database/sql"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}

var db *sql.DB

func handleRequsts() {
	http.HandleFunc("/", HomePage)
	http.HandleFunc("/contacts/", ContactsPage)
	http.ListenAndServe(":8081", nil)
}

func main() {
	var err error
	db, err = sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/simple_site_go")
	CheckErr(err)
	defer db.Close()

	handleRequsts()
}
