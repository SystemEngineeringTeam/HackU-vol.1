package dbctl

import (
	"database/sql"
	"errors"
	"log"
	"runtime"
)

// Task はデータベースから取得した値を扱うための構造体
type Task struct {
	ID           int    `json:"id"`
	Title        string `json:"title"`
	DeadlineDate string `json:"deadlineDate"`
	DeadlineTime string `json:"deadlineTime"`
	Description  string `json:"description"`
	Weight       string `json:"weight"`
}

// CallTasks はタスク一覧を返す関数
func CallTasks(token string) ([]Task, error) {
	userID, err := callUserIDFromToken(token)
	if err != nil {
		return nil, err
	}

	taskIDs, err := callTaskIDsFromUserID(userID)
	if err != nil {
		return nil, err
	}

	tasks := make([]Task, 0, 0)

	for _, id := range taskIDs {
		t, err := callTaskFromTaskID(id)
		if err != nil {
			return nil, err
		}

		if t.ID != 0 {
			tasks = append(tasks, t)
		}
	}

	if len(tasks) == 0 {
		return nil, nil
	}

	return tasks, nil
}

func callTaskIDsFromUserID(userID int) ([]int, error) {
	taskIDs := make([]int, 0, 0)
	rows, err := db.Query("select task_id from user_and_task_links where user_id=?", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		temporaryID := 0
		rows.Scan(&temporaryID)
		taskIDs = append(taskIDs, temporaryID)
	}

	return taskIDs, nil
}

func callTaskFromTaskID(taskID int) (Task, error) {
	rows, err := db.Query("select id,title,deadline_date,deadline_time,description,weight_id from tasks where id=? and isAchieve=false", taskID)
	if err != nil {
		return Task{}, err
	}
	defer rows.Close()

	task := Task{}

	for rows.Next() {
		id := 0
		title := ""
		var date sql.NullString = sql.NullString{}
		var time sql.NullString = sql.NullString{}
		var description sql.NullString = sql.NullString{}
		weightID := 0

		rows.Scan(&id, &title, &date, &time, &description, &weightID)
		weightDegree, err := callWeightDegreeFromWeightID(weightID)
		if err != nil {
			return Task{}, err
		}

		task = Task{ID: id, Title: title, DeadlineDate: convertToString(date), DeadlineTime: convertToString(time), Description: convertToString(description), Weight: weightDegree}
	}
	return task, nil
}

func callWeightDegreeFromWeightID(weightID int) (string, error) {
	rows, err := db.Query("select degree from weights where id=?", weightID)
	if err != nil {
		return "", err
	}

	degree := ""
	for rows.Next() {
		rows.Scan(&degree)
	}

	return degree, nil
}

func callWeightIDFromWeightDegree(degree string) (int, error) {
	if len(degree) <= 0 {
		return 0, nil
	}

	rows, err := db.Query("select id from weights where degree=?", degree)
	if err != nil {
		return -1, err
	}

	id := -1
	for rows.Next() {
		rows.Scan(&id)
	}

	return id, nil
}

// RegisterNewTask は新しいタスクを登録する関数
func RegisterNewTask(token string, t Task) (int, error) {
	weightID, err := callWeightIDFromWeightDegree(t.Weight)
	if err != nil || weightID == -1 {
		pc, file, line, _ := runtime.Caller(0)
		f := runtime.FuncForPC(pc)
		log.Printf(errFormat, err, f.Name(), file, line)
		return -1, err
	}

	res, err := db.Exec("insert into tasks(title,deadline_date,deadline_time,description,weight_id,isAchieve,registered_datetime) values(?,?,?,?,?,false,Now())", t.Title, convertToNullString(t.DeadlineDate), convertToNullString(t.DeadlineTime), convertToNullString(t.Description), convertToNullInt(weightID))
	if err != nil {
		pc, file, line, _ := runtime.Caller(0)
		f := runtime.FuncForPC(pc)
		log.Printf(errFormat, err, f.Name(), file, line)
		return -1, err
	}

	newTaskID, err := res.LastInsertId()
	if err != nil {
		pc, file, line, _ := runtime.Caller(0)
		f := runtime.FuncForPC(pc)
		log.Printf(errFormat, err, f.Name(), file, line)
		return -1, err
	}

	userID, err := callUserIDFromToken(token)
	if err != nil {
		pc, file, line, _ := runtime.Caller(0)
		f := runtime.FuncForPC(pc)
		log.Printf(errFormat, err, f.Name(), file, line)
		return -1, err
	}

	err = linkTaskIDAndUserID(int(newTaskID), userID)
	if err != nil {
		pc, file, line, _ := runtime.Caller(0)
		f := runtime.FuncForPC(pc)
		log.Printf(errFormat, err, f.Name(), file, line)
		return -1, err
	}

	return int(newTaskID), err
}

func callTaskIDFromTaskTitle(title string) (int, error) {
	rows, err := db.Query("select id from tasks where title=?", title)
	if err != nil {
		pc, file, line, _ := runtime.Caller(0)
		f := runtime.FuncForPC(pc)
		log.Printf(errFormat, err, f.Name(), file, line)
		return -1, err
	}
	id := -1
	for rows.Next() {
		rows.Scan(&id)
	}
	return id, nil
}

func linkTaskIDAndUserID(taskID, userID int) error {
	_, err := db.Query("insert into user_and_task_links(task_id,user_id) values(?,?)", taskID, userID)
	if err != nil {
		pc, file, line, _ := runtime.Caller(0)
		f := runtime.FuncForPC(pc)
		log.Printf(errFormat, err, f.Name(), file, line)
		return err
	}
	return nil
}

// TaskAchieveFlagChangeToTrue はタスクの完了状況を達成済みにする関数
func TaskAchieveFlagChangeToTrue(token string, taskID int) error {
	userID, err := callUserIDFromToken(token)
	if err != nil {
		pc, file, line, _ := runtime.Caller(0)
		f := runtime.FuncForPC(pc)
		log.Printf(errFormat, err, f.Name(), file, line)
		return err
	}
	tIDs, err := callTaskIDsFromUserID(userID)
	if err != nil {
		pc, file, line, _ := runtime.Caller(0)
		f := runtime.FuncForPC(pc)
		log.Printf(errFormat, err, f.Name(), file, line)
		return err
	}
	isInvalidUserID := true
	for _, tID := range tIDs {
		if tID == taskID {
			isInvalidUserID = false
		}
	}
	if isInvalidUserID {
		return errors.New("Invalid userID")
	}

	_, err = db.Exec("update tasks set isAchieve=true where id=?", taskID)
	if err != nil {
		pc, file, line, _ := runtime.Caller(0)
		f := runtime.FuncForPC(pc)
		log.Printf(errFormat, err, f.Name(), file, line)
		return err
	}
	return nil
}

// TaskAchieveFlagChangeToFalse はタスクの完了状況を未達成にする関数
func TaskAchieveFlagChangeToFalse(token string, taskID int) error {
	userID, err := callUserIDFromToken(token)
	if err != nil {
		pc, file, line, _ := runtime.Caller(0)
		f := runtime.FuncForPC(pc)
		log.Printf(errFormat, err, f.Name(), file, line)
		return err
	}
	tIDs, err := callTaskIDsFromUserID(userID)
	if err != nil {
		pc, file, line, _ := runtime.Caller(0)
		f := runtime.FuncForPC(pc)
		log.Printf(errFormat, err, f.Name(), file, line)
		return err
	}

	isInvalidUserID := true
	for _, tID := range tIDs {
		if tID == taskID {
			isInvalidUserID = false
		}
	}
	if isInvalidUserID {
		return errors.New("Invalid userID")
	}

	_, err = db.Exec("update tasks set isAchieve=false where id=?", taskID)
	if err != nil {
		pc, file, line, _ := runtime.Caller(0)
		f := runtime.FuncForPC(pc)
		log.Printf(errFormat, err, f.Name(), file, line)
		return err
	}
	return nil
}

// CallWeightsList は難易度一覧を返す関数
func CallWeightsList() (string, error) {
	rows, err := db.Query("select degree from weights")
	if err != nil {
		pc, file, line, _ := runtime.Caller(0)
		f := runtime.FuncForPC(pc)
		log.Printf(errFormat, err, f.Name(), file, line)
		return "", err
	}

	weights := make([]string, 0, 0)
	for rows.Next() {
		w := sql.NullString{}
		rows.Scan(&w)
		if w.Valid {
			weights = append(weights, w.String)
		}
	}
	arrayString := convertStringArrayToJSONArray(weights)

	return arrayString, nil
}

//CompleteTask はtaskが全て完了しているかを判断する
func CompleteTask(token string) (bool, error) {

	userID, err := callUserIDFromToken(token)
	if err != nil {
		return false, err
	}
	taskIDs, err := callTaskIDsFromUserID(userID)
	if err != nil {
		return false, err
	}

	completeTaskCount := 0
	for _, taskID := range taskIDs {

		rowsTaskFlag, err := db.Query("select isAchieve from tasks where id=?", taskID)
		if err != nil {
			return false, err
		}
		var taskFlag int
		for rowsTaskFlag.Next() {
			taskFlag = 0
			rowsTaskFlag.Scan(&taskFlag)
		}
		if taskFlag == 1 {
			completeTaskCount++
		}
	}

	//タスクが全て完了している場合
	if len(taskIDs) == completeTaskCount {
		return true, nil
	}

	return false, nil
}
