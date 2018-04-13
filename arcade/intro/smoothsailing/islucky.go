package smoothsailing

import "strconv"

func isLucky(n int) bool {
	str := strconv.Itoa(n)
	firstHalf := 0
	secondHalf := 0
	strLen := len(str)
	for index, num := range str {
		val := int(num)
		if index >= strLen/2 {
			secondHalf += val
		} else {
			firstHalf += val
		}
	}

	return firstHalf == secondHalf
}
