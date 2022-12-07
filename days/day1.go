package days

import (
	"bufio"
	_ "embed"
	"sort"
	"strconv"
	"strings"
)

//go:embed inputs/day1.txt
var input1 string

func Day1() string {
	input := input1
	var elves map[int]uint64
	elves = make(map[int]uint64)

	i := 0

	scannner := bufio.NewScanner(strings.NewReader(input))
	for scannner.Scan() {
		line := scannner.Text()
		if line == "" {
			i++
		} else {
			l, _ := strconv.ParseUint(line, 10, 64)
			if val, ok := elves[i]; ok {
				elves[i] = val + l
			} else {
				elves[i] = l
			}
		}
	}

	var _ int
	var highest uint64
	highest = 0

	for elve, number := range elves {
		if number > highest {
			highest = number
			_ = elve
		}
	}

	return strconv.FormatUint(highest, 10)
}

func Day1_2() string {
	input := input1
	var elves map[int]uint64
	elves = make(map[int]uint64)

	i := 0

	scannner := bufio.NewScanner(strings.NewReader(input))
	for scannner.Scan() {
		line := scannner.Text()
		if line == "" {
			i++
		} else {
			l, _ := strconv.ParseUint(line, 10, 64)
			if val, ok := elves[i]; ok {
				elves[i] = val + l
			} else {
				elves[i] = l
			}
		}
	}

	elven := make([]Pair, 0, len(elves))

	i = 0
	for k, v := range elves {
		elven = append(elven, Pair{
			Key:   k,
			Value: v,
		})
	}

	sort.Slice(elven, func(i, j int) bool {
		return elven[i].Value > elven[j].Value
	})

	//fmt.Printf("First: %d - %d\n", elven[0].Key, elven[0].Value)
	//fmt.Printf("Second: %d - %d\n", elven[1].Key, elven[1].Value)
	//fmt.Printf("Third: %d - %d\n", elven[2].Key, elven[2].Value)
	//fmt.Printf("Solution: %d\n", elven[0].Value+elven[1].Value+elven[2].Value)

	return strconv.FormatUint(elven[0].Value+elven[1].Value+elven[2].Value, 10)

}

type Pair struct {
	Key   int
	Value uint64
}
