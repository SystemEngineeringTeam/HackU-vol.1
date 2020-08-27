package main

import (
	"fmt"
	"time"
)

var layout = "2006-01-02 15:04:05"

func main() {
	//jst, _ := time.LoadLocation("Asia/Tokyo")

	nowTime := time.Now()
	nextTime := nowTime.AddDate(0, 0, -7)

	fmt.Println(nowTime)
	fmt.Println(nextTime)
}
