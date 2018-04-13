package islandofknowledge

import "testing"

var areEquallyStrongData = []struct {
	yourLeft     int
	yourRight    int
	friendsLeft  int
	friendsRight int
	expected     bool
}{
	{10, 15, 15, 10, true},
	{15, 10, 15, 10, true},
	{15, 10, 15, 9, false},
	{10, 5, 5, 10, true},
	{10, 15, 5, 20, false},
	{10, 15, 5, 20, false},
	{10, 20, 10, 20, true},
	{5, 20, 20, 5, true},
	{20, 15, 5, 20, false},
	{5, 10, 5, 10, true},
	{1, 10, 10, 0, false},
	{5, 5, 10, 10, false},
	{10, 5, 10, 6, false},
	{1, 1, 1, 1, true},
}

func TestAreEquallyStrong(t *testing.T) {
	for _, tt := range areEquallyStrongData {
		actual := areEquallyStrong(tt.yourLeft, tt.yourRight, tt.friendsLeft, tt.friendsRight)
		if actual != tt.expected {
			t.Errorf("\n\tareEquallyStrong(%v, %v, %v, %v): expected %v, actual %v\n\n", tt.yourLeft, tt.yourRight, tt.friendsLeft, tt.friendsRight, tt.expected, actual)
		}
	}
}
