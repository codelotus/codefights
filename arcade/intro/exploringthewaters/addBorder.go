package exploringthewaters

import (
	"bytes"
)

func addBorder(picture []string) []string {
	result := []string{}

	for index, val := range picture {
		if index == 0 { // First Row
			result = append(result, buildHeaderFooter(len(val)))
			result = append(result, buildRow(val))
		} else { // Row
			result = append(result, buildRow(val))
		}
		if index == len(picture)-1 { // this is the last row
			result = append(result, buildHeaderFooter(len(val)))
		}
	}
	return result
}

func buildHeaderFooter(length int) string {
	result := bytes.NewBufferString("")

	for x := 0; x < length+2; x++ {
		result.WriteString("*")
	}
	return result.String()
}

func buildRow(input string) string {
	result := bytes.NewBufferString("*")
	for _, val := range input {
		result.WriteString(string(val))
	}
	result.WriteString("*")
	return result.String()
}
