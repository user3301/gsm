package gsm

import "unicode/utf8"

func LevenshteinDistanceByteCount(a, b string) int {
	if a == b {
		return 0
	}
	if len(a) == 0 {
		return utf8.RuneCountInString(b)
	}
	if len(b) == 0 {
		return utf8.RuneCountInString(a)
	}

	ra, rb := []rune(a), []rune(b)

	if len(ra) > len(rb) {
		ra, rb = rb, ra
	}

	lenra, lenrb := len(ra), len(rb)

	matrix := make([]uint16, lenra+1)
	for i := 1; i < len(matrix); i++ {
		matrix[i] = uint16(i)
	}

	_ = matrix[lenra]

	for i := 1; i <= lenrb; i++ {
		previous := uint16(i)
		for j := 1; j <= lenra; j++ {
			current := matrix[j-1]
			if rb[i-1] != ra[j-1] {
				current = min(min(matrix[j-1]+1, previous+1), matrix[j]+1)
			}
			matrix[j-1] = previous
			previous = current
		}
		matrix[lenra] = previous
	}
	return int(matrix[lenra])
}
