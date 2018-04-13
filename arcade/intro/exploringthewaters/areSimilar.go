package exploringthewaters

import (
	"reflect"
)

func areSimilar(a []int, b []int) bool {
	moveIndex1 := -1
	moveIndex2 := -1

	for index := 0; index < len(a); index++ {
		if !containsInt(b, a[index]) {
			return false
		}
		if !containsInt(a, b[index]) {
			return false
		}
		if a[index] != b[index] {
			if moveIndex1 == -1 {
				moveIndex1 = index
			} else if moveIndex2 == -1 {
				moveIndex2 = index
				a = performMove(a, moveIndex1, moveIndex2)
				return reflect.DeepEqual(a, b)
			}
		}
	}
	return true
}

func performMove(a []int, index1 int, index2 int) []int {
	x := a[index1]
	y := a[index2]

	a = append(a[:index1], a[index1+1:]...) // delete element
	a = insertAfter(a, index1, y)
	a = append(a[:index2], a[index2+1:]...) // delete element
	a = insertAfter(a, index2, x)
	return a
}

func insertAfter(collection []int, loc int, data int) []int {
	collection = append(collection, 0)
	copy(collection[loc+1:], collection[loc:])
	collection[loc] = data
	return collection
}

func containsInt(input []int, val int) bool {
	for _, v := range input {
		if v == val {
			return true
		}
	}
	return false
}
