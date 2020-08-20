package dbctl

import (
	"database/sql"
	"errors"
	"log"
	"runtime"
)

//Hp はデータベースから取得した値を扱うための構造体
type Hp struct{
	Hp int `json:"hp"`
	MaxHp int `json:"maxHp"`
}

//CallHp は現在のhpを返す関数
func CallHp(token string)([]Hp,error){
	userID,err:=callUserIDFromToken(token)
	if err != nil {
		return nil, err
	}

	hpID,err:=callHpIDsFromUserID(userID)

			
}

func callHpIDFromUserID(userID int)(int,error){
	
	rows,err:=db.Query("select hp_id from users where user_id=?",token)

	if err != nil {
		return nil, err		
	}
	defer rows.Close()

	for rows.Next(){
		
	}

	
}

























