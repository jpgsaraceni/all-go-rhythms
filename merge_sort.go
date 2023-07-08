package main

// MergeSort is an implementation of the Merge Sort algorithm which
// given an array of n length, recursively devides the array into sub-arrays,
// sorts those sub-arrays, then merges the sorted sub-arrays.
// It has a worst-case runtime of O(nlogn).
func MergeSort(sequence []int) []int {
	size := len(sequence)

	// base case
	if size <= 1 {
		return sequence
	}

	leftSide := MergeSort(sequence[:size/2])
	rightSide := MergeSort(sequence[size/2:])

	output := merge(leftSide, rightSide)

	return output
}

func merge(leftSide, rightSide []int) []int {
	output := make([]int, len(leftSide)+len(rightSide))
	var leftIndex, rightIndex int

	for i := 0; i < len(output); i++ {
		if leftIndex >= len(leftSide) { // means left side is over, so copy right side
			output[i] = rightSide[rightIndex]
			rightIndex++

			continue
		}
		if rightIndex >= len(rightSide) { // means right side is over, so copy left side
			output[i] = leftSide[leftIndex]
			leftIndex++

			continue
		}
		if leftSide[leftIndex] < rightSide[rightIndex] {
			output[i] = leftSide[leftIndex]
			leftIndex++

			continue
		}
		output[i] = rightSide[rightIndex]
		rightIndex++
	}

	return output
}
