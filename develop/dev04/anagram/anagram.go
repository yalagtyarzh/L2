package anagram

import (
	"fmt"
	"sort"
	"strings"
)

// GenerateAnagram составляет словарь анаграммы и возвращает его
func GenerateAnagram(input []string) map[string][]string {
	res := fillMap(input)
	fmt.Println(res)
	changeKeys(res)
	deleteNotAnagrams(res)
	for k := range res {
		sort.Strings(res[k])
	}

	return res
}

func fillMap(input []string) map[string][]string {
	res := make(map[string][]string)
	checked := make(map[string]struct{})

	for _, v := range input {
		// Подготавливаем строку для записи в мапу
		v = strings.ToLower(v)
		array := strings.Split(v, "")
		// Сортируем символы строки в алфавитном порядке
		sort.Strings(array)
		// Записываем в мапу лишь те анаграммы, которые не имеют дубликатов
		if _, ok := checked[v]; !ok {
			res[strings.Join(array, "")] = append(res[strings.Join(array, "")], v)
			checked[v] = struct{}{}
		}
	}

	return res
}

// changeKeys меняет ключи мапы на первый встречающийся элемент в срезах, находящихся под этими ключами
func changeKeys(m map[string][]string) {
	c := copyMap(m)
	for k, v := range c {
		m[v[0]] = v
		delete(m, k)
	}
}

// deleteNotAnagrams удаляет те анаграммы, имеющие лишь одно слово в своей группе
func deleteNotAnagrams(m map[string][]string) {
	for k, v := range m {
		if len(v) < 2 {
			delete(m, k)
		}
	}
}

// copyMap копирует все ключи и значения одной мапы в другую
func copyMap(m map[string][]string) map[string][]string {
	res := make(map[string][]string)
	for k, v := range m {
		res[k] = v
	}

	return res
}
