package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	input  string
	answer int
}

func TestDay8(t *testing.T) {
	testCases := []TestCase{
		{
			input: `RL

AAA = (BBB, CCC)
BBB = (DDD, EEE)
CCC = (ZZZ, GGG)
DDD = (DDD, DDD)
EEE = (EEE, EEE)
GGG = (GGG, GGG)
ZZZ = (ZZZ, ZZZ)`,
			answer: 2,
		}, {
			input: `LLR

AAA = (BBB, BBB)
BBB = (AAA, ZZZ)
ZZZ = (ZZZ, ZZZ)`,
			answer: 6,
		},
	}

	for i, testCase := range testCases {
		input := bytes.NewBufferString(testCase.input)

		assert.Equalf(t, testCase.answer, run(input), "test case #%d failed", i)
	}
}
