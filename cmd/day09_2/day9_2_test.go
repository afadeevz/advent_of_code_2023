package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay9(t *testing.T) {
	t.Parallel()

	input := bytes.NewBufferString(`0 3 6 9 12 15
1 3 6 10 15 21
10 13 16 21 30 45`)

	assert.Equal(t, 2, run(input))
}
