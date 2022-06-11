package grep

import (
	"fmt"
	"strings"

	"dev05/config"
	"dev05/utils"
)

func Grep(pattern string, input []string, params config.Flags) map[int]string {
	if params.IgnoreCase {
		pattern = strings.ToLower(pattern)
		//pattern = utils.PreAppend(pattern, " ")
	}

	res := make(map[int]string)
	for i, v := range input {
		str := v
		if params.IgnoreCase {
			str = strings.ToLower(str)
		}

		var found bool
		if params.Fixed {
			found = str == pattern
		} else {
			found = strings.Contains(str, pattern)
		}

		if !found == params.Invert {
			res[i] = v
		}
	}

	if params.Count {
		fmt.Println(len(res))
		return res
	}

	before := utils.FindMax(params.Before, params.Context)
	after := utils.FindMax(params.After, params.Context)

	GetStringsAroundFound(res, input, before, after)

	return res
}

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
