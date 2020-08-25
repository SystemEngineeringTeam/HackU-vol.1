package apifuncs

import (
	"fmt"
	"net/http"
	"os"

	"set1.ie.aitech.ac.jp/HackU_vol_1/dbctl"
)

// DBTest is ..
func DBTest(w http.ResponseWriter, r *http.Request) {
	// err := dbctl.RegisterNewUser(dbctl.User{Name: "Hoge", Email: "hoge@hoge.jp", Pass: "hogehoge"})
	// if err != nil {
	// 	log.Fatal(err)
	// 	return	
	// }
	// token := "56f91b5f3668c470912be72ea6cbb0567cfdc0e6ab2266505f3f4b30bab989c6"
	// id, err := dbctl.RegisterNewTask("56f91b5f3668c470912be72ea6cbb0567cfdc0e6ab2266505f3f4b30bab989c6", dbctl.Task{Title: "fuga", Weight: "えぐい"})
	// if err != nil {
	// 	return
	// }
	// fmt.Println(id)	
	weights, err := dbctl.CallWeightsList()
	if err != nil {
		return
	}

	fmt.Fprintln(w, weights)
	fmt.Fprintln(os.Stdout, weights)
}
