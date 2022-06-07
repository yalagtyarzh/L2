package pack

import (
	"fmt"
	"strconv"
	"unicode"
)

func Unpack(str string) (string, error) {
	res := make([]rune, 0)
	runes := []rune(str)

	for i := 0; i < len(runes); i++ {
		if unicode.IsDigit(runes[i]) {
			return "", fmt.Errorf("got unshielded digit in position %d", i+1)
		}

		if i+1 == len(runes) {
			res = append(res, runes[i])
			break
		}

		next := runes[i+1]

		if unicode.IsDigit(next) {
			num, err := strconv.Atoi(string(next))
			if err != nil {
				return "", fmt.Errorf("converting string->int error")
			}

			up := MultipleRunes(runes[i], num)
			res = append(res, up...)
			i++

			continue
		} else {
			res = append(res, runes[i])
		}
	}

	return string(res), nil
}

func MultipleRunes(r rune, mult int) []rune {
	res := make([]rune, mult)
	for i, _ := range res {
		res[i] = r
	}

	return res
}
