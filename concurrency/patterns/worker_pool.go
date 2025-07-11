package main

import (
	"fmt"
	"math"
	"sync"
)

// worker
func getMax(jobs chan []int, results chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		res := math.MinInt
		for _, v := range job {
			res = max(res, v)
		}
		results <- res
	}
}

func main() {
	arr := []int{1, 3, 7, 0, 5, 6, 2, 8, 4, 9, 11, 1111}
	jobSize := 3
	jobs := make(chan []int)
	results := make(chan int)
	var wg sync.WaitGroup

	// Start worker pool
	workerCount := 10
	for range workerCount {
		wg.Add(1)
		go getMax(jobs, results, &wg)
	}

	// send jobs
	go func() {
		for i := 0; i < len(arr); i += jobSize {
			job := arr[i:min(i+jobSize, len(arr))]
			jobs <- job
		}
		close(jobs)
	}()

	// Wait for all workers, then close results
	go func() {
		wg.Wait()
		close(results)
	}()

	// Collect results
	maxVal := math.MinInt
	for num := range results {
		maxVal = max(maxVal, num)
	}

	fmt.Println(maxVal)
}
