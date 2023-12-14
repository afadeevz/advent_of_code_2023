package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay7(t *testing.T) {
	t.Parallel()

	input := bytes.NewBufferString(`32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483`)

	assert.Equal(t, 6440, run(input))
}
