package main

import (
	"fmt"
	"time"
)

func main() {

	date := time.Now()
	fmt.Println(date)

	format := "2006-01-02 15:04:05"
	then, _ := time.Parse(format, "2020-08-22 11:58:06")
	fmt.Println(then)

	diff := date.Sub(then)

	//func Since(t Time) Duration
	//Since returns the time elapsed since t.
	//It is shorthand for time.Now().Sub(t).

	fmt.Println(diff.Hours())       //number of Hours
	fmt.Println(diff.Nanoseconds()) //number of Nanoseconds
	fmt.Println(diff.Minutes())     //number of Minutes
	fmt.Println(diff.Seconds())     //number of Seconds


	fmt.Println(int(diff.Hours() / 24)) //number of days

	t := time.Now()
	fmt.Println(t)
	fmt.Println(t.Year())
	fmt.Println(t.Month())
	fmt.Println(t.Day())
	fmt.Println(t.Hour())
	fmt.Println(t.Minute())
	fmt.Println(t.Second())
	fmt.Println(t.Weekday())
	
}
