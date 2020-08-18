package apifuncs

import (
	"log"
	"net/http"

	"set1.ie.aitech.ac.jp/HackU_vol_1/dbctl"
)

// DBTest is ..
func DBTest(w http.ResponseWriter, r *http.Request) {
	err := dbctl.RegisterNewUser(dbctl.User{Name: "Hoge", Email: "hoge@hoge.jp", Pass: "hogehoge"})
	if err != nil {
		log.Fatal(err)
		return
	}
}
