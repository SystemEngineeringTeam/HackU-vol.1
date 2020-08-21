package apifuncs

import (
	"fmt"
	"net/http"
)

//HpResponse は/hpに対する処理をする
func HpResponse(w http.ResponseWriter, r *http.Request) {

	//セキュリティ設定
	w.Header().Set("Access-Control-Allow-Origin", "*")                       // Allow any access.
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE") // Allowed methods.
	w.Header().Set("Access-Control-Allow-Headers", "*")

	q := r.URL.Query()
	var userToken string

	if len(q["userToken"]) > 0 {
		userToken = q["userToken"][0]
	} else {
		fmt.Println("out of index")
	}


	if r.Method==http.MethodGet{
		
		
	}







}
