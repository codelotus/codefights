package exploringthewaters

func arrayChange(inputArray []int) int {
	result := 0
	for index, val := range inputArray {
		if index != 0 {
			if val <= inputArray[index-1] {
				inputArray = append(inputArray[:index], inputArray[index+1:]...) // delete val at index
				// replace val at index incrementing by 1
				inputArray = append(inputArray, 0)
				copy(inputArray[index+1:], inputArray[index:])
				increaseBy := (inputArray[index-1] - val) + 1
				inputArray[index] = val + increaseBy
				result = result + increaseBy
			}
		}
	}
	return result
}
