package main

import (
	"flag"
	"log"

	"dev10/telnet"
)

func main() {
	host := flag.String("host", "localhost", "хост")
	port := flag.String("port", "3000", "порт")
	timeout := flag.Int("timeout", 10, "таймаут на подключение к серверу")

	flag.Parse()

	client, err := telnet.NewTelnetClient(*host, *port, *timeout)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Conn.Close()

	client.Start()
}
