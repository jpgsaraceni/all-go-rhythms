package main

import (
	"testing"
)

func TestMergeSort(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name   string
		input  []int
		expect []int
	}{
		{
			name:   "should sort array with even amount of elements",
			input:  []int{1, 3, 10, 2, 4, 5, 9, 6, 8, 11},
			expect: []int{1, 2, 3, 4, 5, 6, 8, 9, 10, 11},
		},
		{
			name:   "should sort array with odd amount of elements",
			input:  []int{1, 3, 10, 2, 4, 5, 9, 6, 11},
			expect: []int{1, 2, 3, 4, 5, 6, 9, 10, 11},
		},
		{
			name:   "should return same array when it is already sorted",
			input:  []int{1, 2, 3, 4, 5, 6, 8, 9, 10, 11},
			expect: []int{1, 2, 3, 4, 5, 6, 8, 9, 10, 11},
		},
	}

	for _, tc := range testCases {
		test := tc
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			got := MergeSort(test.input)
			assertEqual(t, test.name, test.expect, got)
		})
	}
}
