package exploringthewaters

import (
	"reflect"
	"testing"
)

var areSimilarData = []struct {
	a        []int
	b        []int
	expected bool
}{
	{
		[]int{1, 2, 3},
		[]int{1, 2, 3},
		true,
	},
	{
		[]int{1, 2, 3},
		[]int{2, 1, 3},
		true,
	},
	{
		[]int{1, 2, 2},
		[]int{2, 1, 1},
		false,
	},
	{
		[]int{1, 1, 4},
		[]int{1, 2, 3},
		false,
	},
	{
		[]int{1, 2, 3},
		[]int{1, 10, 2},
		false,
	},
	{
		[]int{2, 3, 1},
		[]int{1, 3, 2},
		true,
	},
	{
		[]int{2, 3, 9},
		[]int{10, 3, 2},
		false,
	},
	{
		[]int{4, 6, 3},
		[]int{3, 4, 6},
		false,
	},
	{
		[]int{832, 998, 148, 570, 533, 561, 894, 147, 455, 279},
		[]int{832, 998, 148, 570, 533, 561, 455, 147, 894, 279},
		true,
	},
	{
		[]int{832, 998, 148, 570, 533, 561, 894, 147, 455, 279},
		[]int{832, 570, 148, 998, 533, 561, 455, 147, 894, 279},
		false,
	},
	{
		[]int{1, 2, 3, 4, 5},
		[]int{5, 4, 3, 2, 1},
		false,
	},
	{
		[]int{1, 1, 1, 1, 3},
		[]int{1, 1, 3, 1, 1},
		true,
	},
}

func TestAreSimilar(t *testing.T) {
	for _, tt := range areSimilarData {
		actual := areSimilar(tt.a, tt.b)
		if !reflect.DeepEqual(actual, tt.expected) {
			t.Errorf("\n\tareSimilar(%v, %v): expected %v, actual %v\n\n", tt.a, tt.b, tt.expected, actual)
		}
	}
}
