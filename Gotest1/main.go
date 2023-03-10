package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// geekole.com
type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type error interface {
	Error() string
}

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	db, err := sql.Open("mysql", "root:123456@tcp(localhost:49153)/mytest?parseTime=true")
	if err != nil {
		fmt.Println(err.Error())
	}
	defer db.Close()
	results, err := db.Query("SELECT id, username, password, email FROM users")
	if err != nil {
		fmt.Println(err)
	}
	for results.Next() {
		var user User
		err = results.Scan(&user.ID, &user.Username, &user.Password, &user.Email)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(user.Username + " " + user.Password + " " + user.Email)
	}

}
