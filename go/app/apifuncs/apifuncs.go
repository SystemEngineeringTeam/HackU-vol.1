package apifuncs

import (
	"net/http"
	"net/url"
)

//TaskResponse は/tasksに対する処理をする
func TaskResponse(w http.ResponseWriter, r *http.Request) {

	//セキリティ設定
	w.Header().Set("Access-Control-Allow-Origin", "*")                       // Allow any access.
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE") // Allowed methods.
	
	q:=r.URL.Query()
	userToken:=q["userToken"]


	if r.Method == http.MethodGet {
	
	} else if r.Method == http.MethodPost {

	}

	w.WriteHeader(http.StatusOK)
}

//TaskSuccess は/tasks/successに対する処理(taskを達成した時の処理)
func TaskSuccess(w http.ResponseWriter, r *http.Request) {
	
	q:=r.URL.Query()
	userToken:=q["userToken"]
	userTOken:=q["usertaskID"]

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

	w.WriteHeader(http.StatusOK)
}

//UsersSignUp は/users/signupに対する処理
func UsersSignUp(w http.ResponseWriter, r *http.Request) {
	//セキリティ設定
	w.Header().Set("Access-Control-Allow-Origin", "*")                       // Allow any access.
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE") // Allowed methods.

	w.WriteHeader(http.StatusOK)
}




	
	


