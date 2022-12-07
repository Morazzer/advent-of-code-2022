package days

import (
	"bufio"
	_ "embed"
	"strconv"
	"strings"
)

//go:embed inputs/day5.txt
var input5 string

func Day5() string {
	input := input5

	var finishedInit = false

	stacks := make(map[int][]rune)

	scannner := bufio.NewScanner(strings.NewReader(input))
	for scannner.Scan() {
		line := scannner.Text()

		if line == "" {

			for k, v := range stacks {
				for i, j := 0, len(v)-1; i < j; i, j = i+1, j-1 {
					v[i], v[j] = v[j], v[i]
				}
				stacks[k] = v
			}

			finishedInit = true
			continue
		}

		stack := 0
		if !finishedInit {
			for i, char := range line {
				if i%4 == 0 {
					stack++
				}
				if char == ' ' || char == '[' || char == ']' {
					continue
				}
				if char == '1' {
					break
				}
				stacks[stack] = append(stacks[stack], char)
			}
			continue
		}

		split := strings.Split(line, " ")

		amount, _ := strconv.Atoi(split[1])
		from, _ := strconv.Atoi(split[3])
		to, _ := strconv.Atoi(split[5])

		for i := 0; i < amount; i++ {
			if len(stacks[from]) > 0 {
				item := stacks[from][len(stacks[from])-1]
				stacks[from] = stacks[from][:len(stacks[from])-1]
				stacks[to] = append(stacks[to], item)
			}
		}
	}

	s := ""
	for i := 1; i < 10; i++ {
		v := stacks[i]
		if len(v) > 0 {
			s += string(v[len(v)-1])
		} else {
			s += " "
		}
	}

	return s
}

func Day5_2() string {
	input := input5

	var finishedInit = false

	stacks := make(map[int][]rune)

	scannner := bufio.NewScanner(strings.NewReader(input))
	for scannner.Scan() {
		line := scannner.Text()

		if line == "" {

			for k, v := range stacks {
				for i, j := 0, len(v)-1; i < j; i, j = i+1, j-1 {
					v[i], v[j] = v[j], v[i]
				}
				stacks[k] = v
			}

			finishedInit = true
			continue
		}

		stack := 0
		if !finishedInit {
			for i, char := range line {
				if i%4 == 0 {
					stack++
				}
				if char == ' ' || char == '[' || char == ']' {
					continue
				}
				if char == '1' {
					break
				}
				stacks[stack] = append(stacks[stack], char)
			}
			continue
		}

		split := strings.Split(line, " ")

		amount, _ := strconv.Atoi(split[1])
		from, _ := strconv.Atoi(split[3])
		to, _ := strconv.Atoi(split[5])

		items := stacks[from][len(stacks[from])-amount:]
		stacks[from] = stacks[from][:len(stacks[from])-amount]
		stacks[to] = append(stacks[to], items...)
	}

	s := ""
	for i := 1; i < 10; i++ {
		v := stacks[i]
		if len(v) > 0 {
			s += string(v[len(v)-1])
		} else {
			s += " "
		}
	}

	return s
}
