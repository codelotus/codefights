package exploringthewaters

import (
	"reflect"
	"testing"
)

var addBorderData = []struct {
	data     []string
	expected []string
}{
	{
		[]string{"abc", "ded"},
		[]string{"*****", "*abc*", "*ded*", "*****"},
	},
	{
		[]string{"a"},
		[]string{"***", "*a*", "***"},
	},
	{
		[]string{"aa", "**", "zz"},
		[]string{"****", "*aa*", "****", "*zz*", "****"},
	},
	{
		[]string{"abcde", "fghij", "klmno", "pqrst", "uvwxy"},
		[]string{"*******", "*abcde*", "*fghij*", "*klmno*", "*pqrst*", "*uvwxy*", "*******"},
	},
	{
		[]string{"wzy**"},
		[]string{"*******", "*wzy***", "*******"},
	},
}

func TestAddBorder(t *testing.T) {
	for _, tt := range addBorderData {
		actual := addBorder(tt.data)
		if !reflect.DeepEqual(actual, tt.expected) {
			t.Errorf("\n\taddBorder(%v): expected %v, actual %v\n\n", tt.data, tt.expected, actual)
		}
	}
}
