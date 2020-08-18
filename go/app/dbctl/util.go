package dbctl

import (
	"database/sql"
	"log"
	"runtime"
)

//エラーの内容:err 関数の名前:f.Name() ファイルのパス:file runtimeが呼ばれた行数:line
const errFormat = "%v\nfunction:%v file:%v line:%v\n"

var db *sql.DB

// packageがimportされたときに呼び出される関数
func init() {
	pc, file, line, _ := runtime.Caller(0)
	f := runtime.FuncForPC(pc)

	var err error

	db, err = sql.Open("mysql", "gopher:setsetset@tcp(vol1_mysql:3306)/production_db")
	if err != nil {
		log.Printf(errFormat, err, f.Name(), file, line)

		//データベースを開けないと動作が継続できないためpanicを発生させる
		panic("Can't Open database.")
	}
}
