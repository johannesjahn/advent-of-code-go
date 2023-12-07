package main

import (
	_ "embed"
	"flag"
	"fmt"
	"sort"
	"strings"

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
	turns := parseInput(input)

	sort.SliceStable(turns, func(i, j int) bool {
		s1 := score(turns[i].hand)
		s2 := score(turns[j].hand)

		if s1 == s2 {
			return handIsBiggerByChar(turns[i].hand, turns[j].hand)
		}

		return s1 >= s2
	})

	result := 0
	for i := 0; i < len(turns); i++ {
		result += turns[i].bid * (len(turns) - i)
	}

	return result
}

func convertCardToValue(c string) int {
	if c == "T" {
		return 10
	}
	if c == "J" {
		return 11
	}
	if c == "Q" {
		return 12
	}
	if c == "K" {
		return 13
	}
	if c == "A" {
		return 14
	}
	return util.MustAtoi(c)
}

func convertCardToValueWithJoker(c string) int {
	if c == "T" {
		return 10
	}
	if c == "J" {
		return 1
	}
	if c == "Q" {
		return 12
	}
	if c == "K" {
		return 13
	}
	if c == "A" {
		return 14
	}
	return util.MustAtoi(c)
}

func handIsBiggerByCharWithJoker(hand1, hand2 string) bool {
	for i := 0; i < len(hand1); i++ {

		val1 := convertCardToValueWithJoker(string(hand1[i]))
		val2 := convertCardToValueWithJoker(string(hand2[i]))

		if val1 > val2 {
			return true
		}
		if val2 > val1 {
			return false
		}
	}
	return true
}

func handIsBiggerByChar(hand1, hand2 string) bool {
	for i := 0; i < len(hand1); i++ {

		val1 := convertCardToValue(string(hand1[i]))
		val2 := convertCardToValue(string(hand2[i]))

		if val1 > val2 {
			return true
		}
		if val2 > val1 {
			return false
		}
	}
	return true
}

func score(hand string) int {
	valueMap := make(map[string]int)
	for _, c := range hand {
		valueMap[string(c)]++
	}
	values := make([]int, 0, len(valueMap))
	for _, v := range valueMap {
		values = append(values, v)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(values)))

	if values[0] == 5 {
		return 7
	}
	if values[0] == 4 {
		return 6
	}
	if values[0] == 3 && values[1] == 2 {
		return 5
	}
	if values[0] == 3 {
		return 4
	}
	if values[0] == 2 && values[1] == 2 {
		return 3
	}
	if values[0] == 2 {
		return 2
	}
	return 1
}

func scoreWithJoker(hand string) int {
	valueMap := make(map[string]int)
	for _, c := range hand {
		valueMap[string(c)]++
	}

	jokers := valueMap["J"]
	if jokers == 5 {
		return 7
	}
	delete(valueMap, "J")

	values := make([]int, 0, len(valueMap))
	for _, v := range valueMap {
		values = append(values, v)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(values)))

	values[0] += jokers

	if values[0] == 5 {
		return 7
	}
	if values[0] == 4 {
		return 6
	}
	if values[0] == 3 && values[1] == 2 {
		return 5
	}
	if values[0] == 3 {
		return 4
	}
	if values[0] == 2 && values[1] == 2 {
		return 3
	}
	if values[0] == 2 {
		return 2
	}
	return 1
}

func part2(input string) int {
	turns := parseInput(input)

	sort.SliceStable(turns, func(i, j int) bool {
		s1 := scoreWithJoker(turns[i].hand)
		s2 := scoreWithJoker(turns[j].hand)

		if s1 == s2 {
			return handIsBiggerByCharWithJoker(turns[i].hand, turns[j].hand)
		}

		return s1 >= s2
	})

	result := 0
	for i := 0; i < len(turns); i++ {
		result += turns[i].bid * (len(turns) - i)
	}

	return result
}

type turn struct {
	hand string
	bid  int
}

func parseInput(input string) (turns []turn) {
	for _, line := range strings.Split(input, "\n") {
		r := strings.Split(line, " ")
		turns = append(turns, turn{r[0], util.MustAtoi(r[1])})
	}
	return turns
}
