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
	t.Parallel()

	testCases := []TestCase{
		{
			input: `LR

11A = (11B, XXX)
11B = (XXX, 11Z)
11Z = (11B, XXX)
22A = (22B, XXX)
22B = (22C, 22C)
22C = (22Z, 22Z)
22Z = (22B, 22B)
XXX = (XXX, XXX)`,
			answer: 6,
		},
	}

	for i, testCase := range testCases {
		input := bytes.NewBufferString(testCase.input)

		assert.Equalf(t, testCase.answer, run(input), "test case #%d failed", i)
	}
}
