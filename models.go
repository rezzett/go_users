// Package dd:w
package main

import "database/sql"

type User struct {
	Id   uint
	Name string `json:"name"`
	Age  uint16 `json:"age"`
}

func AllUsers(DB *sql.DB) []User {
	var users []User
	res, err := DB.Query("SELECT `id`, `name`, `age` FROM `users`")
	if err != nil {
		panic(err)
	}

	for res.Next() {
		var user User
		err := res.Scan(&user.Id, &user.Name, &user.Age)
		if err != nil {
			panic(err)
		}
		//fmt.Printf("%d)User: %s - %d\n", user.Id, user.Name, user.Age)
		users = append(users, user)
	}
	return users
}
