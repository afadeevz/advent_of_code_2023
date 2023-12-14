package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay12(t *testing.T) {
	t.Parallel()

	input := bytes.NewBufferString(`???.### 1,1,3
.??..??...?##. 1,1,3
?#?#?#?#?#?#?#? 1,3,1,6
????.#...#... 4,1,1
????.######..#####. 1,6,5
?###???????? 3,2,1`)

	assert.Equal(t, 21, run(input, foldRatePart1))
}

//go:embed input12.txt
var inputFileData string

func TestDay12InputFile(t *testing.T) {
	t.Parallel()

	input := bytes.NewBufferString(inputFileData)

	assert.Equal(t, 0x1E49, run(input, foldRatePart1))
}

func TestDay12Lines(t *testing.T) {
	t.Parallel()

	type TestCase struct {
		input  string
		output int
	}

	testCases := []TestCase{
		{"???.### 1,1,3", 1},
		{".??..??...?##. 1,1,3", 4},
		{"?#?#?#?#?#?#?#? 1,3,1,6", 1},
		{"????.#...#... 4,1,1", 1},
		{"????.######..#####. 1,6,5", 4},
		{"?###???????? 3,2,1", 10},
		{"#??????? 2,1", 5},
	}

	for i, testCase := range testCases {
		testCase := testCase
		t.Run(fmt.Sprintf("Case%d", i+1), func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, testCase.output, processLine(testCase.input, foldRatePart1))
		})
	}
}

// func TestDay12Part2(t *testing.T) {
// 	t.Parallel()

// 	input := bytes.NewBufferString(`???.### 1,1,3
// .??..??...?##. 1,1,3
// ?#?#?#?#?#?#?#? 1,3,1,6
// ????.#...#... 4,1,1
// ????.######..#####. 1,6,5
// ?###???????? 3,2,1`)

// 	assert.Equal(t, 525152, run(input, foldRatePart2))
// }

// func TestDay12LinesPart2(t *testing.T) {
// 	t.Parallel()

// 	type TestCase struct {
// 		input  string
// 		output int
// 	}

// 	testCases := []TestCase{
// 		{"???.### 1,1,3", 1},
// 		{".??..??...?##. 1,1,3", 16384},
// 		{"?#?#?#?#?#?#?#? 1,3,1,6", 1},
// 		{"????.#...#... 4,1,1", 16},
// 		{"????.######..#####. 1,6,5", 2500},
// 		{"?###???????? 3,2,1", 506250},
// 	}

// 	for i, testCase := range testCases {
// 		testCase := testCase
// 		t.Run(fmt.Sprintf("Case%d", i+1), func(t *testing.T) {
// 			t.Parallel()
// 			assert.Equal(t, testCase.output, processLine(testCase.input, foldRatePart2))
// 		})
// 	}
// }
