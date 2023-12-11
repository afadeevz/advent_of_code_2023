package main

import (
	"bufio"
	"fmt"
	"io"
	"log/slog"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Card rune

func (c Card) Strength() int {
	mapping := map[Card]int{
		'2': 1,
		'3': 2,
		'4': 3,
		'5': 4,
		'6': 5,
		'7': 6,
		'8': 7,
		'9': 8,
		'T': 9,
		'J': 10,
		'Q': 11,
		'K': 12,
		'A': 13,
	}

	result, ok := mapping[c]
	if !ok {
		panic(fmt.Sprintf("unknown card: %c", c))
	}

	return result
}

func (c Card) Compare(other Card) int {
	if c.Strength() > other.Strength() {
		return 1
	}
	if c.Strength() < other.Strength() {
		return -1
	}
	return 0
}

type HandType int

const (
	HighCard HandType = iota
	OnePair
	TwoPair
	ThreeOfAKind
	FullHouse
	FourOfAKind
	FiveOfAKind
)

func (ht HandType) Compare(other HandType) int {
	if ht > other {
		return 1
	}
	if ht < other {
		return -1
	}
	return 0
}

type Hand struct {
	cards []Card
	bid   int
}

func (h Hand) Type() HandType {
	counts := make(map[Card]int)
	for _, card := range h.cards {
		counts[card]++
	}

	countsCounts := make(map[int]int)
	for _, count := range counts {
		countsCounts[count]++
	}

	slog.Info("calc type",
		"cards", string(h.cards),
		"counts", counts,
		"countsCounts", countsCounts,
	)

	if countsCounts[5] == 1 {
		return FiveOfAKind
	}
	if countsCounts[4] == 1 {
		return FourOfAKind
	}
	if countsCounts[3] == 1 {
		if countsCounts[2] == 1 {
			return FullHouse
		} else {
			return ThreeOfAKind
		}
	}
	if countsCounts[2] == 2 {
		return TwoPair
	}
	if countsCounts[2] == 1 {
		return OnePair
	}
	return HighCard
}

func (h Hand) Compare(other Hand) int {
	cmpRes := h.Type().Compare(other.Type())
	if cmpRes != 0 {
		return cmpRes
	}

	for i := range h.cards {
		cmpRes = h.cards[i].Compare(other.cards[i])
		if cmpRes != 0 {
			return cmpRes
		}
	}

	return 0
}

func compareHands(a, b Hand) int {
	return a.Compare(b)
}

func main() {
	slog.Info("got answer", "answer", run(os.Stdin))
}

func run(input io.Reader) int {
	hands := parseInput(input)

	slices.SortFunc[[]Hand](hands, compareHands)

	answer := 0
	for i, hand := range hands {
		slog.Info("got hand",
			"cards", string(hand.cards),
			"bid", hand.bid,
			"rank", i+1,
			"type", hand.Type(),
		)

		answer += (i + 1) * hand.bid
	}

	return answer
}

func parseInput(input io.Reader) (hands []Hand) {
	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		line := scanner.Text()
		hand := parseLine(line)
		hands = append(hands, hand)
	}

	return
}

func parseLine(line string) Hand {
	parts := strings.Split(line, " ")
	cards := []Card(parts[0])
	bid, _ := strconv.ParseInt(parts[1], 10, 64)

	return Hand{
		cards: cards,
		bid:   int(bid),
	}
}
