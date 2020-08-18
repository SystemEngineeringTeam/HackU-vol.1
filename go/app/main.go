package main

import (
	"net/http"
		
	_ "github.com/go-sql-driver/mysql"
	"set1.ie.aitech.ac.jp/HackU_vol_1/apifuncs"
)

func main() {
	http.HandleFunc("/tasks", apifuncs.TaskResponse)
	http.HandleFunc("tasks/success", apifuncs.TaskSuccess)
	http.HandleFunc("/users/login", apifuncs.UsersLogin)
	http.HandleFunc("/users/signup", apifuncs.UsersSignUp)

	http.ListenAndServe(":80", nil)
}
