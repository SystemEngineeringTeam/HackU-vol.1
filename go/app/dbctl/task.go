package dbctl

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
	rows, err := db.Query("select id,title,deadlineDate,deadlineTime,description,weight from tasks where id=? and isAchieve=false", taskID)
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
