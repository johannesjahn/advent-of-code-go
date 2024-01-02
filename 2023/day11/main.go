package main

import (
	_ "embed"
	"flag"
	"fmt"
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

type galaxy struct {
	x int
	y int
}

func part1(input string) int {
	parsed := parseInput(input)
	_ = parsed

	rowsContainingGalaxies := make(map[int]bool)
	colsContainingGalaxies := make(map[int]bool)

	for x, line := range parsed {
		for y, char := range line {
			if char == '#' {
				rowsContainingGalaxies[x] = true
				colsContainingGalaxies[y] = true
			}
		}
	}

	rowsExpanded := 0
	rowsCount := len(parsed)
	for x := 0; x < rowsCount; x++ {
		if !rowsContainingGalaxies[x] {
			parsed = expandRow(parsed, x+rowsExpanded)
			rowsExpanded++
		}
	}

	colsExpanded := 0
	colsCount := len(parsed[0])
	for y := 0; y < colsCount; y++ {
		if !colsContainingGalaxies[y] {
			parsed = expandCol(parsed, y+colsExpanded)
			colsExpanded++
		}
	}

	galaxies := make([]galaxy, 0)

	for i, line := range parsed {
		for j, char := range line {
			if char == '#' {
				galaxies = append(galaxies, galaxy{j, i})
			}
		}
	}

	result := 0

	for i := 0; i < len(galaxies)-1; i++ {
		for j := i + 1; j < len(galaxies); j++ {
			result += calculateDistance(galaxies[i], galaxies[j])
		}
	}

	return result
}

func calculateDistanceSpec(g1, g2 galaxy, rowsContainingGalaxies, colsContainingGalaxies map[int]bool) int {

	factor := 1000000

	result := 0

	a := 0
	b := 0

	if g1.x < g2.x {
		a = g1.x
		b = g2.x
	} else {
		a = g2.x
		b = g1.x
	}

	for i := a; i < b; i++ {
		if colsContainingGalaxies[i] {
			result++
		} else {
			result += factor
		}
	}

	if g1.y < g2.y {
		a = g1.y
		b = g2.y
	} else {
		a = g2.y
		b = g1.y
	}

	for i := a; i < b; i++ {
		if rowsContainingGalaxies[i] {
			result++
		} else {
			result += factor
		}
	}

	return result
}

func calculateDistance(g1, g2 galaxy) int {

	a := g1.x - g2.x
	b := g1.y - g2.y

	if a < 0 {
		a = -a
	}
	if b < 0 {
		b = -b
	}
	return a + b
}

func expandRow(input []string, row int) []string {
	result := append(input[0:row], strings.Repeat(".", len(input[0])))
	result = append(result, input[row:]...)
	return result
}

func expandCol(input []string, col int) []string {
	result := make([]string, len(input))
	for i, line := range input {
		result[i] = line[0:col] + "." + line[col:]
	}
	return result
}

func part2(input string) int {
	parsed := parseInput(input)
	_ = parsed

	rowsContainingGalaxies := make(map[int]bool)
	colsContainingGalaxies := make(map[int]bool)

	for x, line := range parsed {
		for y, char := range line {
			if char == '#' {
				rowsContainingGalaxies[x] = true
				colsContainingGalaxies[y] = true
			}
		}
	}

	rowsCount := len(parsed)
	for x := 0; x < rowsCount; x++ {
		if !rowsContainingGalaxies[x] {
			rowsContainingGalaxies[x] = false
		}
	}

	colsCount := len(parsed[0])
	for y := 0; y < colsCount; y++ {
		if !colsContainingGalaxies[y] {
			colsContainingGalaxies[y] = false
		}
	}

	galaxies := make([]galaxy, 0)

	for i, line := range parsed {
		for j, char := range line {
			if char == '#' {
				galaxies = append(galaxies, galaxy{j, i})
			}
		}
	}

	result := 0

	for i := 0; i < len(galaxies)-1; i++ {
		for j := i + 1; j < len(galaxies); j++ {
			result += calculateDistanceSpec(galaxies[i], galaxies[j], rowsContainingGalaxies, colsContainingGalaxies)
		}
	}

	return result
}

func parseInput(input string) (ans []string) {
	ans = append(ans, strings.Split(input, "\n")...)
	return ans
}
