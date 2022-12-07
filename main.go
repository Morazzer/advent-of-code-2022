package main

import (
	"fmt"
	"morazzer.dev/advent-of-code/days"
	"strconv"
)

func main() {
	var day string
	fmt.Print("Please enter the day: ")
	_, err := fmt.Scanln(&day)
	if err != nil {
		panic(err)
	}

	if day == "all" {
		fmt.Println()
		for i := 0; i < 25; i++ {
			fmt.Printf("Solutions for day %d\n", i+1)

			p1 := days.Days["day"+strconv.Itoa(i+1)]
			p2 := days.Days["day"+strconv.Itoa(i+1)+"_2"]

			p1s := "Not implemented yet"
			p2s := p1s
			if p1 != nil {
				p1s = p1()
			}
			if p2 != nil {
				p2s = p2()
			}

			fmt.Printf("Part 1: %s\n", p1s)
			fmt.Printf("Part 2: %s\n", p2s)
			fmt.Println()
		}
		return
	}

	fun := days.Days["day"+day]
	result := fun()
	fmt.Printf("Solution for day %s: %s", day, result)
}
