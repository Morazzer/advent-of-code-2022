package days

import (
	"bufio"
	_ "embed"
	"fmt"
	"strings"
)

//go:embed inputs/day6.txt
var input6 string

func Day6() {
	input := input6
	scannner := bufio.NewScanner(strings.NewReader(input))
	for scannner.Scan() {
		line := scannner.Text()

		recieved := make([]rune, 0)

		for _, r := range line {
			recieved = append(recieved, r)
			i := len(recieved) - 1
			if i >= 4 {
				if recieved[i] != recieved[i-1] && recieved[i] != recieved[i-2] && recieved[i] != recieved[i-3] && recieved[i-1] != recieved[i-2] && recieved[i-1] != recieved[i-3] && recieved[i-2] != recieved[i-3] {
					fmt.Printf("Solution: %d", i+1)
					return
				}
			}
		}
	}
}

func Day6_2() {
	input := input6
	scannner := bufio.NewScanner(strings.NewReader(input))
	for scannner.Scan() {
		line := scannner.Text()

		recieved := make([]rune, 0)

		for _, r := range line {
			recieved = append(recieved, r)
			i := len(recieved) - 1
			if i >= 14 {
				real := true
				for h := 0; h < 14; h++ {
					for k := 0; k < 14; k++ {
						if h == k {
							continue
						}
						if recieved[i-k] == recieved[i-h] {
							real = false
							break
						}
					}
					if !real {
						break
					}
				}
				if real {
					fmt.Printf("Solution: %d\n", i+1)
				}
			}
		}
	}
}
