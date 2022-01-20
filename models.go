package main

import "database/sql"

type User struct {
	Id   uint
	Name string `json:"name"`
	Age  uint16 `json:"age"`
}

func AllUsers(db *sql.DB) []User {
	var users []User
	res, err := db.Query("SELECT `id`, `name`, `age` FROM `users`")
	CheckErr(err)

	for res.Next() {
		var user User
		err := res.Scan(&user.Id, &user.Name, &user.Age)
		CheckErr(err)
		users = append(users, user)
	}
	return users
}

func InsertUser(db sql.DB, name string, age uint) {
	result, err := db.Query("INSERT INTO `users` (`name`, `age`) VALUES('Issac', 56)")
	CheckErr(err)
	defer result.Close()
}
