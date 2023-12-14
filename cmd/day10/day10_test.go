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

func TestDay10(t *testing.T) {
	t.Parallel()

	testCases := []TestCase{
		{
			input: `.....
.S-7.
.|.|.
.L-J.
.....`,
			answer: 4,
		}, {
			input: `..F7.
.FJ|.
SJ.L7
|F--J
LJ...`,
			answer: 8,
		},
	}

	for i, testCase := range testCases {
		input := bytes.NewBufferString(testCase.input)

		assert.Equalf(t, testCase.answer, run(input), "test case #%d failed", i)
	}
}
