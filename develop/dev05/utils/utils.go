package utils

//func PreAppend(s string, r rune) string {
//	var res strings.Builder
//	res.WriteRune(r)
//	res.WriteString(s)
//	res.WriteRune(r)
//
//	return res.String()
//}

// FindMax возвращает наибольшее число из переданных двух
func FindMax(a, b uint) uint {
	if a > b {
		return a
	}

	return b
}

// CopyKeys возвращает все ключи мапы в качестве среза
func CopyKeys(m map[int]string) []int {
	c := make([]int, 0)
	for k := range m {
		c = append(c, k)
	}

	return c
}
