package apifuncs

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"set1.ie.aitech.ac.jp/HackU_vol_1/dbctl"
	
)

//TaskResponse は/tasksに対する処理をする
func TaskResponse(w http.ResponseWriter, r *http.Request) {

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

	//getメソッドの時
	if r.Method == http.MethodGet {

		//タスクを取得
		tasks, err := dbctl.CallTasks(userToken)

		if err != nil {
			w.WriteHeader(http.StatusServiceUnavailable)
			log.Fatal(err)
			return
		}

		//バイト型のjsonで受け取る
		jsonBytes, err := json.Marshal(tasks)
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

	} else if r.Method == http.MethodPost {

		//body読み込み
		jsonBytes, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusServiceUnavailable)
			fmt.Println("io error")
			return
		}

		//構造体の初期化
		data := dbctl.Task{}

		//taskの構造体にbodyの値を入れる
		if err := json.Unmarshal(jsonBytes, &data); err != nil {
			w.WriteHeader(http.StatusServiceUnavailable)
			fmt.Println("JSON Unmarshal error:", err)
			return
		}

		//taskの登録
		taskID, err := dbctl.RegisterNewTask(userToken, data)
		if err != nil {
			w.WriteHeader(http.StatusServiceUnavailable)
			fmt.Println("database error", err)
			return
		}

		w.WriteHeader(http.StatusOK)
		r.Header.Set("Content-Type", "application/json")

		//クライアントに返す
		fmt.Fprintln(w, taskID)

	}
}


//TaskSuccess は/tasks/successに対する処理(taskを達成した時の処理)
func TaskSuccess(w http.ResponseWriter, r *http.Request) {

	//セキリティ設定
	w.Header().Set("Access-Control-Allow-Origin", "*")                       // Allow any access.
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE") // Allowed methods.
	w.Header().Set("Access-Control-Allow-Headers", "*")

	//クエリのパラメータの取得
	q := r.URL.Query()

	var userToken string
	var stringTaskID string

	if len(q["userToken"]) > 0 {
		userToken = q["userToken"][0]
	} else {
		fmt.Println("out of index")
		return
	}

	if len(q["taskID"]) > 0 {
		stringTaskID = q["taskID"][0]
	} else {
		fmt.Println("out of index")
		return
	}
			
	if r.Method == http.MethodPost {

		//数値に変換
		numberTaskID, err := strconv.Atoi(stringTaskID)
		if err != nil {
			w.WriteHeader(http.StatusServiceUnavailable)
			fmt.Println("changeNumber error", err)
			return
		}

		if err := dbctl.TaskAchieveFlagChangeToTrue(userToken, numberTaskID); err != nil {
			w.WriteHeader(http.StatusServiceUnavailable)
			fmt.Println("database error", err)						
			return
		}

	}

	w.WriteHeader(http.StatusOK)
}

//TaskDifficulty は/tasks/weightsに対する処理
func TaskDifficulty(w http.ResponseWriter, r *http.Request) {
	//セキリティ設定
	w.Header().Set("Access-Control-Allow-Origin", "*")                       // Allow any access.
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE") // Allowed methods.
	w.Header().Set("Access-Control-Allow-Headers", "*")

	//難易度をデータベースからもらう
	weight, err := dbctl.CallWeightsList()
	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		fmt.Println("database err", err)
		return
	}

	w.WriteHeader(http.StatusOK)

	r.Header.Set("Content-Type", "application/json")

	fmt.Fprintf(w, weight)
}
