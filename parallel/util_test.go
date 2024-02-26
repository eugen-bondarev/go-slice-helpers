package parallel

import (
	"math/rand"
	"testing"
	"time"
)

// [min, max)
func getInt(r *rand.Rand, min, max int) int {
	return rand.Intn(max-min) + min
}

func sum(values []int) int {
	res := 0
	for _, v := range values {
		res += v
	}
	return res
}

func Test_createBatches(t *testing.T) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < 100; i++ {
		sliceLength := getInt(r, 0, 700000)
		workers := getInt(r, 1, 128)

		batches := createBatches(sliceLength, workers)

		should := sliceLength
		is := sum(batches)

		if should != is {
			t.Fatalf("should: %v, is: %v", should, is)
		}
	}
}
