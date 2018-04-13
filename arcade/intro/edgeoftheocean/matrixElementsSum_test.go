package edgeoftheocean

import "testing"

var matrixSumData = []struct {
	data     [][]int
	expected int
}{
	{[][]int{
		{0, 1, 1, 2},
		{0, 5, 0, 0},
		{2, 0, 3, 3}}, 9},
	{[][]int{
		{1, 1, 1, 0},
		{0, 5, 0, 1},
		{2, 1, 3, 10}}, 9},
	{[][]int{
		{1, 1, 1},
		{2, 2, 2},
		{3, 3, 3}}, 18},
	{[][]int{{0}}, 0},
}

func TestMatrixElementsSum(t *testing.T) {
	for _, tt := range matrixSumData {
		actual := matrixElementsSum(tt.data)
		if actual != tt.expected {
			t.Errorf("matrixElementsSum(%v): expected %d, actual %d", tt.data, tt.expected, actual)
		}
	}
}
