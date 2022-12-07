package days

import (
	"bufio"
	_ "embed"
	"strconv"
	"strings"
)

//go:embed inputs/day2.txt
var input2 string

func Day2() string {
	input := input2

	var score uint64
	score = 0

	scannner := bufio.NewScanner(strings.NewReader(input))
	for scannner.Scan() {
		line := scannner.Text()
		split := strings.Split(line, " ")
		they, you := split[0], split[1]
		if (they == "A" && you == "X") || (they == "B" && you == "Y") || (they == "C" && you == "Z") {
			score += 3
			/*
			 * A = ROCK
			 * B = PAPER
			 * C = SCISSORS
			 * ------------
			 * X = ROCK
			 * Y = PAPER
			 * Z = SCISSORS
			 */
		} else if (they == "A" && you == "Y") || (they == "B" && you == "Z") || (they == "C" && you == "X") {
			score += 6
		}
		if you == "X" {
			score += 1
		} else if you == "Y" {
			score += 2
		} else if you == "Z" {
			score += 3
		}
	}

	return strconv.FormatUint(score, 10)
}

func Day2_2() string {
	input := input2

	var score uint64
	score = 0

	loses := make(map[string]string)
	wins := make(map[string]string)
	draws := make(map[string]string)

	loses["A"] = "Z"
	loses["B"] = "X"
	loses["C"] = "Y"

	draws["A"] = "X"
	draws["B"] = "Y"
	draws["C"] = "Z"

	wins["A"] = "Y"
	wins["B"] = "Z"
	wins["C"] = "X"

	/*
	 * A = ROCK
	 * B = PAPER
	 * C = SCISSORS
	 * ------------
	 * X = ROCK
	 * Y = PAPER
	 * Z = SCISSORS
	 */

	scannner := bufio.NewScanner(strings.NewReader(input))
	for scannner.Scan() {
		line := scannner.Text()
		split := strings.Split(line, " ")
		they, you := split[0], split[1]

		if you == "X" {
			you = loses[they]
		} else if you == "Y" {
			score += 3
			you = draws[they]
		} else if you == "Z" {
			score += 6
			you = wins[they]
		}

		if you == "X" {
			score += 1
		} else if you == "Y" {
			score += 2
		} else if you == "Z" {
			score += 3
		}
	}
	return strconv.FormatUint(score, 10)
}
