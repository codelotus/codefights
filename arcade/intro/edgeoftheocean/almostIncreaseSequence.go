package edgeoftheocean

import "math"

func almostIncreasingSequence(sequence []int) bool {
	return validateSequence(findElementToRemove(sequence))
}

func findElementToRemove(sequence []int) []int {
	lastVal := math.MinInt64
	var num int

	// Handle first element:
	if sequence[0] >= sequence[1] {
		return append(sequence[:0], sequence[1:]...)
	}
	for i := 0; i < len(sequence); i++ {
		num = sequence[i]
		if num < lastVal &&
			i+1 < len(sequence) && // Index Out of Range check
			sequence[i+1] > sequence[i-2] && // see if the next number is greater than two numbers back
			num > sequence[i-2] { // if current number is greater than two numbers back
			sequence = append(sequence[:i-1], sequence[i:]...)
			return sequence
		} else if num < lastVal {
			sequence = append(sequence[:i], sequence[i+1:]...)
		}
		lastVal = num
	}
	return sequence
}

func validateSequence(sequence []int) bool {
	lastVal := -1000
	for _, num := range sequence {
		if num <= lastVal {
			return false
		}
		lastVal = num
	}
	return true
}
