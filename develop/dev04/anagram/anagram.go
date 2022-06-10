package anagram

import (
	"sort"
	"strings"
)

func GenerateAnagram(input []string) map[string][]string {
	res := fillMap(input)
	changeKeys(res)
	deleteNotAnagrams(res)
	for k, _ := range res {
		sort.Strings(res[k])
	}

	return res
}

func fillMap(input []string) map[string][]string {
	res := make(map[string][]string)
	checked := make(map[string]struct{})

	for _, v := range input {
		v = strings.ToLower(v)
		array := strings.Split(v, "")
		sort.Strings(array)
		if _, ok := checked[v]; !ok {
			res[strings.Join(array, "")] = append(res[strings.Join(array, "")], v)
			checked[v] = struct{}{}
		}
	}

	return res
}

func changeKeys(m map[string][]string) {
	c := copyMap(m)
	for k, v := range c {
		m[v[0]] = v
		delete(m, k)
	}
}

func deleteNotAnagrams(m map[string][]string) {
	for k, v := range m {
		if len(v) < 2 {
			delete(m, k)
		}
	}
}

func copyMap(m map[string][]string) map[string][]string {
	res := make(map[string][]string)
	for k, v := range m {
		res[k] = v
	}

	return res
}
