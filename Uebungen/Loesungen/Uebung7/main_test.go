package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDivide(t *testing.T) {
	result, err := divide(4, 2)

	assert.NoError(t, err)
	assert.Equal(t, float32(2), result)
}

func TestDivideWithError(t *testing.T) {
	_, err := divide(4, 0)
	assert.Error(t, err)
}
