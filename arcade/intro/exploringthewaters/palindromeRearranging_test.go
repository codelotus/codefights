package exploringthewaters

import "testing"

var palindromeRearrangingData = []struct {
	data     string
	expected bool
}{
	{
		"aabb",
		true,
	},
	{
		"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaabc",
		false,
	},
	{
		"abbcabb",
		true,
	},
	{
		"zyyzzzzz",
		true,
	},
	{
		"z",
		true,
	},
	{
		"zaa",
		true,
	},
	{
		"abca",
		false,
	},
}

func TestPalindromeRearranging(t *testing.T) {
	for _, tt := range palindromeRearrangingData {
		actual := palindromeRearranging(tt.data)
		if actual != tt.expected {
			t.Errorf("\n\tpalindromeRearranging(%v): expected %v, actual %v\n\n", tt.data, tt.expected, actual)
		}
	}
}
