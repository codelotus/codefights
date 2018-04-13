package smoothsailing

import (
	"reflect"
	"testing"
)

var sortByHeightData = []struct {
	data     []int
	expected []int
}{
	{
		[]int{-1, 150, 190, 170, -1, -1, 160, 180},
		[]int{-1, 150, 160, 170, -1, -1, 180, 190},
	},
	{
		[]int{-1, -1, -1, -1, -1},
		[]int{-1, -1, -1, -1, -1},
	},
	{
		[]int{4, 2, 9, 11, 2, 16},
		[]int{2, 2, 4, 9, 11, 16},
	},
}

func TestSortByHeight(t *testing.T) {
	for _, tt := range sortByHeightData {
		actual := sortByHeight(tt.data)
		if !reflect.DeepEqual(actual, tt.expected) {
			t.Errorf("\n\tsortByHeight(%v): expected %v, actual %v\n\n", tt.data, tt.expected, actual)
		}
	}
}
