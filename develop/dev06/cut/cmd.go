package cut

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"dev06/config"
)

// StartCut запускает shell утилиту, имитирующая input unix утилиты cat
func StartCut(params config.Flags) {
	output := make([][]string, 0)
	neededCols := GetIntegersFromString(params.Fields)

	// Стартуем наш Reader
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("$ ")
		row, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		// В случае выхода из программы - выводим все данные, удовлетворяющие флагам программы
		if row == "\\end\n" {
			for _, v := range output {
				fmt.Print(strings.Join(v, " "))
			}

			break
		}

		// Если задано, что выводим лишь строки с разделителем, то проверяем, если в заданной строке разделитель
		if params.SeparatedOnly {
			if !strings.Contains(row, params.Delimiter) {
				continue
			}
		}

		// Разделяем строку с помощью разделителя
		splitRow := strings.Split(row, params.Delimiter)

		if neededCols != nil {
			// Если заданы лишь определенные столбцы, то добавляем в output лишь определенные столбцы
			filtredRow := make([]string, 0)
			for i, v := range splitRow {
				if _, ok := neededCols[i+1]; ok {
					filtredRow = append(filtredRow, v)
				}
			}

			output = append(output, filtredRow)
		} else {
			// В иных случаях просто добавляем разделенную строку в output
			output = append(output, splitRow)
		}
	}
}

// GetIntegersFromString извлекает из строки числа и возвращает мапу, содержащие числа в качестве ключей
func GetIntegersFromString(s string) map[int]struct{} {
	if s == "" {
		return nil
	}

	nums := strings.Split(s, ",")
	res := make(map[int]struct{})
	for _, v := range nums {
		n, err := strconv.Atoi(v)
		if err != nil {
			continue
		}

		res[n] = struct{}{}
	}

	return res
}
