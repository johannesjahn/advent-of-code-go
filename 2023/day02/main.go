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

type ball struct {
	color string
	count int
}

type turn struct {
	balls []ball
}

type game struct {
	id    int
	turns []turn
}

func parseGames(input []string) []game {
	games := []game{}
	for _, inputStr := range input {
		game := game{-1, nil}
		relevant := strings.Split(inputStr, ":")
		game.id = util.MustAtoi(strings.Split(relevant[0], " ")[1])
		turns := relevant[1]
		for _, turnInput := range strings.Split(turns, ";") {
			turnR := turn{}
			for _, balls := range strings.Split(turnInput, ",") {
				b := ball{}
				for ncIdx, numberColor := range strings.Split(balls, " ") {
					if ncIdx == 0 {
						continue
					}
					if ncIdx%2 == 0 {
						b.color = numberColor
					} else {
						b.count = util.MustAtoi(numberColor)
					}
				}
				turnR.balls = append(turnR.balls, b)
			}
			game.turns = append(game.turns, turnR)
		}
		games = append(games, game)
	}
	return games
}

func part1(input string) int {
	parsed := parseInput(input)
	_ = parsed

	games := parseGames(parsed)

	redPossible := 12
	greenPossible := 13
	bluePossible := 14

	possibleIdsSum := 0
	for _, game := range games {
		possible := true
		for _, turn := range game.turns {
			if !possible {
				break
			}
			for _, ball := range turn.balls {
				if !possible {
					break
				}
				switch color := ball.color; color {
				case "red":
					possible = ball.count <= redPossible
				case "green":
					possible = ball.count <= greenPossible
				case "blue":
					possible = ball.count <= bluePossible
				}
			}
		}
		if possible {
			possibleIdsSum += game.id
		}
	}

	return possibleIdsSum
}

func part2(input string) int {

	parsed := parseInput(input)
	games := parseGames(parsed)

	result := 0

	for _, game := range games {
		maxRed := 1
		maxGreen := 1
		maxBlue := 1

		for _, turn := range game.turns {
			for _, ball := range turn.balls {
				switch color := ball.color; color {
				case "red":
					maxRed = max(maxRed, ball.count)
				case "green":
					maxGreen = max(maxGreen, ball.count)
				case "blue":
					maxBlue = max(maxBlue, ball.count)
				}
			}
		}
		result += maxRed * maxGreen * maxBlue
	}

	return result
}

func parseInput(input string) (ans []string) {
	for _, line := range strings.Split(input, "\n") {
		ans = append(ans, line)
	}
	return ans
}
