package parallel_test

import (
	"go-slice-helpers/parallel"
	"reflect"
	"strconv"
	"testing"
)

func normalMap[T any, U any](input []T, mutate func(input T) U) []U {
	output := make([]U, len(input))

	for i := 0; i < len(output); i++ {
		output[i] = mutate(input[i])
	}

	return output
}

func TestMap(t *testing.T) {
	type testCase struct {
		input    []int
		callback func(v int) string
	}

	testCases := []testCase{
		{
			input: []int{},
			callback: func(v int) string {
				return strconv.Itoa((v + 1) * (v + 1))
			},
		},
		{
			input: []int{1, 2, 3},
			callback: func(v int) string {
				return strconv.Itoa(v * v)
			},
		},
		{
			input: []int{1, 2, 3},
			callback: func(v int) string {
				return strconv.Itoa((v + 1) * (v + 1))
			},
		},
	}

	for _, testCase := range testCases {
		should := normalMap(testCase.input, testCase.callback)
		is := parallel.Map(testCase.input, testCase.callback)

		if !reflect.DeepEqual(should, is) {
			t.Fatalf("should: %v, is: %v\n", should, is)
		}
	}
}
