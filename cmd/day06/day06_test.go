package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay6(t *testing.T) {
	t.Parallel()

	input := bytes.NewBufferString(`Time:      7  15   30
Distance:  9  40  200`)

	assert.Equal(t, 288, run(input))
}
