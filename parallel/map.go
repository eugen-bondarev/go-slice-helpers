package parallel

import (
	"runtime"
	"sync"
)

func MapVerbose[TIn any, TOut any](input []TIn, mutate func(t TIn) TOut, workers int) []TOut {
	output := make([]TOut, len(input))

	batches := createBatches(len(input), workers)

	start := 0

	wg := sync.WaitGroup{}
	wg.Add(len(batches))

	for i := 0; i < len(batches); i++ {
		end := start + batches[i]

		go func(start, end int) {
			for j := start; j < end; j++ {
				output[j] = mutate(input[j])
			}
			wg.Done()
		}(start, end)

		start += batches[i]
	}

	wg.Wait()

	return output
}

func Map[TIn any, TOut any](input []TIn, mutate func(t TIn) TOut) []TOut {
	return MapVerbose(input, mutate, runtime.NumCPU())
}
