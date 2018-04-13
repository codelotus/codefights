package smoothsailing

import (
	"testing"
)

var isLuckyData = []struct {
	data     int
	expected bool
}{
	{1230, true},
	{239017, false},
	{134008, true},
	{10, false},
	{11, true},
	{1010, true},
	{261534, false},
	{100000, false},
	{999999, true},
	{123321, true},
}

func TestIsLucky(t *testing.T) {
	for _, tt := range isLuckyData {
		actual := isLucky(tt.data)
		if actual != tt.expected {
			t.Errorf("\n\tisLucky(%v): expected %v, actual %v\n\n", tt.data, tt.expected, actual)
		}
	}
}
