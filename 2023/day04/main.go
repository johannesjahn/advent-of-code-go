package main

import (
	_ "embed"
	"flag"
	"fmt"
	"math"
	"regexp"
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
	parsed := parseInput(input)

	var re = regexp.MustCompile(`(?m)\d+`)

	result := float64(0)
	for _, game := range parsed {
		v := strings.Split(strings.Split(game, ":")[1], "|")
		winnings := v[0]
		nos := v[1]
		r := 0

		wMap := map[int]bool{}
		for _, iw := range re.FindAllString(winnings, -1) {
			wMap[util.MustAtoi(iw)] = true
		}
		for _, in := range re.FindAllString(nos, -1) {
			_, ok := wMap[util.MustAtoi(in)]
			if ok {
				r++
			}
		}

		if r != 0 {
			r -= 1
			result += math.Pow(2, float64(r))
		}
	}

	return int(result)
}

func part2(input string) int {
	parsed := parseInput(input)

	var re = regexp.MustCompile(`(?m)\d+`)

	result := 0
	numberOfWins := []int{}
	for _, game := range parsed {
		v := strings.Split(strings.Split(game, ":")[1], "|")
		winnings := v[0]
		nos := v[1]
		r := 0

		wMap := map[int]bool{}
		for _, iw := range re.FindAllString(winnings, -1) {
			wMap[util.MustAtoi(iw)] = true
		}
		for _, in := range re.FindAllString(nos, -1) {
			_, ok := wMap[util.MustAtoi(in)]
			if ok {
				r++
			}
		}

		numberOfWins = append(numberOfWins, r)
	}

	numberOfCards := []int{}
	for range parsed {
		numberOfCards = append(numberOfCards, 1)
	}

	for idx, _ := range numberOfWins {
		if idx == 0 {
			continue
		}
		numberOfCardsAtIdx := 1
		for i := 0; i < idx; i++ {
			if numberOfWins[i]+i >= idx {
				numberOfCardsAtIdx += 1 * numberOfCards[i]
			}
		}
		numberOfCards[idx] = numberOfCardsAtIdx
	}

	for _, c := range numberOfCards {
		result += c
	}

	return result
}

func parseInput(input string) (ans []string) {
	for _, line := range strings.Split(input, "\n") {
		ans = append(ans, line)
	}
	return ans
}
