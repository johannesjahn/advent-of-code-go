package main

import (
	_ "embed"
	"flag"
	"fmt"
	"strconv"
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
	result := 0

	for _, line := range parsed {

		first := "0"
		for _, char := range line {
			if unicode.IsDigit(char) {
				first = string(char)
				break
			}
		}
		last := "0"
		for i := len(line) - 1; i >= 0; i-- {
			if unicode.IsDigit(rune(line[i])) {
				last = string(line[i])
				break
			}
		}
		result += mustAtoi(first + last)
	}
	return result
}

func part2(input string) int {
	parsed := parseInput(input)
	result := 0

	for _, line := range parsed {

		first := "0"
		for i, char := range line {
			if unicode.IsDigit(char) {
				first = string(char)
				break
			}
			chk := checkIfNumberString(line, i)
			if chk != -1 {
				first = strconv.Itoa(chk)
				break
			}
		}
		last := "0"
		for i := len(line) - 1; i >= 0; i-- {
			if unicode.IsDigit(rune(line[i])) {
				last = string(line[i])
				break
			}
			chk := checkIfNumberString(line, i)
			if chk != -1 {
				last = strconv.Itoa(chk)
				break
			}
		}
		result += mustAtoi(first + last)
	}
	return result
}

func parseInput(input string) (ans []string) {
	for _, line := range strings.Split(input, "\n") {
		ans = append(ans, line)
	}
	return ans
}

func checkIfNumberString(s string, index int) int {

	if index+3 <= len(s) {
		if s[index:index+3] == "one" {
			return 1
		}
		if s[index:index+3] == "two" {
			return 2
		}
		if s[index:index+3] == "six" {
			return 6
		}
	}
	if index+4 <= len(s) {
		if s[index:index+4] == "four" {
			return 4
		}
		if s[index:index+4] == "five" {
			return 5
		}
		if s[index:index+4] == "nine" {
			return 9
		}
	}
	if index+5 <= len(s) {
		if s[index:index+5] == "three" {
			return 3
		}
		if s[index:index+5] == "seven" {
			return 7
		}
		if s[index:index+5] == "eight" {
			return 8
		}
	}
	return -1
}

func mustAtoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}
