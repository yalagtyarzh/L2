package cat

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"dev06/config"
)

func StartCat(params config.Flags) {
	output := make([][]string, 0)
	neededCols := GetIntegersFromString(params.Fields)

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("$ ")
		row, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		if row == "\\end\n" {
			for _, v := range output {
				fmt.Print(strings.Join(v, " "))
			}

			break
		}

		if params.SeparatedOnly {
			if !strings.Contains(row, params.Delimiter) {
				continue
			}
		}

		splitRow := strings.Split(row, params.Delimiter)

		if neededCols != nil {
			frow := make([]string, 0)
			for i, v := range splitRow {
				if _, ok := neededCols[i+1]; ok {
					frow = append(frow, v)
				}
			}

			output = append(output, frow)
		} else {
			output = append(output, splitRow)
		}
	}
}

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
