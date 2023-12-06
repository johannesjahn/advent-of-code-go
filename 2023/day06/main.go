package main

import (
	_ "embed"
	"flag"
	"fmt"
	"regexp"
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

func part1(input string) int {
	times, distances := parseInput(input)
	result := 1

	for roundIdx := 0; roundIdx < len(times); roundIdx++ {
		numberOfWays := 0
		time := times[roundIdx]
		goalDistance := distances[roundIdx]
		for chargeTime := 1; chargeTime < time; chargeTime++ {
			distance := chargeTime * (time - chargeTime)
			if distance > goalDistance {
				numberOfWays++
			}
		}
		result *= numberOfWays
	}

	return result
}

func part2(input string) int {

	times, distances := parseInput(input)
	timeStr := ""
	for _, time := range times {
		timeStr += fmt.Sprintf("%d", time)
	}
	time := cast.ToInt(timeStr)

	distanceStr := ""
	for _, dist := range distances {
		distanceStr += fmt.Sprintf("%d", dist)
	}
	goalDistance := cast.ToInt(distanceStr)

	numberOfWays := 0
	for chargeTime := 1; chargeTime < time; chargeTime++ {
		distance := chargeTime * (time - chargeTime)
		if distance > goalDistance {
			numberOfWays++
		}
	}

	return numberOfWays
}

func parseInput(input string) (times []int, distances []int) {
	for idx, line := range strings.Split(input, "\n") {
		if idx == 0 {
			for _, time := range regexp.MustCompile(`\d+`).FindAllString(line, -1) {
				times = append(times, cast.ToInt(time))
			}
		} else if idx == 1 {
			for _, dist := range regexp.MustCompile(`\d+`).FindAllString(line, -1) {
				distances = append(distances, cast.ToInt(dist))
			}
		}
	}
	return times, distances
}
