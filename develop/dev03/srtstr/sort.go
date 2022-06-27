package srtstr

import (
	"fmt"
	"reflect"
	"sort"
	"strings"

	"dev03/config"
	"dev03/utils"
)

// Sort сортирует строки с помощью переданных параметров, хранящихся в params
func Sort(strs []string, params config.Flags) []string {
	s := make([]string, 0)
	cells := make([][]string, 0)
	colNumber := params.Column

	// Конструируем строки и ячейки для сортировки
	for _, v := range strs {
		if len(strings.Fields(v)) >= colNumber && colNumber != 0 {
			// Если задано, что мы сортируем по столбцам, и при этом наша строка имеет равное или большое количество
			// столбцов, чем требуется - заполняем массив ячеек
			cells = append(cells, strings.Fields(v))
		} else {
			// Иначе заполняем массив строк
			s = append(s, v)
		}
	}

	// Алгоритм сортировки для строк
	sort.SliceStable(
		s, func(i, j int) bool {
			if params.Num {
				inum, _ := utils.GetNumFromString(s[i])
				jnum, _ := utils.GetNumFromString(s[j])
				if jnum == 0 && inum == 0 {
					return s[i] < s[j]
				}

				return inum < jnum
			}

			return s[i] < s[j]
		},
	)

	// Алгоритм сортировки для полей
	if len(cells) != 0 {
		sort.SliceStable(
			cells, func(i, j int) bool {
				if params.Num {
					inum, _ := utils.GetNumFromString(cells[i][colNumber-1])
					jnum, _ := utils.GetNumFromString(cells[j][colNumber-1])
					if jnum == 0 && inum == 0 {
						return cells[i][colNumber-1] < cells[j][colNumber-1]
					}

					return inum < jnum
				}

				return cells[i][colNumber-1] < cells[j][colNumber-1]
			},
		)
	}

	// Объединяем отдельные ячейки в строки
	rows := make([]string, 0)
	for _, v := range cells {
		rows = append(rows, strings.Join(v, " "))
	}

	// Составляем результат сортировки
	res := make([]string, 0)
	res = append(res, s...)
	res = append(res, rows...)

	// Если задан флаг, то проверяем, одинаковы ли у нас получился результат с входными данными
	if params.IsSorted {
		if reflect.DeepEqual(res, strs) {
			fmt.Println("Sorted")
		} else {
			fmt.Println("Not sorted")
		}
	}

	// Если задан флаг на уникальность - удаляем дубликаты
	if params.Unique {
		res = utils.RemoveDuplicates(res)
	}

	// Если задан флаг обратной сортировки - просто меняем местами все элементы среза
	if params.Reverse {
		utils.Reverse(res)
	}

	// Если задан параметр игнора хвостовых пробелов - стираем их
	if params.B {
		for i := range res {
			res[i] = strings.TrimSpace(res[i])
		}
	}

	return res
}
