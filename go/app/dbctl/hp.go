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
	if err != nil {
		return Hp{}, err
	}
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
	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		return -1, err
	}
	nowTime := time.Now()

	//フォーマットの整形time型の日本時間で表示される
	thenUpdateDate, err := time.ParseInLocation(layout, updateDate, jst)
	if err != nil {
		return -1, err
	}

	//それぞれのtaskのダメージ計算
	for _, taskID := range taskIDs {

		rowsTaskFlag, err := db.Query("select isAchieve from tasks where id=?", taskID)
		if err != nil {
			return -1, err
		}
		var taskFlag int
		for rowsTaskFlag.Next() {
			taskFlag = 0
			rowsTaskFlag.Scan(&taskFlag)
		}
		//task成功してるなら
		if taskFlag == 1 {

			continue
		}

		judgmentOneweekagoFlag, err := judgmentTaskDealineOneWeekAgo(taskID)
		if err != nil {
			return -1, err
		}
		//１週間以上の場合計算処理をしない
		if judgmentOneweekagoFlag == true {
			continue
		}

		//それぞれのタスクの重さ
		rowsWeightIDs, err := db.Query("select weight_id from tasks where id=?", taskID)
		if err != nil {
			return -1, err
		}
		defer rowsWeightIDs.Close()
		//タスク一つの重さ
		var weightID int
		for rowsWeightIDs.Next() {
			weightID = 0
			rowsWeightIDs.Scan(&weightID)
		}
		if weightID == 0 {
			weightID = 1
		}

		//hpをアップデートした日(タスクを登録した時にもされる)と現在時刻の差
		diffUpdateDate := nowTime.Sub(thenUpdateDate)
		//time型をint型に変換したもの
		var intDiffUpdateDate int = int(diffUpdateDate.Seconds())

		var DeaadlineExist bool = false
		if DeaadlineExist, err = judgmentTaskDeadlineDateExist(taskID); err != nil {
			return -1, err
		}
		if DeaadlineExist == true {

			//戻り値はint型
			diffTaskDeadlineOneWeekAgo, err := returndiffTaskDeadlineOneWeekAgo(taskID)
			if err != nil {
				return -1, err
			}
			//taskが1週間以内に入ってきたときのダメージを正常化する処理
			if intDiffUpdateDate > diffTaskDeadlineOneWeekAgo {
				intDiffUpdateDate = diffTaskDeadlineOneWeekAgo
			}
		}

		totalDamage = totalDamage + intDiffUpdateDate*weightID

	}

	//time型をstring型に変換したもの"2020-08-24 22:46:04"のような形になる
	//stringUpdateNowTime := nowTime.Format(layout)
	//データベースのupdate_datetimeを現在時刻に変更
	_, err = db.Exec("update user_parameters set updated_datetime=Now() where id=?", temporaryUserID)
	if err != nil {
		return -1, err
	}

	currentHp := pastHp - totalDamage

	//hpが0以下の処理
	if currentHp < 0 {
		currentHp = 0
	}

	return currentHp, nil

}

func judgmentTaskDealineOneWeekAgo(taskID int) (bool, error) {
	deadlineDateExist, err := judgmentTaskDeadlineDateExist(taskID)
	if err != nil {
		return false, err
	}
	//締め切りが存在しない場合falseを返して終了
	if deadlineDateExist == false {
		return false, nil
	}

	//以下締め切りが存在する場合の処理
	//現在の日付と時刻
	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		return false, err
	}
	nowTime := time.Now()
	//日付と時刻を連絵したもの
	TaskDeadlineDateAndTime, err := returnTaskDeadlineDateAndTime(taskID)
	//フォーマットの整形time型の"2020-08-22 11:58:06のような形式で表示される
	thenTaskDeadlineDateAndTime, err := time.ParseInLocation(layout, TaskDeadlineDateAndTime, jst)
	//締め切りと現在時刻の差をとる
	diffTaskDeadlineDateAndTime := nowTime.Sub(thenTaskDeadlineDateAndTime)
	//締め切りと現在時刻の差が1週間以上の場合
	if diffTaskDeadlineDateAndTime.Hours() < -168 {

		return true, nil
	}
	return false, nil
}

func judgmentTaskDeadlineDateExist(taskID int) (bool, error) {
	rowsTaskDeadlineDate, err := db.Query("select deadline_date from tasks where id=?", taskID)
	if err != nil {
		return true, err
	}
	defer rowsTaskDeadlineDate.Close()
	var TaskDeadlineDate string
	for rowsTaskDeadlineDate.Next() {
		rowsTaskDeadlineDate.Scan(&TaskDeadlineDate)
	}
	//taskの締め切りが存在しない場合
	if TaskDeadlineDate == "" {

		return false, nil
	}
	return true, nil
}

func returnTaskDeadlineDateAndTime(taskID int) (string, error) {
	//締め切り日時の取得
	rowsTaskDeadline, err := db.Query("select deadline_date,deadline_time from tasks where id=?", taskID)
	if err != nil {
		return "", err
	}
	defer rowsTaskDeadline.Close()
	var taskDeadlineDate string
	var taskDeadlineTime string
	for rowsTaskDeadline.Next() {
		rowsTaskDeadline.Scan(&taskDeadlineDate, &taskDeadlineTime)
	}

	//締め切り日時と時刻を連結
	taskDeadlineDateAndTime := taskDeadlineDate + " " + taskDeadlineTime

	return taskDeadlineDateAndTime, nil
}

func returndiffTaskDeadlineOneWeekAgo(taskID int) (int, error) {
	//締め切りの日付と時間を取得
	taskDeadlineTypeString, err := returnTaskDeadlineDateAndTime(taskID)
	if err != nil {
		return -1, err
	}
	//string型→time型に変換
	taskDeadlineTypeTime, err := stringToTime(taskDeadlineTypeString)
	if err != nil {
		return -1, err
	}

	//締め切りを1週間前のtime型の取得
	TaskDeadlineOneWeekAgoTypeTime := taskDeadlineTypeTime.AddDate(0, 0, -7)
	//time型→string型
	TaskDeadlineOneWeekAgoTypeString := timeToString(TaskDeadlineOneWeekAgoTypeTime)
	//2018-04-06 17:59:20のような表示形式になる

	//現在の日付と時刻
	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		return -1, err
	}
	nowTime := time.Now()

	//フォーマットの整形time型の日本時間で表示される
	DeadlineOneWeekAgo, err := time.ParseInLocation(layout, TaskDeadlineOneWeekAgoTypeString, jst)
	if err != nil {
		return -1, err
	}

	//締め切り1週間前と現在時刻の差
	diffTaskDeadlineOneWeekAgo := nowTime.Sub(DeadlineOneWeekAgo)
	//float型をint型に変換したもの
	var intDiffTaskDeadlineOneWeekAgo int = int(diffTaskDeadlineOneWeekAgo.Seconds())

	return intDiffTaskDeadlineOneWeekAgo, nil

}

func timeToString(t time.Time) string {
	str := t.Format(layout)
	return str
}

func stringToTime(str string) (time.Time, error) {
	t, err := time.Parse(layout, str)
	if err != nil {
		return t, err
	}
	return t, nil
}

//RecoveryHp はタスクが達成されたときに20万回復する処理を行う
func RecoveryHp(token string) error {

	//トークンからparamIDを取得
	rows, err := db.Query("select param_id from users where token=?", token)
	if err != nil {
		return err
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
		return err
	}
	defer rows.Close()
	var pastHp int
	for rows.Next() {
		pastHp = 0
		rows.Scan(&pastHp)
	}

	recoveryAfterHp := pastHp + 200000

	if recoveryAfterHp > 1000000 {
		recoveryAfterHp = 1000000
	}

	//user_parametersの更新
	_, err = db.Exec("update user_parameters set hp=? where id=?", recoveryAfterHp, temporaryUserID)
	if err != nil {
		return err
	}

	return nil
}

//TaskIDUpdateTime はtaskIDの数を数えて時刻をアップデートする
func TaskIDUpdateTime(token string) error {

	/* 	taskIDs, err := callTaskIDsFromUserToken(token)
	   	if err != nil {
	   		return err
		   } */
	//usersのtokenからuser_parametersのidを取得
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

	_, err = db.Exec("update user_parameters set updated_datetime=Now() where id=?", temporaryUserID)
	if err != nil {
		return err
	}

	return nil
}

//ChangeHpMillion はhpを100万にする
func ChangeHpMillion(userToken string) error {
	rows, err := db.Query("select param_id from users where token=?", userToken)
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

	_, err = db.Exec("update user_parameters set hp=? where id=?", 1000000, temporaryUserID)
	if err != nil {
		return err
	}
	return nil

}
