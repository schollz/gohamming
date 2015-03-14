package hamming

func Calc(one string, two string) int {
	// check length and bail
	if len(one) != len(two) {
		return -1
	}

	var result int = 0
	for i, r := range one {
		if uint8(r) != two[i] {
			result++
		}
	}
	return result
}
