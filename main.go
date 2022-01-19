package main

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type DBUser struct {
	Id   uint
	Name string `json:"name"`
	Age  uint16 `json:"age"`
}

func handleRequsts() {
	http.HandleFunc("/", HomePage)
	http.HandleFunc("/contacts/", ContactsPage)
	http.ListenAndServe(":8081", nil)
}

func main() {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/simple_site_go")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	/* insert new user
	 *result, err := db.Query("INSERT INTO `users` (`name`, `age`) VALUES('Issac', 56)")
	 *if err != nil {
	 *    panic(err)
	 *}
	 *defer result.Close()
	 */

	res, err := db.Query("SELECT `id`, `name`, `age` FROM `users`")
	if err != nil {
		panic(err)
	}

	for res.Next() {
		var user DBUser
		err := res.Scan(&user.Id, &user.Name, &user.Age)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%d)User: %s - %d\n", user.Id, user.Name, user.Age)
	}

	handleRequsts()

}
