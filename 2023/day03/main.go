package main

import (
	_ "embed"
	"flag"
	"fmt"
	"strings"
	"unicode"

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

func part1(input string) int {
	parsed := parseInput(input)
	sum := 0

	for i, line := range parsed {
		startIdx := -1
		for j, char := range line {
			if unicode.IsDigit(char) {
				if startIdx == -1 {
					startIdx = j
				}
				if j == len(line)-1 {
					if isAdjacent(startIdx, j, i, parsed) {
						sum += util.MustAtoi(string(line[startIdx : j+1]))
					}
				}
			} else if startIdx != -1 {
				if isAdjacent(startIdx, j-1, i, parsed) {
					sum += util.MustAtoi(string(line[startIdx:j]))
				}
				startIdx = -1
			}
		}
	}

	return sum
}

func isAdjacent(startIdx int, endIdx int, lineIdx int, str []string) bool {

	start := startIdx - 1
	if start < 0 {
		start = 0
	}

	idxAbove := lineIdx - 1
	if idxAbove >= 0 {
		end := endIdx + 2
		if end > len(str[idxAbove]) {
			end = len(str[idxAbove])
		}
		if !containsOnlyDot(str[idxAbove][start:end]) {
			return true
		}
	}
	left := startIdx - 1
	if left >= 0 {
		if !containsOnlyDot(string(str[lineIdx][left])) {
			return true
		}
	}
	right := endIdx + 1
	if right <= len(str[lineIdx])-1 {
		if !containsOnlyDot(string(str[lineIdx][right])) {
			return true
		}
	}
	idxBelow := lineIdx + 1
	if idxBelow <= len(str)-1 {
		end := endIdx + 2
		if end > len(str[idxBelow]) {
			end = len(str[idxBelow])
		}
		if !containsOnlyDot(str[idxBelow][start:end]) {
			return true
		}
	}
	return false
}

func containsOnlyDot(str string) bool {
	for _, char := range str {
		if char != '.' {
			return false
		}
	}
	return true
}

func part2(input string) int {
	return 0
}

func parseInput(input string) (ans []string) {
	for _, line := range strings.Split(input, "\n") {
		ans = append(ans, line)
	}
	return ans
}
