package apifuncs

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"set1.ie.aitech.ac.jp/HackU_vol_1/dbctl"
)

//UsersLogin は/users/loginに対する処理
func UsersLogin(w http.ResponseWriter, r *http.Request) {
	//セキリティ設定
	w.Header().Set("Access-Control-Allow-Origin", "*")                       // Allow any access.
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE") // Allowed methods.
	w.Header().Set("Access-Control-Allow-Headers", "*")

	jsonBytes, err := ioutil.ReadAll(r.Body)

	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		log.Println("io error")		
		return
	}

	//構造体の初期化
	data := dbctl.User{}

	if err := json.Unmarshal(jsonBytes, &data); err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		fmt.Println("JSON Unmarshal error:", err)
		return
	}

	//データベースからトークンを取得(string型)
	data, err = dbctl.Login(data.Email, data.Pass)
	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		log.Println("database error")
		return
	}

	jsonString := `{"name":"` + data.Name + `","token":"` + data.Token + `"}`

	w.WriteHeader(http.StatusOK)
	r.Header.Set("Content-Type", "application/json")
	
	//クライアントに渡す
	fmt.Fprintf(w, jsonString)
}

//UsersSignUp は/users/signupに対する処理
func UsersSignUp(w http.ResponseWriter, r *http.Request) {
	//セキュリティ設定
	w.Header().Set("Access-Control-Allow-Origin", "*")                       // Allow any access.
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE") // Allowed methods.
	w.Header().Set("Access-Control-Allow-Headers", "*")

	jsonBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		log.Println("io error")
		return
	}

	/*
		content := string(jsonBytes)
		log.Fatalln(content) */

	//構造体の初期化
	data := dbctl.User{}

	if err := json.Unmarshal(jsonBytes, &data); err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		fmt.Println("JSON Unmarshal error:", err)
		return		
	}

	//ユーザ登録を行う

	if err := dbctl.RegisterNewUser(data); err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)		
		log.Println("database error")
		return
	}

	w.WriteHeader(http.StatusOK)
}
