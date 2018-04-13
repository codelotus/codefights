package smoothsailing

import "sort"

func sortByHeight(a []int) []int {
	treeLoc := locateTrees(a)
	result := removeTreesAndSort(a)
	return replaceTrees(treeLoc, result)
}

func replaceTrees(treeLoc []int, data []int) []int {
	for _, location := range treeLoc {
		data = insertAfter(data, location, -1)
	}
	return data
}

func removeTreesAndSort(data []int) []int {
	var result []int
	for _, element := range data {
		if element != -1 {
			result = append(result, element)
		}
	}
	sort.Ints(result)
	return result
}

func locateTrees(data []int) []int {
	var result []int
	for index, element := range data {
		if element == -1 {
			result = append(result, index)
		}
	}
	return result
}

func insertAfter(collection []int, loc int, data int) []int {
	collection = append(collection, 0)
	copy(collection[loc+1:], collection[loc:])
	collection[loc] = data
	return collection
}
