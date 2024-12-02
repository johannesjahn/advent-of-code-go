package main

import (
	_ "embed"
	"flag"
	"fmt"
	"slices"
	"strings"

	"github.com/alexchao26/advent-of-code-go/cast"
	"github.com/alexchao26/advent-of-code-go/util"
)

//go:embed input.txt
var input string

func init() {
	// do this in init (not main) so test file has same input
	input = strings.TrimRight(input, "\n")
	if len(input) == 0 {
		panic("empty input.txt file")
	}
}

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()
	fmt.Println("Running part", part)

	if part == 1 {
		ans := part1(input)
		util.CopyToClipboard(fmt.Sprintf("%v", ans))
		fmt.Println("Output:", ans)
	} else {
		ans := part2(input)
		util.CopyToClipboard(fmt.Sprintf("%v", ans))
		fmt.Println("Output:", ans)
	}
}

func isSafe(row []int) bool {

	ascending := true
	for i, num := range row {
		if i == 0 && len(row) > 1 {
			if num > row[i+1] {
				ascending = false
			}
		}
		if i < len(row)-1 {
			value := num - row[i+1]
			if ascending {
				if value != -1 && value != -2 && value != -3 {
					return false
				}
			} else {
				if value != 1 && value != 2 && value != 3 {
					return false
				}
			}
		}
	}

	return true
}

func isWhoopsieSafe(report []int) bool {
	inc := 0
	for idx := 1; idx < len(report); idx++ {
		if report[idx]-report[idx-1] > 0 {
			inc++
		} else {
			inc--
		}
	}
	if inc < len(report)-3 && inc > -(len(report)-3) {
		return false // there are more than one bad number
	}
	isInc := false
	if inc > 0 {
		isInc = true
	}
	for idx := 1; idx < len(report); idx++ {
		if isInc && (report[idx]-report[idx-1]) > 0 && (report[idx]-report[idx-1]) < 4 {
			continue
		}
		if !isInc && (report[idx]-report[idx-1]) < 0 && (report[idx]-report[idx-1]) > -4 {
			continue
		}
		// If there is a wrong level check current one and adjacent
		newReport := slices.Clone(report)
		newReport = append(newReport[:idx], newReport[idx+1:]...)
		if isSafe(newReport) {
			return true
		} else {
			newReport = slices.Clone(report)
			newReport = append(newReport[:idx-1], newReport[idx:]...)
			if isSafe(newReport) {
				return true
			} else {
				newReport = slices.Clone(report)
				if idx < len(report)-2 {
					newReport = append(newReport[:idx+1], newReport[idx+2:]...)
				} else {
					newReport = newReport[:len(newReport)-1]
				}
				if isSafe(newReport) {
					return true
				}
			}
		}
		return false
	}
	return true
}

func part1(input string) int {
	parsed := parseInput(input)

	no_of_safe := 0
	for _, row := range parsed {
		if isSafe(row) {
			no_of_safe++
		}
	}

	return no_of_safe
}

func part2(input string) int {
	parsed := parseInput(input)

	no_of_safe := 0
	for _, row := range parsed {
		r := isWhoopsieSafe(row)
		if r {
			no_of_safe++
		}
	}

	return no_of_safe
}

func parseInput(input string) (ans [][]int) {
	ans = [][]int{}
	for _, line := range strings.Split(input, "\n") {
		row := []int{}
		for _, num := range strings.Split(line, " ") {
			if num == "" {
				continue
			}
			row = append(row, cast.ToInt(num))
		}
		ans = append(ans, row)
	}
	return ans
}
