package exploringthewaters

func alternatingSums(a []int) []int {
	teamA := 0
	teamB := 0
	for index, value := range a {
		if index%2 == 0 {
			teamA = teamA + value
		} else {
			teamB = teamB + value
		}
	}
	return []int{teamA, teamB}
}
