package days

import (
	"bufio"
	_ "embed"
	"strconv"
	"strings"
)

//go:embed inputs/day8.txt
var input8 string

func Day8() string {
	input := input8
	i := 0
	field := make(map[int][]int)

	scannner := bufio.NewScanner(strings.NewReader(input))
	for scannner.Scan() {
		line := scannner.Text()

		for _, r := range line {
			num, _ := strconv.Atoi(string(r))
			field[i] = append(field[i], num)
		}
		i += 1
	}

	rows := len(field)
	columns := len(field[0])

	count := 0

	for i = 1; i < len(field)-1; i++ {
		for l, num := range field[i] {
			if l == 0 || (l == len(field[i])-1) {
				continue
			}
			top, right, left, bottom := true, true, true, true
			for r := 0; r < i; r++ {
				if field[r][l] >= num {
					top = false
				}
			}
			for c := 0; c < l; c++ {
				if field[i][c] >= num {
					left = false
				}
			}

			for r := i + 1; r < rows; r++ {
				if field[r][l] >= num {
					bottom = false
				}
			}
			for c := l + 1; c < columns; c++ {
				if field[i][c] >= num {
					right = false
				}
			}
			if top || bottom || left || right {
				count += 1
			}
		}
	}

	count += rows*2 + columns*2 - 4

	return strconv.Itoa(count)
}

func Day8_2() string {
	input := input8
	i := 0
	field := make(map[int][]int)

	scannner := bufio.NewScanner(strings.NewReader(input))
	for scannner.Scan() {
		line := scannner.Text()

		for _, r := range line {
			num, _ := strconv.Atoi(string(r))
			field[i] = append(field[i], num)
		}
		i += 1
	}

	max := 0

	for y := 1; y < len(field)-1; y++ {
		for x := 1; x < len(field[y])-1; x++ {
			v1, v2, v3, v4 := 0, 0, 0, 0
			num := field[y][x]

			r := y
			for r > 0 {
				r--
				v1 += 1
				if field[r][x] >= num {
					break
				}
			}
			r = y
			for r < len(field)-1 {
				r++
				v2 += 1
				if field[r][x] >= num {
					break
				}
			}

			c := x
			for c > 0 {
				c--
				v3 += 1
				if field[y][c] >= num {
					break
				}
			}
			c = x
			for c < len(field[y])-1 {
				c++
				v4 += 1
				if field[y][c] >= num {
					break
				}
			}
			temp := v1 * v2 * v3 * v4
			if temp > max {
				max = temp
			}
		}
	}

	return strconv.Itoa(max)
}
