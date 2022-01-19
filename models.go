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
	if err != nil {
		panic(err)
	}

	for res.Next() {
		var user User
		err := res.Scan(&user.Id, &user.Name, &user.Age)
		if err != nil {
			panic(err)
		}
		users = append(users, user)
	}
	return users
}

func InsertUser(db sql.DB, name string, age uint) {
	result, err := db.Query("INSERT INTO `users` (`name`, `age`) VALUES('Issac', 56)")
	if err != nil {
		panic(err)
	}
	defer result.Close()
}
