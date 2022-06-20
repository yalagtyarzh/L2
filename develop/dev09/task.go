package main

import (
	"flag"
	"fmt"
	"log"
	"sync"

	"dev09/download"
)

var wg sync.WaitGroup

func main() {
	destination := flag.String("d", "./tmp", "destination for output")

	flag.Parse()

	urls := flag.Args()
	if len(urls) < 1 {
		fmt.Println("usage: task [option] [URL]...")
		return
	}

	for _, v := range urls {
		wg.Add(1)
		err := download.Store(*destination, v, &wg)
		if err != nil {
			log.Fatal(err)
		}
	}

	wg.Wait()
}
