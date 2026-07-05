package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNegativeSize(t *testing.T) {
	var numbers []int

	require.NotPanics(t, func() { numbers = generateRandomElements(-10) })
	assert.Len(t, numbers, 0)
}

func TestZeroElements(t *testing.T) {
	var numbers []int
	var max_v, maxThreaded int

	require.NotPanics(t, func() { numbers = generateRandomElements(0) })
	assert.Len(t, numbers, 0)

	if assert.NotPanics(t, func() { max_v = maximum(numbers) }) {
		assert.Equal(t, 0, max_v)
	}

	if assert.NotPanics(t, func() { max_v = maxChunks(numbers) }) {
		assert.Equal(t, 0, max_v)
	}

	assert.Equal(t, max_v, maxThreaded)
}

func TestOneElement(t *testing.T) {
	var numbers []int
	var max_v int

	require.NotPanics(t, func() { numbers = generateRandomElements(1) })
	assert.Len(t, numbers, 1)

	if assert.NotPanics(t, func() { max_v = maximum(numbers) }) {
		assert.Equal(t, numbers[0], max_v)
	}

	if assert.NotPanics(t, func() { max_v = maxChunks(numbers) }) {
		assert.Equal(t, numbers[0], max_v)
	}
}

func TestSubChunks(t *testing.T) {
	start, end := 2, 15

	for size := start; size <= end; size++ {
		var numbers []int
		var max_v, maxThreaded int

		require.NotPanics(t, func() { numbers = generateRandomElements(size) })
		assert.Len(t, numbers, size)

		assert.NotPanics(t, func() { max_v = maximum(numbers) })

		assert.NotPanics(t, func() { maxThreaded = maxChunks(numbers) })

		assert.Equal(t, max_v, maxThreaded)
	}
}

func TestResult(t *testing.T) {
	variants := []struct {
		in  []int
		out int
	}{
		// 1. Basic
		{in: []int{1, 2, 3, 4}, out: 4},      // Basic Positive Increasing
		{in: []int{-1, -2, -3, -4}, out: -1}, // All Negative (Max is closest to zero)

		// 2. Max at Start/End
		{in: []int{5, 1, 2}, out: 5}, // Max at the beginning
		{in: []int{1, 2, 8}, out: 8}, // Max at the end

		// 3. Mixed Signs
		{in: []int{-10, 5, -3, 8}, out: 8},  // Positive max among negative/positive mix
		{in: []int{10, -5, 2, -1}, out: 10}, // Max is at the start

		// 4. Zero handling
		{in: []int{0, 0, 0, 0}, out: 0}, // All Zeros
		{in: []int{-5, 0, -2}, out: 0},  // Zero is the maximum

		// 5. Repeated Maximums
		{in: []int{7, 2, 7, 1}, out: 7},  // Max repeats in the middle
		{in: []int{-5, -5, -5}, out: -5}, // All same negative value
	}

	for _, variant := range variants {
		var max_v int

		t.Run("TestMaximumFunction", func(t *testing.T) {
			assert.NotPanics(t, func() { max_v = maximum(variant.in) })
			assert.Equal(t, variant.out, max_v, "Expected %d but got %d for input %v", variant.out, max_v, variant.in)
		})

		t.Run("TestMaxChunksFunction", func(t *testing.T) {
			assert.NotPanics(t, func() { max_v = maxChunks(variant.in) })
			assert.Equal(t, variant.out, max_v, "Expected %d but got %d for input %v", variant.out, max_v, variant.in)
		})
	}
}

func TestLargeSet(t *testing.T) {
	const size = 100_000_000
	var numbers []int
	var max_v, maxThreaded int

	require.NotPanics(t, func() { numbers = generateRandomElements(size) })
	assert.Len(t, numbers, size)

	assert.NotPanics(t, func() { max_v = maximum(numbers) })

	assert.NotPanics(t, func() { maxThreaded = maxChunks(numbers) })

	assert.Equal(t, max_v, maxThreaded)
}
