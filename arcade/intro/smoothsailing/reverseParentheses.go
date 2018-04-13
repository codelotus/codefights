package smoothsailing

func reverseParentheses(s string) string {
	var openParens []int
	var closeParens []int
	input := []rune(s)

	// find parens
	openParens, closeParens = findParens(s)

	if len(openParens) == 0 && len(closeParens) == 0 { // there are no more parens
		return s
	}
	currentOpenParen, currentCloseParen := findCurrentParens(openParens, closeParens)
	// find inner parens and reverse
	substr := input[currentOpenParen+1 : currentCloseParen]
	reversedSubStr := reverse(substr)
	prepen := append(input[:currentOpenParen], reversedSubStr...)
	postpen := append(prepen, input[currentCloseParen+1:]...)
	return reverseParentheses(string(postpen))
}

func findParens(input string) ([]int, []int) {
	var openParens []int
	var closeParens []int
	for index, character := range input {
		strChar := string(character)
		if strChar == "(" {
			openParens = append(openParens, index)
		} else if strChar == ")" {
			closeParens = append(closeParens, index)
		}
	}
	return openParens, closeParens
}

func findCurrentParens(openParens []int, closeParens []int) (int, int) {
	firstCloseParen := closeParens[0]
	for index, val := range openParens {
		if val > firstCloseParen {
			return openParens[index-1], firstCloseParen
		}
	}
	return openParens[len(openParens)-1], firstCloseParen
}

func reverse(input []rune) []rune {
	result := make([]rune, len(input))
	resultLen := len(result)
	for _, character := range input {
		resultLen--
		result[resultLen] = character
	}
	return result
}
