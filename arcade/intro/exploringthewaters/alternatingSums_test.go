package exploringthewaters

import (
	"reflect"
	"testing"
)

var alternatingSumsData = []struct {
	data     []int
	expected []int
}{
	{
		[]int{50, 60, 60, 45, 70},
		[]int{180, 105},
	},
	{
		[]int{100, 50},
		[]int{100, 50},
	},
	{
		[]int{80},
		[]int{80, 0},
	},
}

func TestAlternatingSums(t *testing.T) {
	for _, tt := range alternatingSumsData {
		actual := alternatingSums(tt.data)
		if !reflect.DeepEqual(actual, tt.expected) {
			t.Errorf("\n\talternatingSums(%v): expected %v, actual %v\n\n", tt.data, tt.expected, actual)
		}
	}
}
