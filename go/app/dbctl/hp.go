package dbctl

import (
	"log"
	"runtime"
	"time"
)

var layout = "2006-01-02 15:04:05"

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
	//usersテーブルのparam_idの変数
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

	taskIDs, err := callTaskIDsFromUserToken(token)
	if err != nil {
		return Hp{}, err
	}

	//updateした日時を取得
	rows, err = db.Query("select updated_datetime from user_parameters where id=?", temporaryUserID)
	defer rows.Close()
	var updateDate string
	for rows.Next() {
		rows.Scan(&updateDate)
	}

	//ダメージ計算処理
	currentHp, err := calculateCurrentHp(taskIDs, pastHp, updateDate, temporaryUserID)
	if err != nil {
		return Hp{}, err
	}

	//データベースのhpの更新
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

	//task_idsを取得
	rows, err = db.Query("select task_id from user_and_task_links where user_id=?", userID)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	taskIDs := make([]int, 0, 0)

	for rows.Next() {
		temporaryTaskID := 0
		rows.Scan(&temporaryTaskID)
		taskIDs = append(taskIDs, temporaryTaskID)
	}

	return taskIDs, err
}

func calculateCurrentHp(taskIDs []int, pastHp int, updateDate string, temporaryUserID int) (int, error) {

	var totalDamage int = 0
	var err error
	//現在の日付と時刻
	nowTime := time.Now()
	//現在の時刻を秒数に変換したもの
	//dataSecond := (data.Hour() * 3600) + (data.Minute() * 60) + data.Second()

	//フォーマットの整形time型の"2020-08-22 11:58:06 +0000 UTC"のような形式で表示される
	thenUpdateDate, err := time.Parse(layout, updateDate)
	if err != nil {
		return -1, err
	}

	//それぞれのtaskのダメージ計算
	for _, taskID := range taskIDs {

		//それぞれのタスクの重さ
		rowsWeightIDs, err := db.Query("select weight_id from tasks where id=?", taskID)
		if err != nil {
			return -1, err
		}
		defer rowsWeightIDs.Close()
		//タスク一つの重さ
		var WeightID int

		for rowsWeightIDs.Next() {
			WeightID = 0
			rowsWeightIDs.Scan(&WeightID)
		}

		//hpをアップデートした日(タスクを登録した時にもされる)と現在時刻の差
		diffUpdateDate := nowTime.Sub(thenUpdateDate)

		//float型をint型に変換したもの
		var intDiffUpdateDate int = int(diffUpdateDate.Seconds())

		totalDamage = totalDamage + intDiffUpdateDate*WeightID
	}

	//time型をstring型に変換したもの"2020-08-24 22:46:04"のような形になる
	stringUpdateNowTime := nowTime.Format(layout)
	//データベースのupdate_datetimeを現在時刻に変更
	_, err = db.Exec("update user_parameters set  updated_datetime=? where id=?", stringUpdateNowTime, temporaryUserID)

	currentHp := pastHp - totalDamage

	return currentHp, err

}

//RecoveryHp はタスクが達成されたときに20万回復する処理を行う
func RecoveryHp(token string) {

	//トークンからparamIDを取得
	rows, err := db.Query("select param_id from users where token=?", token)
	if err != nil {
		return
	}
	defer rows.Close()
	var temporaryUserID int
	for rows.Next() {
		temporaryUserID = 0
		rows.Scan(&temporaryUserID)
	}

	//現在のhpを取得
	rows, err = db.Query("select hp from user_parameters where id=?", temporaryUserID)
	if err != nil {
		return
	}
	defer rows.Close()
	var pastHp int
	for rows.Next() {
		temporaryUserID = 0
		rows.Scan(&pastHp)
	}

	recoveryAfterHp := pastHp + 200000

	//user_parametersの更新
	_, err = db.Exec("update user_parameters set hp=? where id=?", recoveryAfterHp, temporaryUserID)

}

//CountTaskIDUpdateTime はtaskIDの数を数えて時刻をアップデートする
func CountTaskIDUpdateTime(token string) error {

	taskIDs, err := callTaskIDsFromUserToken(token)
	if err != nil {
		return err
	}

	//usersのtokenからuser)parametersのidを取得
	rows, err := db.Query("select param_id from users where token=?", token)
	if err != nil {
		return err
	}
	//Next が呼び出されて false が返され，それ以上結果セットがない場合， rows は自動的に閉じられる
	defer rows.Close()
	//usersテーブルのparam_idの変数
	var temporaryUserID int
	for rows.Next() {
		temporaryUserID = 0
		rows.Scan(&temporaryUserID)
	}

	//タスクが0だったら
	if len(taskIDs) == 0 {
		nowTime := time.Now()
		//time型をstring型に変換したもの"2020-08-24 22:46:04"のような形になる
		stringUpdateNowTime := nowTime.Format(layout)
		//updated_datetimeの更新
		_, err = db.Exec("update user_parameters set  updated_datetime=? where id=?", stringUpdateNowTime, temporaryUserID)
	}

	return err
}
