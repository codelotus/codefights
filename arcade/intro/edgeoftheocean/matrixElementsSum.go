package edgeoftheocean

func matrixElementsSum(matrix [][]int) int {
	total := 0

	nrow := len(matrix)
	if nrow > 0 {
		ncol := len(matrix[0])
		for col := 0; col < ncol; col++ { // iterate by colum
			for row := 0; row < nrow; row++ {
				room := matrix[row][col]
				if room == 0 { // skip remaining column
					break
				} else {
					total += room
				}
			}
		}
	}

	return total
}
