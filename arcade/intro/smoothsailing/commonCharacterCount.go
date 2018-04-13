package smoothsailing

import "strings"

func commonCharacterCount(s1 string, s2 string) int {
	var result int
	for _, char := range strings.Split(s1, "") {
		if strings.Contains(s2, char) {
			// remove the matching char from s2
			s2 = strings.Replace(s2, char, "", 1)
			result++
		}
	}
	return result
}
