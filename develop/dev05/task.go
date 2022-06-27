package main

import (
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	"dev05/config"
	"dev05/grep"
	"dev05/pkg/fscan"
)

var flags config.Flags

func main() {
	flag.UintVar(&flags.After, "A", 0, "печатать +N строк после совпадения")
	flag.UintVar(&flags.Before, "B", 0, "печатать +N строк до совпадения")
	flag.UintVar(&flags.Context, "C", 0, "печатать +N строк вокруг совпадения")
	flag.BoolVar(&flags.Count, "c", false, "количество строк")
	flag.BoolVar(&flags.IgnoreCase, "i", false, "игнорировать регистр")
	flag.BoolVar(&flags.Invert, "v", false, "вместо совпадения, исключать")
	flag.BoolVar(&flags.Fixed, "F", false, "точное совпадение со строкой, не паттерн")
	flag.BoolVar(&flags.LineNum, "n", false, "напечатать номер строки")

	flag.Parse()

	args := flag.Args()
	if len(args) != 2 {
		fmt.Println("Usage: [flags] [pattern] [file]")
		os.Exit(1)
	}

	pattern := args[0]

	filename := args[1]
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("invalid file:", err)
		os.Exit(1)
	}
	defer file.Close()

	strs, err := fscan.Scan(file)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	res := grep.Grep(pattern, strs, flags)

	if res != nil {
		printMap(res)
	}
}

func printMap(m map[int]string) {
	keys := make([]int, 0)
	for k := range m {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	for _, key := range keys {
		var s strings.Builder
		if flags.LineNum {
			lineNum := fmt.Sprintf("%s:", strconv.Itoa(key+1))
			s.WriteString(lineNum)
		}

		s.WriteString(m[key])

		fmt.Println(s.String())
	}
}
