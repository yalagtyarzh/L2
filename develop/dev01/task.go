package main

import (
	"fmt"
	"os"
	"time"

	"dev01/clock"
)

var (
	host   = "0.ru.pool.nt.org"
	format = time.UnixDate
)

func main() {
	t, err := clock.GetDate(host)
	if err != nil {
		fmt.Fprintln(os.Stderr, "invalid response:", err)
		os.Exit(1)
	}

	fmt.Println(t.Format(format))
}
