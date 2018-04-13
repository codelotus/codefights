package smoothsailing

import (
	"reflect"
	"testing"
)

var longestStringData = []struct {
	data     []string
	expected []string
}{
	{[]string{"aba", "aa", "ad", "vcd", "aba"}, []string{"aba", "vcd", "aba"}},
	{[]string{"aa"}, []string{"aa"}},
	{[]string{"abc", "eeee", "abcd", "dcd"}, []string{"eeee", "abcd"}},
	{[]string{"a", "abc", "cbd", "zzzzzz", "a", "abcdef", "asasa", "aaaaaa"}, []string{"zzzzzz", "abcdef", "aaaaaa"}},
}

func TestAllLongestStrings(t *testing.T) {
	for _, tt := range longestStringData {
		actual := allLongestStrings(tt.data)
		if !reflect.DeepEqual(actual, tt.expected) {
			t.Errorf("\n\tallLongestStrings(%v): expected %v, actual %v\n\n", tt.data, tt.expected, actual)
		}
	}
}
