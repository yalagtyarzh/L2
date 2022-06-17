package main

import (
	"os"

	"chresp/chresp"
)

func main() {
	logger := chresp.FirstLogger{NextChain: nil}
	logger.Next("hi")

	slogger := chresp.SecondLogger{NextChain: nil}

	logger.NextChain = &slogger
	logger.Next("a")

	slogger.NextChain = &chresp.WriterLogger{
		NextChain: nil,
		Writer:    os.Stdin,
	}

	logger.Next("hello")
}
