package exploringthewaters

func palindromeRearranging(inputString string) bool {
	if len(inputString) <= 1 {
		return true
	}
	keys := map[string]int{}
	for _, rkey := range inputString {
		key := string(rkey)
		if _, ok := keys[key]; ok { // check if key exists
			keys[key] = keys[key] + 1
		} else {
			// key does not exist, add it to the map
			keys[key] = 1
		}
	}
	oddCount := 0
	for _, v := range keys {
		if v%2 != 0 {
			oddCount++
			if oddCount == 2 {
				return false
			}
		}
	}
	return true
}
