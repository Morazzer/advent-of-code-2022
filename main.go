package main

import (
	"fmt"
	"morazzer.dev/advent-of-code/days"
)

func main() {
	var day string
	fmt.Print("Please enter the day: ")
	_, err := fmt.Scanln(&day)
	if err != nil {
		panic(err)
	}

	fmt.Println("Running code for day: " + day)

	fun := days.Days["day"+day].(func())
	fun()
}
