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

	parsed := parseInput(input)
	sum := 0

	for i, line := range parsed {
		for j, char := range line {
			if char == '*' {
				numbers := findAllAdjacentNumbers(i, j, parsed)
				if len(numbers) > 1 {
					ans := numbers[0]
					for k := 1; k < len(numbers); k++ {
						ans = ans * numbers[k]
					}
					sum += ans
				}
			}
		}
	}
	return sum
}

func findAllAdjacentNumbers(i int, j int, strs []string) []int {
	numbers := []int{}
	number := -1

	// check above
	idxLineAbove := i - 1
	if idxLineAbove >= 0 {
		if !unicode.IsDigit(rune(strs[idxLineAbove][j])) {
			if j-1 >= 0 {
				number = getNumber(idxLineAbove, j-1, strs)
				if number != -1 {
					numbers = append(numbers, number)
				}
			}
			if j+1 < len(strs[idxLineAbove]) {
				number = getNumber(idxLineAbove, j+1, strs)
				if number != -1 {
					numbers = append(numbers, number)
				}
			}
		} else {
			number = getNumber(idxLineAbove, j, strs)
			numbers = append(numbers, number)
		}
	}

	// check left
	if j-1 >= 0 {
		number = getNumber(i, j-1, strs)
		if number != -1 {
			numbers = append(numbers, number)
		}
	}

	// check right
	if j+1 < len(strs[i]) {
		number = getNumber(i, j+1, strs)
		if number != -1 {
			numbers = append(numbers, number)
		}
	}

	// check below
	idxLineBelow := i + 1
	if idxLineBelow < len(strs) {
		if !unicode.IsDigit(rune(strs[idxLineBelow][j])) {
			if j-1 >= 0 {
				number = getNumber(idxLineBelow, j-1, strs)
				if number != -1 {
					numbers = append(numbers, number)
				}
			}
			if j+1 < len(strs[idxLineBelow]) {
				number = getNumber(idxLineBelow, j+1, strs)
				if number != -1 {
					numbers = append(numbers, number)
				}
			}
		} else {
			number = getNumber(idxLineBelow, j, strs)
			numbers = append(numbers, number)
		}
	}

	return numbers
}

func getNumber(i int, j int, strs []string) int {

	result := []rune{}

	if !unicode.IsDigit(rune(strs[i][j])) {
		return -1
	}

	result = append(result, rune(strs[i][j]))

	for x := j - 1; x >= 0; x-- {
		if !unicode.IsDigit(rune(strs[i][x])) {
			break
		}
		result = append([]rune{rune(strs[i][x])}, result...)
	}
	for x := j + 1; x < len(strs[i]); x++ {
		if !unicode.IsDigit(rune(strs[i][x])) {
			break
		}
		result = append(result, rune(strs[i][x]))
	}

	return util.MustAtoi(string(result))
}

func parseInput(input string) (ans []string) {
	for _, line := range strings.Split(input, "\n") {
		ans = append(ans, line)
	}
	return ans
}
