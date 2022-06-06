package main

import (
	"fmt"
	"os"
	"time"

	"github.com/beevik/ntp"
)

var (
	host   = "0.ru.pool.nt.org"
	format = time.UnixDate
)

func main() {
	t, err := getDate(host)
	if err != nil {
		fmt.Fprintln(os.Stderr, "invalid response:", err)
		os.Exit(1)
	}

	fmt.Println(t.Format(format))
}

func getDate(host string) (time.Time, error) {
	t, err := ntp.Time(host)
	if err != nil {
		return time.Time{}, err
	}

	return t, nil
}
