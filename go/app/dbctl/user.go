package dbctl

import (
	"crypto/sha256"
	"encoding/hex"
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
		return -1, err
	}
	defer rows.Close()

	userID := -1
	for rows.Next() {
		rows.Scan(&userID)
	}

	return userID, nil
}
