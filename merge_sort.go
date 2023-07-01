package main

import "fmt"

// MergeSort is an implementation of the Merge Sort algorithm which
// given an array of n length, recursively devides the array into sub-arrays,
// sorts those sub-arrays, then merges the sorted sub-arrays.
// It has a worst-case runtime of O(nlogn).
func MergeSort(arr []int) []int {
	length := len(arr)

	// base case
	if length <= 1 {
		return arr
	}

	leftArr := MergeSort(arr[0 : length/2])
	rightArr := MergeSort(arr[length/2 : length])

	output := merge(leftArr, rightArr)

	return output
}

func merge(leftArr, rightArr []int) []int {
	output := make([]int, len(leftArr)+len(rightArr))
	var leftIndex, rightIndex int

	for i := 0; i < len(output); i++ {
		if leftIndex >= len(leftArr) {
			output[i] = rightArr[rightIndex]
			rightIndex++

			continue
		}
		if rightIndex >= len(rightArr) {
			output[i] = leftArr[leftIndex]
			leftIndex++

			continue
		}
		if leftArr[leftIndex] < rightArr[rightIndex] {
			output[i] = leftArr[leftIndex]
			leftIndex++

			continue
		}
		output[i] = rightArr[rightIndex]
		rightIndex++
	}

	return output
}

func main() {
	fmt.Println(MergeSort([]int{1, 6, 3, 4, 10, 0, 13}))
}
