package midterm

func compare(Strategies, pattern [][]string, row, col int) bool {
	for i := 0; i < len(pattern); i++ {
		for j := 0; j < len(pattern[i]); j++ {
			if pattern[i][j] != Strategies[row+i][col+j] {
				return false
			}
		}
	}
	return true
}

func paternLength(pattern [][]string) int {
	numcol := len(pattern[0])
	for i := range pattern {
		if len(pattern[i]) > numcol {
			numcol = len(pattern[i])
		}
	}
	return numcol
}

func CountPatternMatches(Strategies [][]string, pattern [][]string) int {
	var count int
	width := paternLength(pattern)
	for i := 0; i <= len(Strategies)-len(pattern); i++ {
		for j := 0; j <= len(Strategies[i])-width; j++ {
			if compare(Strategies, pattern, i, j) == true {
				count += 1
			}
		}
	}
	return count
}
