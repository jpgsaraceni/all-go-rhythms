package main

// CountInversions_BruteForce is an algorithm to detect how mnay inversions i.e. elements
// that break an order in a list of integers, with O(n²) runtime.
func CountInversions_BruteForce(sequence []int) int {
	counter := 0

	for i, number := range sequence {
		for j := i + 1; j < len(sequence); j++ {
			if number > sequence[j] {
				counter++
			}
		}
	}

	return counter
}

// CountInversions_DevideAndConquer is like CountInversions_BruteForce, but with O(nlogn) runtime.
func CountInversions_DevideAndConquer(sequence []int) int {
	var inversionCounter int
	sortAndCountInversions(sequence, &inversionCounter)

	return inversionCounter
}

func sortAndCountInversions(sequence []int, inversionCounter *int) []int {
	size := len(sequence)
	if size <= 1 {
		return sequence
	}

	leftSide := sortAndCountInversions(sequence[:size/2], inversionCounter)
	rightSide := sortAndCountInversions(sequence[size/2:], inversionCounter)
	sorted, splitInversions := countSplitInversions(leftSide, rightSide)
	*inversionCounter += splitInversions

	return sorted
}

func countSplitInversions(leftSide, rightSide []int) ([]int, int) {
	output := make([]int, len(leftSide)+len(rightSide))
	var leftIndex, rightIndex, inversions int

	for i := 0; i < len(output); i++ {
		if leftIndex >= len(leftSide) {
			output[i] = rightSide[rightIndex]
			rightIndex++

			continue
		}
		if rightIndex >= len(rightSide) {
			output[i] = leftSide[leftIndex]
			leftIndex++

			continue
		}
		if leftSide[leftIndex] < rightSide[rightIndex] {
			output[i] = leftSide[leftIndex]
			leftIndex++

			continue
		}
		// means number to the right is greater, so it has as many inversions as
		// the amount of remaining elements to the left.
		output[i] = rightSide[rightIndex]
		rightIndex++
		inversions += len(leftSide) - leftIndex
	}

	return output, inversions
}
