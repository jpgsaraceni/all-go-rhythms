package main

import (
	"testing"
)

func TestCountInversions(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name   string
		input  []int
		expect int
	}{
		{
			name:   "should return inversions for even amount of numbers",
			input:  []int{1, 3, 5, 2, 4, 6},
			expect: 3,
		},
		{
			name:   "should return inversions for odd amount of numbers",
			input:  []int{7, 3, 5, 2, 4, 6, 1},
			expect: 14,
		},
		{
			name:   "should return zero inversion for sorted numbers",
			input:  []int{1, 2, 3, 6, 7, 10, 12},
			expect: 0,
		},
	}

	for _, tt := range testCases {
		test := tt
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			got_bruteForce := CountInversions_BruteForce(test.input)
			assertEqual(t,
				test.name+" by brute force",
				test.expect,
				got_bruteForce)

			got_mergeSort := CountInversions_DevideAndConquer(test.input)
			assertEqual(t,
				test.name+" by devide and conquer",
				test.expect,
				got_mergeSort)
		})
	}
}
