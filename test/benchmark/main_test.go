package main

import (
	assert2 "github.com/stretchr/testify/assert"
	"testing"
)

func TestFibonacci1(t *testing.T) {
	assert := assert2.New(t)

	assert.Equal(0, fibonacci1(-1), "fibonacci1(-1) should be 0")
	assert.Equal(0, fibonacci1(0), "fibonacci1(0) should be 0")
	assert.Equal(1, fibonacci1(1), "fibonacci1(1) should be 1")
	assert.Equal(2, fibonacci1(3), "fibonacci1(3) should be 2")
	assert.Equal(233, fibonacci1(13), "fibonacci1(13) should be 233")
}

func TestFibonacci2(t *testing.T) {
	assert := assert2.New(t)

	assert.Equal(0, fibonacci2(-1), "fibonacci2(-1) should be 0")
	assert.Equal(0, fibonacci2(0), "fibonacci2(0) should be 0")
	assert.Equal(1, fibonacci2(1), "fibonacci2(1) should be 1")
	assert.Equal(2, fibonacci2(3), "fibonacci2(3) should be 2")
	assert.Equal(233, fibonacci2(13), "fibonacci2(13) should be 233")
}
