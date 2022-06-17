package main

import (
	"flag"
	"log"

	"strategy/strategy"
)

func main() {
	output := flag.String("output", "console", "")

	flag.Parse()

	var activeStrategy strategy.PrintStrategy

	// В зависимости от инпута пользователя меняем стратегию рисования квадрата
	switch *output {
	case "console":
		activeStrategy = &strategy.ConsoleSquare{}
	case "image":
		activeStrategy = &strategy.ImageSquare{DestinationFilePath: "/tmp/image.jpg"}
	default:
		activeStrategy = &strategy.ConsoleSquare{}
	}

	err := activeStrategy.Print()
	if err != nil {
		log.Fatal(err)
	}
}
