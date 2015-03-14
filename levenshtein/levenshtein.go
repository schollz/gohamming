package levenshtein

// This function returns the hamming distance (int) for two strings
// returns -1 if the input is not of equal length
func Calc(one string, two string) int {
	// check length and bail
	if len(one) != len(two) {
		return -1
	}

	// dont loop if they are the same
	if one == two {
		return 0
	}
	var result int = 0
	for i, r := range one {
		if uint8(r) != two[i] {
			result++
		}
	}
	return result
}
