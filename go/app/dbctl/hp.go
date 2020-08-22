package dbctl

import (
	"log"
	"runtime"
)

//Hp はデータベースから取得した値を扱うための構造体
type Hp struct {
	CurrentHp int `json:"hp"`
	MaxHp     int `json:"maxHp"`
}

//CallHpFromUserToken は現在のhpとmaxhpを構造体で返す関数
func CallHpFromUserToken(token string) (Hp, error) {

	//rows*型 //tokenは文字列かinterface型
	//usersのtokenからuser)parametersのidを取得
	rows, err := db.Query("select param_id from users where token=?", token)

	if err != nil {
		return Hp{}, err
	}
	//Next が呼び出されて false が返され，それ以上結果セットがない場合， rows は自動的に閉じられる
	defer rows.Close()

	var temporaryUserID int

	for rows.Next() {
		temporaryUserID = 0
		rows.Scan(&temporaryUserID)
	}

	//user_parametersのidからuser_parametersテーブルのhpを取得
	rows, err = db.Query("select hp from user_parameters where id=?", temporaryUserID)

	if err != nil {
		return Hp{}, err
	}

	defer rows.Close()

	//明示的な型宣言
	var pastHp int

	for rows.Next() {
		pastHp = 0
		rows.Scan(&pastHp)
	}

	tasksID, err := callTaskIDsFromUserToken(token)

	if err != nil {
		return Hp{}, err
	}

	currentHp := calculateCurrentHp(tasksID, pastHp)

	//データベースの更新
	_, err = db.Exec("update user_parameters set hp=? where id=?", currentHp, temporaryUserID)

	if err != nil {
		pc, file, line, _ := runtime.Caller(0)
		f := runtime.FuncForPC(pc)
		log.Printf(errFormat, err, f.Name(), file, line)
		return Hp{}, err
	}

	//構造体の初期化
	hp := Hp{}

	hp = Hp{CurrentHp: currentHp, MaxHp: 1000000}

	return hp, nil
}

func callTaskIDsFromUserToken(token string) ([]int, error) {

	//usersのidを取得
	rows, err := db.Query("select id from users where token=?", token)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var userID int

	for rows.Next() {
		userID = 0
		rows.Scan(&userID)
	}

	//task_idを取得
	rows, err = db.Query("select task_id from user_and_task_links where user_id=?", userID)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var taskID []int

	for rows.Next() {
		rows.Scan(&taskID)
	}

	return taskID, err

}

func calculateCurrentHp(tasksID []int, pastHp int) int {

	/* t := time.Now()
	fmt.Println(t)
	fmt.Println(t.Year())
	fmt.Println(t.Month())
	fmt.Println(t.Day())
	fmt.Println(t.Hour())
	fmt.Println(t.Minute())
	fmt.Println(t.Weekday()) */

	numTask := len(tasksID)

	currentHp := pastHp - (numTask * 100000)

	return currentHp

}
