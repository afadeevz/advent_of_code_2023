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

type SeedRange struct {
	start, len int
}

type SeedRanges = []SeedRange

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
	seedRanges, mappingsSlice := parseInput(input)
	slog.Info("parsed input",
		"seeds", seedRanges,
		"mappings", mappingsSlice,
	)

	answer := math.MaxInt

	mappings := make(map[string]Mapping, len(mappingsSlice))
	for _, mapping := range mappingsSlice {
		mappings[mapping.srcCategory] = mapping
	}

	for _, seedRange := range seedRanges {
		for seed := seedRange.start; seed < seedRange.start+seedRange.len; seed++ {
			answer = min(answer, processMappings(mappings, seed))
		}
		slog.Info("range passed", "start", seedRange.start, "len", seedRange.len)
	}

	return answer
}

func processMappings(mappings map[string]Mapping, seed int) int {
	category := "seed"
	value := seed

	for category != "location" {
		mapping := mappings[category]

		// slog.Info("Mapped",
		// 	"from", value,
		// 	"to", mapping.Map(value),
		// )

		value = mapping.Map(value)
		category = mapping.dstCategory
	}

	return value
}

func parseInput(input io.Reader) (seedRanges SeedRanges, mappings []Mapping) {
	scanner := bufio.NewScanner(input)

	scanner.Scan()
	seedRanges = parseSeedRangesLine(scanner.Text())

	for {
		mapping := parseMapping(scanner)
		if mapping == nil {
			return
		}

		mappings = append(mappings, *mapping)
	}
}

func parseSeedRangesLine(line string) SeedRanges {
	line, _ = strings.CutPrefix(line, "seeds: ")
	parts := strings.Split(line, " ")

	seedRanges := make(SeedRanges, 0, len(parts)/2)

	for i := 0; i < len(parts); i += 2 {
		start, _ := strconv.ParseInt(parts[i], 10, 64)
		len, _ := strconv.ParseInt(parts[i+1], 10, 64)
		seedRanges = append(seedRanges, SeedRange{
			start: int(start),
			len:   int(len),
		})
	}

	return seedRanges
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
