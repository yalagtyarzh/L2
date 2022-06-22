package main

import (
	"flag"
	"fmt"
	"log"

	"dev09/download"
)

func main() {
	destination := flag.String("d", "./tmp", "destination for output")

	flag.Parse()

	args := flag.Args()
	if len(args) != 1 {
		fmt.Println("usage: task [option] [URL]")
		return
	}

	err := download.Download(*destination, args[0])
	if err != nil {
		log.Fatal(err)
	}

}
