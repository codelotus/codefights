package edgeoftheocean

import "testing"

var almostIncreaseData = []struct {
	data     []int
	expected bool
}{
	{[]int{1, 3, 2, 1}, false},
	{[]int{1, 3, 2}, true},
	{[]int{123, -17, -5, 1, 2, 3, 12, 43, 45}, true},
}

func TestAlmostIncreasingSequence(t *testing.T) {
	for _, tt := range almostIncreaseData {
		actual := almostIncreasingSequence(tt.data)
		if actual != tt.expected {
			t.Errorf("almostIncreasingSequence(%v): expected %t, actual %t", tt.data, tt.expected, actual)
		}
	}
}
