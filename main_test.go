package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAdd(t *testing.T) {
	assert.Equal(t, 4, add(2, 2), "2 + 2 = 4")
}

func TestSubtract(t *testing.T) {
	assert.Equal(t, 0, subtract(2, 2), "2 - 2 = 0")
}
