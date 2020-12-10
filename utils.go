package gsm

func compare(a, b string) int {
	if len(a) != len(b) {
		return 1
	}
	return 0
}

func min(a, b uint16) uint16 {
	if a < b {
		return a
	}
	return b
}
