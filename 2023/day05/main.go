package main

import (
	_ "embed"
	"flag"
	"fmt"
	"math"
	"strconv"
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
	seeds, mappers := parseInput(input)
	var minLocation uint64 = math.MaxUint64
	for _, seed := range seeds {
		minLocation = min(minLocation, mappers.mapFrom(seed))
	}
	return int(minLocation)
}

func part2(input string) any {
	seeds, mappers := parseInput(input)
	var minLocation uint64 = math.MaxUint64
	for i := 0; i < len(seeds); i += 2 {
		end := seeds[i] + seeds[i+1]
		for s := seeds[i]; s < end; s++ {
			minLocation = min(minLocation, mappers.mapFrom(s))
		}
	}
	return int(minLocation)
}

func parseInput(input string) ([]uint64, multiMapper) {
	lines := strings.Split(input, "\n")
	seedStrings := strings.Split(lines[0], " ")
	seeds := make([]uint64, 0, len(seedStrings)-1)
	for _, seed := range seedStrings[1:] {
		s, _ := strconv.ParseUint(seed, 10, 64)
		seeds = append(seeds, s)
	}
	mappers := make([]mapper, 0, 7)
	for _, line := range lines[1:] {
		if line == "" {
			mappers = append(mappers, mapper{})
			continue
		}
		parts := strings.Split(line, " ")
		if len(parts) == 2 {
			// is the identity of the map; don't care
			continue
		}
		dst, _ := strconv.ParseUint(parts[0], 10, 64)
		src, _ := strconv.ParseUint(parts[1], 10, 64)
		mlen, _ := strconv.ParseUint(parts[2], 10, 64)
		mappers[len(mappers)-1].data = append(mappers[len(mappers)-1].data, mapData{src, dst, mlen})
	}

	return seeds, multiMapper{mappers}
}

type mapData struct {
	src uint64
	dst uint64
	len uint64
}

type mapper struct {
	data []mapData
}

type multiMapper struct {
	mappers []mapper
}

func (m *mapper) mapFrom(src uint64) uint64 {
	for _, d := range m.data {
		if src >= d.src && src < d.src+d.len {
			return d.dst + (src - d.src)
		}
	}
	return src
}

func (m *multiMapper) mapFrom(src uint64) uint64 {
	for _, mapper := range m.mappers {
		src = mapper.mapFrom(src)
	}
	return src
}
