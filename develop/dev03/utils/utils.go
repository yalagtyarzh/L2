package utils

import (
	"strconv"
	"unicode"
)

// Reverse зеркально меняет местами элементы среза строк
func Reverse(strs []string) {
	// Получаем "словарь" строки
	for i, j := 0, len(strs)-1; i < len(strs)/2; i, j = i+1, j-1 {
		// Меняем слова местами
		strs[i], strs[j] = strs[j], strs[i]
	}
}

// RemoveDuplicates создает новый срез строк без дубликатов
func RemoveDuplicates(strs []string) []string {
	res := make([]string, 0)
	for _, v := range strs {
		if !Contains(res, v) {
			res = append(res, v)
		}
	}

	return res
}

// Contains проверяет наличие строки в срезе строк
func Contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

// GetNumFromString составляет посимвольно число в виде строки до момента встречи символа, не являющейся цифровой,
// а позже преобразовывает получившуюся строку в число
func GetNumFromString(s string) (float64, error) {
	var runes []rune
	for i, v := range s {
		if !unicode.IsDigit(v) && !(v == '-' && i == 0) && v != '.' {
			break
		}

		runes = append(runes, v)
	}

	num, err := strconv.ParseFloat(string(runes), 64)
	if err != nil {
		return num, err
	}

	return num, nil
}
