package dbctl

import "crypto/sha256"

// User はデータベースから取得したユーザーのデータを扱うための構造体
type User struct {
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

// RegisterNewUser はユーザの登録を行う関数です
func RegisterNewUser(u User) error {
	hashedPassword := string(sha256.New().Sum([]byte(u.Pass)))
	token := string(sha256.New().Sum([]byte(u.Email)))
	_, err := db.Query("insert into users(name,email,password,token) values (?,?,?,?)", u.Name, u.Email, hashedPassword, token)
	if err != nil {
		return err
	}
	return nil
}
