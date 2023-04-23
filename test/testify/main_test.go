package main

import (
	assert2 "github.com/stretchr/testify/assert"
	"testing"
)

func TestSquare1(t *testing.T) {
	assert := assert2.New(t)
	assert.Equal(81, square(9), "square(9) should be 81")
}
