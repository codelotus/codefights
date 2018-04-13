package exploringthewaters

import (
	"reflect"
	"testing"
)

var arrayChangeData = []struct {
	data     []int
	expected int
}{
	{
		[]int{1, 1, 1},
		3,
	},
	{
		[]int{-1000, 0, -2, 0},
		5,
	},

	{
		[]int{2, 1, 10, 1},
		12,
	},
	{
		[]int{2, 3, 3, 5, 5, 5, 4, 12, 12, 10, 15},
		13,
	},
}

func TestArrayChange(t *testing.T) {
	for _, tt := range arrayChangeData {
		actual := arrayChange(tt.data)
		if !reflect.DeepEqual(actual, tt.expected) {
			t.Errorf("\n\tarrayChange(%v): expected %v, actual %v\n\n", tt.data, tt.expected, actual)
		}
	}
}
