package gsm

import "fmt"

const inValidDistance = -1

func HammingDistanceByteCount(a, b string) (int, error) {
	distance := 0
	if eq := compare(a, b); eq != 0 {
		return inValidDistance, fmt.Errorf("two string has different length")
	}
	for i := range a {
		if b[i] != a[i] {
			distance++
		}
	}
	return distance, nil
}

func HammingDistanceBitCount(a, b string) (int, error) {
	distance := 0
	if eq := compare(a, b); eq != 0 {
		return inValidDistance, fmt.Errorf("two string has different length")
	}
	ab, bb := []byte(a), []byte(b)
	for i := range bb {
		xored := bb[i] ^ ab[i]
		for j := 0; j < 8; j++ {
			distance += int(xored & 1)
			xored = xored >> 1
		}
	}
	return distance, nil
}
