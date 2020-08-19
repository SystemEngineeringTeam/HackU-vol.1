package dbctl

import (
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

		tasks = append(tasks, t)
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
		date := ""
		time := ""
		description := ""
		weightID := 0

		rows.Scan(&id, &title, &date, &time, &description, &weightID)
		weightDegree, err := callWeightDegreeFromWeightID(weightID)
		if err != nil {
			return Task{}, err
		}

		task = Task{ID: id, Title: title, DeadlineDate: date, DeadlineTime: time, Description: description, Weight: weightDegree}
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
	rows, err := db.Query("select id from weights where degree=?", degree)
	if err != nil {
		return -1, err
	}

	id := 0
	for rows.Next() {
		rows.Scan(&id)
	}

	return id, nil
}

// RegisterNewTask は新しいタスクを登録する関数
func RegisterNewTask(token string, t Task) (int, error) {
	weightID, err := callWeightIDFromWeightDegree(t.Weight)
	if err != nil || weightID == 0 {
		pc, file, line, _ := runtime.Caller(0)
		f := runtime.FuncForPC(pc)
		log.Printf(errFormat, err, f.Name(), file, line)
		return -1, err
	}

	_, err = db.Query("insert into tasks(title,deadline_date,deadline_time,description,weight_id,isAchieve) values(?,?,?,?,?,false)", t.Title, convertNullString(t.DeadlineDate), convertNullString(t.DeadlineTime), convertNullString(t.Description), convertNullInt(weightID))
	if err != nil {
		pc, file, line, _ := runtime.Caller(0)
		f := runtime.FuncForPC(pc)
		log.Printf(errFormat, err, f.Name(), file, line)
		return -1, err
	}

	newTaskID, err := callTaskIDFromTaskTitle(t.Title)
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

	err = linkTaskIDAndUserID(newTaskID, userID)
	if err != nil {
		pc, file, line, _ := runtime.Caller(0)
		f := runtime.FuncForPC(pc)
		log.Printf(errFormat, err, f.Name(), file, line)
		return -1, err
	}

	return newTaskID, err
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

	_, err = db.Query("update tasks set isAchieve=true where id=?", taskID)
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

	_, err = db.Query("update tasks set isAchieve=false where id=?", taskID)
	if err != nil {
		pc, file, line, _ := runtime.Caller(0)
		f := runtime.FuncForPC(pc)
		log.Printf(errFormat, err, f.Name(), file, line)
		return err
	}
	return nil
}
