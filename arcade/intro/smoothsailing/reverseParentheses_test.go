package smoothsailing

import (
	"reflect"
	"testing"
)

var reverseParenthesesData = []struct {
	data     string
	expected string
}{
	{
		"a(bc)de",
		"acbde",
	},
	{
		"a(bcdefghijkl(mno)p)q",
		"apmnolkjihgfedcbq",
	},
	{
		"co(de(fight)s)",
		"cosfighted",
	},
	{
		"Code(Cha(lle)nge)",
		"CodeegnlleahC",
	},
	{
		"Where are the parentheses?",
		"Where are the parentheses?",
	},
	{
		"abc(cba)ab(bac)c",
		"abcabcabcabc",
	},

	{
		"The ((quick (brown) (fox) jumps over the lazy) dog)",
		"The god quick nworb xof jumps over the lazy",
	},
}

func TestReverseParentheses(t *testing.T) {
	for _, tt := range reverseParenthesesData {
		actual := reverseParentheses(tt.data)
		if !reflect.DeepEqual(actual, tt.expected) {
			t.Errorf("\n\treverseParentheses(%v): expected %v, actual %v\n\n", tt.data, tt.expected, actual)
		}
	}
}
