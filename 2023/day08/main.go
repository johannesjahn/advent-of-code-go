package main

import (
	_ "embed"
	"flag"
	"fmt"
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

type instruction struct {
	op    string
	left  string
	right string
}

func parseLine(line string) instruction {
	re := regexp.MustCompile(`(?m)(\S+) = \((\S+), (\S+)\)`).FindStringSubmatch(line)
	return instruction{re[1], re[2], re[3]}
}

func part1(input string) int {
	parsed := parseInput(input)
	_ = parsed

	instructions := parsed[0]

	instructionMap := map[string]instruction{}

	for i := 2; i < len(parsed); i++ {
		instruction := parseLine(parsed[i])
		instructionMap[instruction.op] = instruction
	}

	executionCount := 0
	currentInstruction := instructionMap["AAA"]
	currentInstructionIdx := 0

	for true {
		instruction := instructions[currentInstructionIdx]
		if string(currentInstruction.op) == "ZZZ" {
			return executionCount
		}
		executionCount++
		if string(instruction) == "R" {
			currentInstruction = instructionMap[currentInstruction.right]
		} else {
			currentInstruction = instructionMap[currentInstruction.left]
		}
		currentInstructionIdx = (currentInstructionIdx + 1) % len(instructions)
	}

	return 0
}

func part2(input string) int { // TODO: fix this
	parsed := parseInput(input)
	_ = parsed

	instructions := parsed[0]

	instructionMap := map[string]instruction{}
	instructionRunners := make([]string, 0)

	for i := 2; i < len(parsed); i++ {
		instruction := parseLine(parsed[i])
		instructionMap[instruction.op] = instruction
		if instruction.op[2] == 'A' {
			instructionRunners = append(instructionRunners, instruction.op)
		}
	}

	executionCounts := make([]int, len(instructionRunners))
	currentInstructionIdx := 0

	for i := 0; i < len(instructionRunners); i++ {
		executionCount := 0
		for {
			if instructionRunners[i][2] == 'Z' {
				break
			}

			instruction := instructions[currentInstructionIdx]
			if string(instruction) == "R" {
				instructionRunners[i] = instructionMap[instructionRunners[i]].right
			} else {
				instructionRunners[i] = instructionMap[instructionRunners[i]].left
			}
			executionCount++
			currentInstructionIdx = (currentInstructionIdx + 1) % len(instructions)
		}
		executionCounts[i] = executionCount
	}

	return lcm(executionCounts[0], executionCounts[1], executionCounts[2:]...)
}

func gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func lcm(a, b int, integers ...int) int {
	result := a * b / gcd(a, b)

	for i := 0; i < len(integers); i++ {
		result = lcm(result, integers[i])
	}

	return result
}

func allInstructionsAreTerminal(instructions []string) bool {
	for _, instruction := range instructions {
		if instruction[2] != 'Z' {
			return false
		}
	}
	return true
}

func parseInput(input string) (ans []string) {
	for _, line := range strings.Split(input, "\n") {
		ans = append(ans, line)
	}
	return strings.Split(input, "\n")
}
