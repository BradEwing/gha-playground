package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	tests := []struct {
		name     string
		inputA   int
		inputB   int
		expected int
	}{
		{
			name:     "zero",
			inputA:   0,
			inputB:   0,
			expected: 0,
		},
		{
			name:     "negative",
			inputA:   -1,
			inputB:   -2,
			expected: -3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, add(tt.inputA, tt.inputB), tt.expected)
		})
	}
}
