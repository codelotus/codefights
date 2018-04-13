package smoothsailing

import (
	"reflect"
	"testing"
)

var commonCharacterData = []struct {
	s1       string
	s2       string
	expected int
}{
	{"aabcc", "adcaa", 3},
	{"zzzz", "zzzzzzz", 4},
	{"abca", "xyzbac", 3},
	{"a", "b", 0},
	{"a", "aaa", 1},
}

func TestCommonCharacterCount(t *testing.T) {
	for _, tt := range commonCharacterData {
		actual := commonCharacterCount(tt.s1, tt.s2)
		if !reflect.DeepEqual(actual, tt.expected) {
			t.Errorf("\n\tcommonCharacterCount(%v, %v): expected %v, actual %v\n\n", tt.s1, tt.s2, tt.expected, actual)
		}
	}
}
