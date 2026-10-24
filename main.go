package main

import (
	"fmt"
	"math"
	"math/rand"
	"sync"
	"time"
)

const (
	SIZE   = 100_000_000
	CHUNKS = 8
)

// generateRandomElements generates random elements.
func generateRandomElements(size int) []int {
	if size < 0 {
		size = 0
	}
	slice := make([]int, size)

	for i := 0; i < size; i++ {
		slice[i] = rand.Int()
	}

	return slice
}

func maxUnsafe(data []int) int {
	max_v := math.MinInt

	for _, v := range data {
		max_v = max(max_v, v)
	}

	return max_v
}

// maximum returns the maximum number of elements.
func maximum(data []int) int {
	if len(data) == 0 {
		return 0
	}

	return maxUnsafe(data)
}

// maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data []int) int {
	data_len := len(data)
	if data_len == 0 {
		return 0
	}

	chunk_len := int(math.Ceil(float64(data_len) / float64(CHUNKS)))

	chunk_max := [CHUNKS]int{}
	for i := range chunk_max {
		chunk_max[i] = math.MinInt
	}

	var wg sync.WaitGroup

	for i, start := 0, 0; start < data_len; i++ {
		wg.Add(1)
		end := min(start+chunk_len, data_len)

		go func(start int, end int) {
			chunk_max[i] = maxUnsafe(data[start:end])
			wg.Done()
		}(start, end)

		start = end
	}

	wg.Wait()

	return maxUnsafe(chunk_max[:])
}

func main() {
	fmt.Printf("Генерируем %d целых чисел\n", SIZE)
	numbers := generateRandomElements(SIZE)

	fmt.Println("Ищем максимальное значение в один поток")
	start := time.Now()
	max := maximum(numbers)
	elapsed := time.Since(start).Microseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d μs\n", max, elapsed)

	fmt.Printf("Ищем максимальное значение в %d потоков\n", CHUNKS)
	start = time.Now()
	max = maxChunks(numbers)
	elapsed = time.Since(start).Microseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d μs\n", max, elapsed)
}
