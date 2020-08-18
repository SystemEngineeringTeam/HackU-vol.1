package apifuncs

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"set1.ie.aitech.ac.jp/HackU_vol_1/dbctl"
)

//TaskResponse は/tasksに対する処理をする
func TaskResponse(w http.ResponseWriter, r *http.Request) {

	//セキュリティ設定
	w.Header().Set("Access-Control-Allow-Origin", "*")                       // Allow any access.
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE") // Allowed methods.
	w.Header().Set("Access-Control-Allow-Headers","*")

	q := r.URL.Query()
	userToken := q["userToken"][0]

	//getメソッドの時
	if r.Method == http.MethodGet {

		//タスクを取得
		tasks, err := dbctl.CallTasks(userToken)

		if err != nil {
			w.WriteHeader(http.StatusServiceUnavailable)
			log.Fatal(err)
		}

		//バイト型のjsonで受け取る
		jsonBytes, err := json.Marshal(tasks)
		if err != nil {
			w.WriteHeader(http.StatusServiceUnavailable)
			log.Fatal(err)
		}

		jsonString := string(jsonBytes)

		// httpステータスコードを返す<-New
		w.WriteHeader(http.StatusOK)
		// JSONを返す
		fmt.Fprintln(w, jsonString)

	} else if r.Method == http.MethodPost {

	}

	w.WriteHeader(http.StatusOK)
}

//TaskSuccess は/tasks/successに対する処理(taskを達成した時の処理)
func TaskSuccess(w http.ResponseWriter, r *http.Request) {

/* 	q := r.URL.Query()
	taskID := q["userToken"][0]
	userToken := q["usertaskID"][0] */
	
	//セキリティ設定
	w.Header().Set("Access-Control-Allow-Origin", "*")                       // Allow any access.
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE") // Allowed methods.

	if r.Method == http.MethodPost {

	}

	w.WriteHeader(http.StatusOK)
}

//UsersLogin は/users/loginに対する処理
func UsersLogin(w http.ResponseWriter, r *http.Request) {
	//セキリティ設定
	w.Header().Set("Access-Control-Allow-Origin", "*")                       // Allow any access.
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE") // Allowed methods.
	w.Header().Set("Access-Control-Allow-Headers","*")

	w.WriteHeader(http.StatusOK)
}

//UsersSignUp は/users/signupに対する処理
func UsersSignUp(w http.ResponseWriter, r *http.Request) {
	//セキリティ設定
	w.Header().Set("Access-Control-Allow-Origin", "*")                       // Allow any access.
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE") // Allowed methods.
	w.Header().Set("Access-Control-Allow-Headers","*")
	
	w.WriteHeader(http.StatusOK)
}
