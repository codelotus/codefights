package smoothsailing

func allLongestStrings(inputArray []string) []string {
	maxlen := 0
	currentLen := 0
	var result []string
	for _, str := range inputArray {
		currentLen = len(str)
		if currentLen > maxlen {
			maxlen = currentLen
			result = []string{str}
		} else if currentLen == maxlen {
			result = append(result, str)
		}
	}

	return result
}
