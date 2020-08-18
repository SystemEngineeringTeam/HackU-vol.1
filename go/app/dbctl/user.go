package dbctl

import (
	"crypto/sha256"
	"encoding/hex"
	"log"
	"runtime"
)

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

		pc, file, line, _ := runtime.Caller(0)
		f := runtime.FuncForPC(pc)
		log.Printf(errFormat, err, f.Name(), file, line)

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

	hashPass := sha256.Sum256([]byte(password))
	encodePass := hex.EncodeToString(hashPass[:])
	rows, err := db.Query("select token from users where email=? and password=?", email, encodePass)
	if err != nil {

		pc, file, line, _ := runtime.Caller(0)
		f := runtime.FuncForPC(pc)
		log.Printf(errFormat, err, f.Name(), file, line)

		return User{}, err
	}
	defer rows.Close()

	name := ""
	token := ""
	for rows.Next() {
		rows.Scan(&name, &token)
	}

	return User{Name: name, Token: token}, nil
}

// RegisterNewUser はユーザの登録を行う関数です
func RegisterNewUser(u User) error {
	hashPass := sha256.Sum256([]byte(u.Pass))
	encodePass := hex.EncodeToString(hashPass[:])
	hashEmail := sha256.Sum256([]byte(u.Email))
	token := hex.EncodeToString(hashEmail[:])
	_, err := db.Query("insert into users(name,email,password,token) values (?,?,?,?)", u.Name, u.Email, encodePass, token)
	if err != nil {

		pc, file, line, _ := runtime.Caller(0)
		f := runtime.FuncForPC(pc)
		log.Printf(errFormat, err, f.Name(), file, line)

		return err
	}

	return nil
}
