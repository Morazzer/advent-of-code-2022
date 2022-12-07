package days

import (
	"bufio"
	_ "embed"
	"strconv"
	"strings"
)

//go:embed inputs/day3.txt
var input3 string

func Day3() string {
	input := input3

	var priority int
	priority = 0

	charMapping := make(map[rune]int)

	for i := 1; i <= 26; i++ {
		charMapping[rune(97+i-1)] = i
		charMapping[rune(65+i-1)] = i + 26
	}

	scannner := bufio.NewScanner(strings.NewReader(input))
	for scannner.Scan() {
		line := scannner.Text()
		first, second := line[0:len(line)/2], line[len(line)/2:]
		for _, firstChar := range first {
			for _, secondChar := range second {
				if (firstChar == secondChar) && (strings.ContainsRune(first, firstChar) && strings.ContainsRune(second, secondChar)) {
					first = strings.ReplaceAll(first, string(firstChar), "")
					second = strings.ReplaceAll(second, string(firstChar), "")
					priority += charMapping[firstChar]
					break
				}
			}
		}
	}

	return strconv.Itoa(priority)
}

func Day3_2() string {
	input := input3

	var priority int
	priority = 0

	charMapping := make(map[rune]int)

	for i := 1; i <= 26; i++ {
		charMapping[rune(97+i-1)] = i
		charMapping[rune(65+i-1)] = i + 26
	}

	scannner := bufio.NewScanner(strings.NewReader(input))
	for scannner.Scan() {
		line1 := scannner.Text()
		scannner.Scan()
		line2 := scannner.Text()
		scannner.Scan()
		line3 := scannner.Text()

		for _, rune := range line1 {
			if strings.ContainsRune(line2, rune) && strings.ContainsRune(line3, rune) {
				priority += charMapping[rune]
				break
			}
		}
	}

	return strconv.Itoa(priority)
}
