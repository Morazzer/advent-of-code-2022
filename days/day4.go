package days

import (
	"bufio"
	_ "embed"
	"strconv"
	"strings"
)

//go:embed inputs/day4.txt
var input4 string

func Day4() string {
	input := input4

	count := 0

	scannner := bufio.NewScanner(strings.NewReader(input))
	for scannner.Scan() {
		line := scannner.Text()
		split := strings.Split(line, ",")
		first, second := split[0], split[1]

		firstSplit := strings.Split(first, "-")
		firstStart, _ := strconv.Atoi(firstSplit[0])
		firstEnd, _ := strconv.Atoi(firstSplit[1])

		secondSplit := strings.Split(second, "-")
		secondStart, _ := strconv.Atoi(secondSplit[0])
		secondEnd, _ := strconv.Atoi(secondSplit[1])

		if (firstStart <= secondStart && firstEnd >= secondEnd) || (secondStart <= firstStart && secondEnd >= firstEnd) {
			count += 1
		}

	}

	return strconv.Itoa(count)
}

func Day4_2() string {
	input := input4

	count := 0

	scannner := bufio.NewScanner(strings.NewReader(input))
	for scannner.Scan() {
		line := scannner.Text()
		split := strings.Split(line, ",")
		first, second := split[0], split[1]

		firstSplit := strings.Split(first, "-")
		firstStart, _ := strconv.Atoi(firstSplit[0])
		firstEnd, _ := strconv.Atoi(firstSplit[1])

		secondSplit := strings.Split(second, "-")
		secondStart, _ := strconv.Atoi(secondSplit[0])
		secondEnd, _ := strconv.Atoi(secondSplit[1])

		firstList := makeRange(firstStart, firstEnd)
		secondList := makeRange(secondStart, secondEnd)

		if contains(firstList, secondList) {
			count++
		}
	}

	return strconv.Itoa(count)
}

func contains(first, second []int) bool {
	for _, i := range first {
		for _, j := range second {
			if i == j {
				return true
			}
		}
	}
	return false
}

func makeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}
