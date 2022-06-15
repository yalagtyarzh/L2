package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	url := os.Args[1]
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("invalid url")
		os.Exit(1)
	}
	defer resp.Body.Close()

	f, err := os.Create("index.html")
	defer f.Close()

	io.Copy(f, resp.Body)
}
