package main

import (
	_ "embed"
	"flag"
	"fmt"
	"strings"

	"github.com/alexchao26/advent-of-code-go/util"
)

type coord struct {
	i, j int
}

type dir coord

var (
	dirNorth = dir{-1, 0}
	dirSouth = dir{1, 0}
	dirWest  = dir{0, -1}
	dirEast  = dir{0, 1}
)

var pipes = map[byte]map[dir]dir{
	'|': {
		dirNorth: dirNorth,
		dirSouth: dirSouth,
	},
	'-': {
		dirEast: dirEast,
		dirWest: dirWest,
	},
	'L': {
		dirSouth: dirEast,
		dirWest:  dirNorth,
	},
	'J': {
		dirEast:  dirNorth,
		dirSouth: dirWest,
	},
	'7': {
		dirEast:  dirSouth,
		dirNorth: dirWest,
	},
	'F': {
		dirNorth: dirEast,
		dirWest:  dirSouth,
	},
}

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

	y, x := findStart(parsed)

	ys, xs := findNext(parsed, y, x, -1, -1)

	distance := 1

	for {
		oldY, oldX := ys, xs
		ys, xs = findNext(parsed, ys, xs, y, x)
		distance++
		y, x = oldY, oldX

		if parsed[ys][xs] == 'S' {
			break
		}
	}

	return distance / 2
}

func findStart(input []string) (row, col int) {
	for row, line := range input {
		for col, char := range line {
			if char == 'S' {
				return row, col
			}
		}
	}
	panic("no start found")
}

func findNext(input []string, y, x, preY, preX int) (nextRow, nextCol int) {
	current := input[y][x]
	switch current {
	case 'S':
		// find the next char
		if x+1 < len(input[y]) && input[y][x+1] != '.' {
			return y, x + 1
		} else if x-1 >= 0 && input[y][x-1] != '.' {
			return y, x - 1
		} else if y-1 >= 0 && input[y-1][x] != '.' {
			return y - 1, x
		} else if y+1 < len(input) && input[y+1][x] != '.' {
			return y + 1, x
		}
	case '-':
		if preX == x+1 {
			return y, x - 1
		} else {
			return y, x + 1
		}
	case '7':
		if preY == y {
			return y + 1, x
		} else {
			return y, x - 1
		}
	case '|':
		if preY == y+1 {
			return y - 1, x
		} else {
			return y + 1, x
		}
	case 'J':
		if preX == x {
			return y, x - 1
		} else {
			return y - 1, x
		}
	case 'L':
		if preY == y {
			return y - 1, x
		} else {
			return y, x + 1
		}
	case 'F':
		if preX == x {
			return y, x + 1
		} else {
			return y + 1, x
		}
	}
	panic("no next found")
}

func findSStart(grid [][]byte) coord {
	var s coord

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 'S' {
				s.i = i
				s.j = j
				return s
			}
		}
	}

	panic("no S found in grid")
}

func getGrid(input string) [][]byte {
	splitted := strings.Split(input, "\n")
	grid := make([][]byte, len(splitted))
	for i := range splitted {
		grid[i] = []byte(splitted[i])
	}

	return grid
}

func part2(input string) int {
	grid := getGrid(input)
	s := findSStart(grid)
	loop := findLoop(s, grid)

	// https://en.wikipedia.org/wiki/Shoelace_formula
	polygonArea := 0
	for i := 0; i < len(loop); i++ {
		cur := loop[i]
		next := loop[(i+1)%len(loop)]

		polygonArea += cur.i*next.j - cur.j*next.i
	}

	if polygonArea < 0 {
		polygonArea = -polygonArea
	}
	polygonArea /= 2

	// https://en.wikipedia.org/wiki/Pick%27s_theorem
	return polygonArea - len(loop)/2 + 1
}

func findLoop(s coord, grid [][]byte) []coord {
	for _, pipe := range "|-LJ7F" {
		grid[s.i][s.j] = byte(pipe)
		loop := checkLoop(s, grid)
		if loop != nil {
			return loop
		}
	}

	panic("no loop found")
}

func anyKey(m map[dir]dir) dir {
	for k := range m {
		return k
	}

	panic("empty map")
}

func checkLoop(s coord, grid [][]byte) []coord {
	cur := s
	dir := anyKey(pipes[grid[s.i][s.j]])

	res := []coord{}

	for {
		res = append(res, cur)
		newDir, ok := pipes[grid[cur.i][cur.j]][dir]
		if !ok {
			return nil
		}

		newCoord := coord{cur.i + newDir.i, cur.j + newDir.j}

		if newCoord.i < 0 || newCoord.i >= len(grid) || newCoord.j < 0 || newCoord.j >= len(grid[newCoord.i]) {
			return nil
		}
		if newCoord == s {
			if _, ok := pipes[grid[s.i][s.j]][newDir]; !ok {
				return nil
			}
			break
		}
		cur = newCoord
		dir = newDir
	}

	return res
}

func myPart2(input string) int {

	parsed := parseInput(input)
	enclosed := 0

	for x := 0; x < len(parsed[0]); x++ {
		for y := 0; y < len(parsed); y++ {

			if parsed[y][x] != '.' && !isConnected(parsed, y, x) {
				continue
			}

			noCrossWest := 0
			lastPipeturn := ""
			didCross := false
			for i := 0; i < x; i++ {
				didCross, lastPipeturn = getDidCrossAndLastPipeturnHorizontal(parsed, y, i, lastPipeturn)
				if didCross {
					noCrossWest++
				}
			}

			noCrossEast := 0
			didCross = false
			lastPipeturn = ""
			for i := x + 1; i < len(parsed[y]); i++ {
				didCross, lastPipeturn = getDidCrossAndLastPipeturnHorizontal(parsed, y, i, lastPipeturn)
				if didCross {
					noCrossEast++
				}
			}

			noCrossNorth := 0
			didCross = false
			lastPipeturn = ""
			for i := 0; i < y; i++ {
				didCross, lastPipeturn = getDidCrossAndLastPipeturnVertical(parsed, i, x, lastPipeturn)
				if didCross {
					noCrossNorth++
				}
			}

			noCrossSouth := 0
			didCross = false
			lastPipeturn = ""
			for i := y + 1; i < len(parsed); i++ {
				didCross, lastPipeturn = getDidCrossAndLastPipeturnVertical(parsed, i, x, lastPipeturn)
				if didCross {
					noCrossSouth++
				}
			}

			if y == 4 && x == 10 {
				fmt.Println(noCrossWest, noCrossEast, noCrossNorth, noCrossSouth)
				fmt.Println(parsed[y])
			}

			if noCrossWest%2 == 1 && noCrossEast%2 == 1 && noCrossNorth%2 == 1 && noCrossSouth%2 == 1 {
				enclosed++
			}

		}
	}

	return enclosed
}

func getDidCrossAndLastPipeturnHorizontal(parsed []string, y, x int, lastPipeturn string) (didCross bool, newlastPipeturn string) {
	if parsed[y][x] == '|' {
		return true, ""
	} else if parsed[y][x] == 'L' {
		if lastPipeturn == "7" {
			return true, ""
		} else {
			return false, "L"
		}
	} else if parsed[y][x] == '7' {
		if lastPipeturn == "L" {
			return true, ""
		} else {
			return false, "7"
		}
	} else if parsed[y][x] == 'J' {
		if lastPipeturn == "F" {
			return true, ""
		} else {
			return false, "J"
		}
	} else if parsed[y][x] == 'F' {
		if lastPipeturn == "J" {
			return true, ""
		} else {
			return false, "F"
		}
	} else if parsed[y][x] == '-' {
		return false, lastPipeturn
	}
	return false, ""
}

func getDidCrossAndLastPipeturnVertical(parsed []string, y, x int, lastPipeturn string) (didCross bool, newlastPipeturn string) {
	if parsed[y][x] == '-' {
		return true, ""
	} else if parsed[y][x] == 'L' {
		if lastPipeturn == "7" {
			return true, ""
		} else {
			return false, "L"
		}
	} else if parsed[y][x] == '7' {
		if lastPipeturn == "L" {
			return true, ""
		} else {
			return false, "7"
		}
	} else if parsed[y][x] == 'J' {
		if lastPipeturn == "F" {
			return true, ""
		} else {
			return false, "J"
		}
	} else if parsed[y][x] == 'F' {
		if lastPipeturn == "J" {
			return true, ""
		} else {
			return false, "F"
		}
	} else if parsed[y][x] == '|' {
		return false, lastPipeturn
	}
	return false, ""
}

func isConnected(parsed []string, y, x int) bool {
	if parsed[y][x] == 'S' {
		return true
	}
	if parsed[y][x] == '.' {
		return false
	}

	if parsed[y][x] == '-' {
		if x == 0 {
			return false
		}
		if x == len(parsed[y])-1 {
			return false
		}
		if (parsed[y][x-1] == 'F' ||
			parsed[y][x-1] == 'L' ||
			parsed[y][x-1] == '-') &&
			(parsed[y][x+1] == 'J' ||
				parsed[y][x+1] == '7' ||
				parsed[y][x+1] == '-') {
			return true
		}
	}
	if parsed[y][x] == '|' {
		if y == 0 {
			return false
		}
		if y == len(parsed)-1 {
			return false
		}
		if (parsed[y-1][x] == 'F' ||
			parsed[y-1][x] == '7' ||
			parsed[y-1][x] == '|') &&
			(parsed[y+1][x] == 'J' ||
				parsed[y+1][x] == 'L' ||
				parsed[y+1][x] == '|') {
			return true
		}
	}
	if parsed[y][x] == 'F' {
		if y == len(parsed)-1 {
			return false
		}
		if x == len(parsed[y])-1 {
			return false
		}
		if (parsed[y][x+1] == 'J' ||
			parsed[y][x+1] == '7' ||
			parsed[y][x+1] == '-') &&
			(parsed[y+1][x] == 'J' ||
				parsed[y+1][x] == 'L' ||
				parsed[y+1][x] == '|') {
			return true
		}
	}
	if parsed[y][x] == 'J' {
		if x == 0 {
			return false
		}
		if y == len(parsed)-1 {
			return false
		}
		if (parsed[y][x-1] == 'F' ||
			parsed[y][x-1] == 'L' ||
			parsed[y][x-1] == '-') &&
			(parsed[y+1][x] == 'F' ||
				parsed[y+1][x] == '7' ||
				parsed[y+1][x] == '|') {
			return true
		}
	}
	if parsed[y][x] == 'L' {
		if y == 0 {
			return false
		}
		if x == len(parsed[y])-1 {
			return false
		}
		if (parsed[y][x+1] == '7' ||
			parsed[y][x+1] == 'J' ||
			parsed[y][x+1] == '-') &&
			(parsed[y-1][x] == 'F' ||
				parsed[y-1][x] == '7' ||
				parsed[y-1][x] == '|') {
			return true
		}
	}
	if parsed[y][x] == '7' {
		if x == 0 {
			return false
		}
		if y == len(parsed)-1 {
			return false
		}
		if (parsed[y][x-1] == 'F' ||
			parsed[y][x-1] == 'L' ||
			parsed[y][x-1] == '-') &&
			(parsed[y+1][x] == 'J' ||
				parsed[y+1][x] == 'L' ||
				parsed[y+1][x] == '|') {
			return true
		}
	}
	return false
}

func parseInput(input string) (ans []string) {
	for _, line := range strings.Split(input, "\n") {
		ans = append(ans, line)
	}
	return ans
}
