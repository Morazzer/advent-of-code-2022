package days

import (
	"bufio"
	_ "embed"
	"math"
	"strconv"
	"strings"
)

//go:embed inputs/day9.txt
var input9 string

func Day9() string {
	input := input9
	visited := make(map[string]int)

	hx, hy := 1, 5
	tx, ty := hx, hy

	scannner := bufio.NewScanner(strings.NewReader(input))
	for scannner.Scan() {
		line := scannner.Text()

		split := strings.Split(line, " ")

		direction := split[0]
		steps, _ := strconv.Atoi(split[1])

		for i := 0; i < steps; i++ {
			switch direction {
			case "R":
				hx += 1
			case "U":
				hy -= 1
			case "D":
				hy += 1
			case "L":
				hx -= 1
			}

			if math.Abs(float64(hx-tx)) > 1 || math.Abs(float64(hy-ty)) > 1 {
				if hx != tx {
					add := 0
					if hx > tx {
						add += 1
					} else {
						add -= 1
					}
					tx += add
				}
				if hy != ty {
					add := 0
					if hy > ty {
						add++
					} else {
						add--
					}
					ty += add
				}
			}

			visited[strconv.Itoa(tx)+","+strconv.Itoa(ty)] = 1

		}
	}

	return strconv.Itoa(len(visited))
}

func Day9_2() string {
	input := input9
	visited := make(map[string]int)

	elemnts := make([]TailElement, 10)
	for i := 0; i < 10; i++ {
		elemnts[i] = TailElement{
			x: 1,
			y: 5,
		}
	}

	scannner := bufio.NewScanner(strings.NewReader(input))
	for scannner.Scan() {
		line := scannner.Text()

		split := strings.Split(line, " ")

		direction := split[0]
		steps, _ := strconv.Atoi(split[1])

		for i := 0; i < steps; i++ {
			switch direction {
			case "R":
				elemnts[0].x += 1
			case "U":
				elemnts[0].y -= 1
			case "D":
				elemnts[0].y += 1
			case "L":
				elemnts[0].x -= 1
			}

			for i := 1; i < 10; i++ {
				hx, hy := elemnts[i-1].x, elemnts[i-1].y
				tx, ty := elemnts[i].x, elemnts[i].y
				if math.Abs(float64(hx-tx)) > 1 || math.Abs(float64(hy-ty)) > 1 {
					if hx != tx {
						add := 0
						if hx > tx {
							add += 1
						} else {
							add -= 1
						}
						elemnts[i].x += add
					}
					if hy != ty {
						add := 0
						if hy > ty {
							add++
						} else {
							add--
						}
						elemnts[i].y += add
					}
				}
			}

			last := elemnts[9]
			visited[strconv.Itoa(last.x)+","+strconv.Itoa(last.y)] = 1
		}
	}

	return strconv.Itoa(len(visited))
}

type TailElement struct {
	x int
	y int
}
