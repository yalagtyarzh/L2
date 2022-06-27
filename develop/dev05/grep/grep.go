package grep

import (
	"fmt"
	"strings"

	"dev05/config"
	"dev05/utils"
)

// Grep ищет строки, в котором есть паттерн переданный нами паттерн
func Grep(pattern string, input []string, params config.Flags) map[int]string {
	// Если задано, что игнорируем регистр - приводим паттерн к нижнему регистру
	if params.IgnoreCase {
		pattern = strings.ToLower(pattern)
	}

	res := make(map[int]string)
	for i, v := range input {
		str := v
		if params.IgnoreCase {
			str = strings.ToLower(str)
		}

		var found bool
		if params.Fixed {
			// Если ищем конкретную строку, а не паттерн, то осуществляем сравнивание строк
			found = str == pattern
			//pattern = utils.PreAppend(pattern, " ")
		} else {
			// В иных случаях проверяем на наличие заданной строки в искомой строке
			found = strings.Contains(str, pattern)
		}

		// Заполняем срез результатов с учетом заданного параметра Invert
		if !found == params.Invert {
			res[i] = v
		}
	}

	// Если задано выдать лишь количество найденных строк - выводим его размер и выходим из функции
	if params.Count {
		fmt.Println(len(res))
		return nil
	}

	// Определяем область поиска строк
	before := utils.FindMax(params.Before, params.Context)
	after := utils.FindMax(params.After, params.Context)

	// Заполняем результирующий срез теми строками, которые находятся "вокруг" строк, имеющих нужный паттерн
	GetStringsAroundFound(res, input, before, after)

	return res
}

// GetStringsAroundFound находит все строки, которые находятся "вокруг" найденных строк
func GetStringsAroundFound(m map[int]string, input []string, before uint, after uint) {
	c := utils.CopyKeys(m)

	for _, v := range c {
		for j := 1; j <= int(before); j++ {
			if v-j < 0 {
				continue
			}

			m[v-j] = input[v-j]
		}

		for j := 1; j <= int(after); j++ {
			if v+j > len(input)-1 {
				continue
			}

			m[v+j] = input[v+j]
		}
	}
}
