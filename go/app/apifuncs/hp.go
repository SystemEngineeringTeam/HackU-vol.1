package apifuncs

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"set1.ie.aitech.ac.jp/HackU_vol_1/dbctl"
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

	if r.Method == http.MethodGet {

		hp, err := dbctl.CallHpFromUserToken(userToken)
		if err != nil {
			w.WriteHeader(http.StatusServiceUnavailable)
			fmt.Println("database error", err)
			return
		}

		//バイト型のjsonで受け取る
		jsonBytes, err := json.Marshal(hp)
		if err != nil {
			w.WriteHeader(http.StatusServiceUnavailable)
			log.Fatal(err)
			return
		}

		jsonString := string(jsonBytes)

		// httpステータスコードを返す<-New
		w.WriteHeader(http.StatusOK)
		r.Header.Set("Content-Type", "application/json")
		// JSONを返す
		fmt.Fprintln(w, jsonString)
			
	}

}
