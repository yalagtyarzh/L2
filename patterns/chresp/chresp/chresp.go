package chresp

import (
	"fmt"
	"io"
	"strings"
)

// ChainLogger - интерфейс, позволяющий объектам, имплементирующих данных интерфейс,
// при определенных условиях вызывать логику следующего объекта на который ссылается исходный
type ChainLogger interface {
	Next(s string)
}

// FirstLogger - объект, который хранит в себе ссылку на следующий логгер
type FirstLogger struct {
	NextChain ChainLogger
}

// Next выводит s на экран и, если имеется ссылка на следующий логгер - вызывает у того метод Next
func (l *FirstLogger) Next(s string) {
	fmt.Printf("First logger: %s\n", s)

	if l.NextChain != nil {
		l.NextChain.Next(s)
	}
}

// SecondLogger - объект, который хранит в себе ссылку на следующий логгер
type SecondLogger struct {
	NextChain ChainLogger
}

// Next выводит s на экран, если в s имеется подстрока hello, а после по цепи вызывает метод Next у следующего логгера
// В иных просто выводит информацию на экран и не вызывает метод у следующей цепи
func (l *SecondLogger) Next(s string) {
	if strings.Contains(strings.ToLower(s), "hello") {
		fmt.Printf("Second logger: %s\n", s)

		if l.NextChain != nil {
			l.NextChain.Next(s)
		}

		return
	}

	fmt.Printf("Finishing in second logging\n")
}

// WriterLogger - объект, который хранит в себе ссылку на следующий логгер и писателя
type WriterLogger struct {
	NextChain ChainLogger
	Writer    io.Writer
}

// Next пишет s с помощью писателя и вызывает следующую цепочку, если она существует
func (l *WriterLogger) Next(s string) {
	if l.Writer != nil {
		l.Writer.Write([]byte("WriterLogger: " + s))
	}

	if l.NextChain != nil {
		l.NextChain.Next(s)
	}
}
