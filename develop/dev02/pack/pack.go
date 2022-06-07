package pack

import (
	"fmt"
	"strconv"
	"unicode"
)

func Unpack(str string) (string, error) {
	res := make([]rune, 0)
	runes := []rune(str)
	shield := false

	for i := 0; i < len(runes); i++ {
		// Обрабатываем неэкранированные цифры
		if unicode.IsDigit(runes[i]) && !shield {
			return "", fmt.Errorf("got unshielded digit in position %d", i+1)
		}

		// Обрабатываем последний символ
		if i+1 == len(runes) {
			if runes[i] == '\\' && !shield {
				return "", fmt.Errorf("unshielded backslash in end of string")
			}

			res = append(res, runes[i])
			break
		}

		// Обрабатываем бэкслыши
		if runes[i] == '\\' && shield == false {
			shield = true
			continue
		}

		next := runes[i+1]

		// Обработка всего остального xd
		if unicode.IsDigit(next) {
			num, err := strconv.Atoi(string(next))
			if err != nil {
				return "", fmt.Errorf("converting string->int error")
			}

			up := MultipleRunes(runes[i], num)
			res = append(res, up...)
			i++

			shield = false

			continue
		} else {
			shield = false

			res = append(res, runes[i])
		}
	}

	return string(res), nil
}

func MultipleRunes(r rune, mult int) []rune {
	res := make([]rune, mult)
	for i := range res {
		res[i] = r
	}

	return res
}
