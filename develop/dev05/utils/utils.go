package utils

//func PreAppend(s string, r rune) string {
//	var res strings.Builder
//	res.WriteRune(r)
//	res.WriteString(s)
//	res.WriteRune(r)
//
//	return res.String()
//}

func FindMax(a, b uint) uint {
	if a > b {
		return a
	}

	return b
}

func CopyKeys(m map[int]string) []int {
	c := make([]int, 0)
	for k := range m {
		c = append(c, k)
	}

	return c
}
