package dbctl

import "crypto/sha256"

// User はデータベースから取得したユーザーのデータを扱うための構造体
type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Pass  string `json:"pass"`
	Token string `json:"token"`
}

func callUserIDFromToken(token string) (int, error) {
	rows, err := db.Query("select id from users where token=?", token)
	if err != nil {
		return -1, err
	}
	defer rows.Close()

	userID := -1
	for rows.Next() {
		rows.Scan(&userID)
	}

	return userID, nil
}

// Login はメールアドレスとパスワードからトークンを取得するための関数
func Login(email, password string) (User, error) {
	hashedPassword := string(sha256.New().Sum([]byte(password)))
	rows, err := db.Query("select token from users where email=? and password=?", email, hashedPassword)
	if err != nil {
		return User{}, err
	}

	name := ""
	token := ""
	for rows.Next() {
		rows.Scan(&name, &token)
	}

	return User{Name: name, Token: token}, nil
}
