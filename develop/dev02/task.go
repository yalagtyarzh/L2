package main

import (
	"fmt"
	"os"

	"dev02/pack"
)

func main() {
	input := os.Args[1:]
	if len(input) != 1 {
		fmt.Println("Please enter only one argument")
		os.Exit(1)
	}
	res, err := pack.Unpack(input[0])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(res)
}
