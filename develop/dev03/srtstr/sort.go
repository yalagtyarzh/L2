package srtstr

import (
	"fmt"
	"reflect"
	"sort"
	"strings"

	"dev03/config"
	"dev03/utils"
)

func Sort(strs []string, params config.Flags) []string {
	s := make([]string, 0)
	cols := make([][]string, 0)
	colNumber := params.Column

	for _, v := range strs {
		if len(strings.Fields(v)) >= colNumber && colNumber != 0 {
			cols = append(cols, strings.Fields(v))
		} else {
			s = append(s, v)
		}
	}

	//fmt.Println(cols)

	// Если не лень, солнце, добей сортировку с выбранным столбцом + числовое значение
	// ТОЧЬ-В-ТОЧЬ как в юниксе, (просто в начале выводятся ВСЕ СТРОКИ В ОТСОРТИРОВАННОМ
	// ВИДЕ, КОТОРЫЕ НЕ УДОВЛЯТВОРЕЮТ КОЛИЧЕСТВУ СТОЛБЦОВ, ПОТОМ СТРОКИ С УДОВЛЕТВОРЯЮЩИМ
	// СТОЛБЦОМ С ОТРИЦАТЕЛЬНЫМ ЗНАЧЕНИЕМ, ПОТОМ СТРОКИ С НУЖНЫМ СТОЛБЦОМ, НО БЕЗ ЧИСЛА,
	// ПОТОМ СТРОКИ С НУЖНЫМ СТОЛБЦОМ И ПОЛОЖИТЕЛЬНЫМ ЗНАЧЕНИЕМ)
	sort.SliceStable(
		s, func(i, j int) bool {
			if params.Num && params.Column == 0 {
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

	if len(cols) != 0 {
		sort.SliceStable(
			cols, func(i, j int) bool {
				if params.Num {
					inum, _ := utils.GetNumFromString(cols[i][colNumber-1])
					jnum, _ := utils.GetNumFromString(cols[j][colNumber-1])
					if jnum == 0 && inum == 0 {
						return cols[i][colNumber-1] < cols[j][colNumber-1]
					}

					return inum < jnum
				}

				return cols[i][colNumber-1] < cols[j][colNumber-1]
			},
		)
	}

	rcols := make([]string, 0)
	for _, v := range cols {
		rcols = append(rcols, strings.Join(v, " "))
	}

	res := make([]string, 0)
	res = append(res, s...)
	res = append(res, rcols...)

	if params.IsSorted {
		if reflect.DeepEqual(res, strs) {
			fmt.Println("Sorted")
		} else {
			fmt.Println("Not sorted")
		}
	}

	if params.Unique {
		utils.RemoveDuplicates(res)
	}

	if params.Reverse {
		utils.Reverse(res)
	}

	if params.B {
		for i, _ := range res {
			res[i] = strings.TrimSpace(res[i])
		}
	}

	return res
}
