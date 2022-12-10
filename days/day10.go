package days

import (
	"bufio"
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed inputs/day10.txt
var input10 string

func Day10() string {
	input := input10

	var x int = 1
	var cycle int = 0
	var signalstrength uint64 = 0

	test := func() {
		if cycle == 20 || (cycle-20)%40 == 0 {
			signalstrength += uint64(x * cycle)
		}
	}

	scannner := bufio.NewScanner(strings.NewReader(input))
	for scannner.Scan() {
		line := scannner.Text()

		cycle += 1
		test()

		if strings.HasPrefix(line, "addx") {
			cycle += 1
			num, _ := strconv.Atoi(line[5:])
			test()
			x += num
		}
	}
	return strconv.FormatUint(signalstrength, 10)
}

func Day10_2() string {
	input := input10

	var x = 1
	var cycle = 0
	out := ""

	draw := func() {

		temp := ((cycle + 1) % 40) - 1

		if x-1 == temp || x == temp || x+1 == temp {
			out += "#"
		} else {
			out += " "
		}

		if (cycle+1)%40 == 0 {
			fmt.Println(out)
			out = ""
		}
	}

	scannner := bufio.NewScanner(strings.NewReader(input))
	for scannner.Scan() {
		line := scannner.Text()

		draw()

		cycle += 1

		if strings.HasPrefix(line, "addx") {
			draw()
			cycle += 1
			num, _ := strconv.Atoi(line[5:])
			x += num
		}
	}
	return "^^^^^^^^^^^^^^^^"
}
