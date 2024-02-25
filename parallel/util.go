package parallel

func createBatches(length, workers int) []int {
	divided, remainder := length/workers, length%workers
	batches := make([]int, workers)

	for i := 0; i < len(batches); i++ {
		batches[i] = divided + min(max(remainder, 0), 1)
		remainder--
	}

	return batches
}
