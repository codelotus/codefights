package islandofknowledge

func areEquallyStrong(yourLeft int, yourRight int, friendsLeft int, friendsRight int) bool {
	if (yourLeft == friendsLeft ||
		yourLeft == friendsRight) &&
		(yourRight == friendsRight ||
			yourRight == friendsLeft) {
		return true
	}
	return false
}
