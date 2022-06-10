package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"

	"dev03/config"
	"dev03/scan"
	"dev03/srtstr"
)

func main() {
	flags := config.Flags{}

	flag.IntVar(&flags.Column, "k", 0, "колонки для сортировки, по умолчанию сортируем всю строку")
	flag.BoolVar(&flags.Reverse, "r", false, "сортировка в обратную сторону")
	flag.BoolVar(&flags.Unique, "u", false, "не выводить повторяющиеся строки")
	flag.BoolVar(&flags.Num, "n", false, "сортировать по числовому значению")
	flag.BoolVar(&flags.B, "b", false, "игнорировать хвостовые пробелы")
	flag.BoolVar(&flags.IsSorted, "c", false, "проверять отсортированы ли данные")

	flag.Parse()

	filename := flag.Arg(0)
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("invalid file:", err)
		os.Exit(1)
	}
	defer file.Close()

	strs, err := scan.Scan(file)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	strs = srtstr.Sort(strs, flags)
	output, err := os.Create("output.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer output.Close()

	w := bufio.NewWriter(output)
	for _, str := range strs {
		_, err := w.WriteString(str + "\n")
		if err != nil {
			log.Fatalln(err)
		}
	}

	w.Flush()
}
