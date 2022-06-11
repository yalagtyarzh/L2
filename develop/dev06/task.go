package main

import (
	"flag"

	"dev06/cat"
	"dev06/config"
)

func main() {
	flags := config.Flags{}
	flag.StringVar(&flags.Fields, "f", "", "выбрать поля (колонки)")
	flag.StringVar(&flags.Delimiter, "d", "\t", "использовать другой разделитель")
	flag.BoolVar(&flags.SeparatedOnly, "s", false, "только строки с разделителем")

	flag.Parse()

	cat.StartCat(flags)
}
