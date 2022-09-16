package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSubtestOperation(t *testing.T) {
	t.Run("add", func(t *testing.T) {
		result := Add(5, 2)
		assert.Equal(t, 7.0, result)
	})

	t.Run("subtract", func(t *testing.T) {
		result := Subtract(5, 2)
		assert.Equal(t, 3.0, result)
	})

	t.Run("times", func(t *testing.T) {
		result := Times(5, 2)
		assert.Equal(t, 10.0, result)
	})

	t.Run("divided", func(t *testing.T) {
		result := Divide(5, 2)
		assert.Equal(t, 2.5, result)
	})
}
