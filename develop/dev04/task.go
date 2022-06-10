package main

import (
	"fmt"
	"strings"

	"dev04/anagram"
)

func main() {
	input := []string{"ПЯТКА", "ПЯТАК", "ТЯПКА", "ЛИСТОК", "актяп", "СЛИТОК", "СТОЛИК", "xd"}
	m := anagram.GenerateAnagram(input)
	for k, v := range m {
		fmt.Printf("%s: %s\n", k, strings.Join(v, ", "))
	}
}
