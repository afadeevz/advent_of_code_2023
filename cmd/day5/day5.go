package main

import (
	"bufio"
	"fmt"
	"io"
	"log/slog"
	"math"
	"os"
	"strconv"
	"strings"
)

type Seeds = []int

type MappingEntry struct {
	src, dst, len int
}

type Mapping struct {
	srcCategory, dstCategory string

	entries []MappingEntry
}

func (m Mapping) Map(value int) int {
	for _, entry := range m.entries {
		if entry.src <= value && value < entry.src+entry.len {
			return value - entry.src + entry.dst
		}
	}

	return value
}

func main() {
	fmt.Println(run(os.Stdin))
}

func run(input io.Reader) int {
	seeds, mappingsSlice := parseInput(input)
	slog.Info("parsed input",
		"seeds", seeds,
		"mappings", mappingsSlice,
	)

	answer := math.MaxInt

	mappings := make(map[string]Mapping, len(mappingsSlice))
	for _, mapping := range mappingsSlice {
		mappings[mapping.srcCategory] = mapping
	}

	for _, seed := range seeds {
		answer = min(answer, processMappings(mappings, seed))
	}

	return answer
}

func processMappings(mappings map[string]Mapping, seed int) int {
	category := "seed"
	value := seed

	for category != "location" {
		mapping := mappings[category]

		slog.Info("Mapped",
			"from", value,
			"to", mapping.Map(value),
		)

		value = mapping.Map(value)
		category = mapping.dstCategory
	}

	return value
}

func parseInput(input io.Reader) (seeds Seeds, mappings []Mapping) {
	scanner := bufio.NewScanner(input)

	scanner.Scan()
	seeds = parseSeedsLine(scanner.Text())

	for {
		mapping := parseMapping(scanner)
		if mapping == nil {
			return
		}

		mappings = append(mappings, *mapping)
	}
}

func parseSeedsLine(line string) []int {
	line, _ = strings.CutPrefix(line, "seeds: ")
	seedsStrs := strings.Split(line, " ")

	seeds := make([]int, 0, len(seedsStrs))

	for _, seedStr := range seedsStrs {
		seed, _ := strconv.ParseInt(seedStr, 10, 64)
		seeds = append(seeds, int(seed))
	}

	return seeds
}

func parseMapping(scanner *bufio.Scanner) *Mapping {
	for {
		if !scanner.Scan() {
			return nil
		}

		if len(scanner.Text()) > 0 {
			break
		}
	}

	var result Mapping

	line := scanner.Text()
	line, _ = strings.CutSuffix(line, " map:")
	parts := strings.Split(line, "-to-")
	result.srcCategory = parts[0]
	result.dstCategory = parts[1]

	for scanner.Scan() && len(scanner.Text()) > 0 {
		line := scanner.Text()
		parts := strings.Split(line, " ")
		dst, _ := strconv.ParseInt(parts[0], 10, 64)
		src, _ := strconv.ParseInt(parts[1], 10, 64)
		len, _ := strconv.ParseInt(parts[2], 10, 64)

		result.entries = append(result.entries, MappingEntry{
			src: int(src),
			dst: int(dst),
			len: int(len),
		})
	}

	return &result
}
