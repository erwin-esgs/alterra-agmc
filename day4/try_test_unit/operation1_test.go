package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOperation1(t *testing.T) {
	result := Add(5, 2)
	assert.Equal(t, 7.0, result)
}

func TestOperation2(t *testing.T) {
	result := Subtract(5, 2)
	assert.Equal(t, 3.0, result)
}

func TestOperation3(t *testing.T) {
	result := Times(5, 2)
	assert.Equal(t, 10.0, result)
}

func TestOperation4(t *testing.T) {
	result := Divide(5, 2)
	assert.Equal(t, 2.5, result)
}
