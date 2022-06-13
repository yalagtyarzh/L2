package main

import (
	"fmt"
	"sync"
	"time"
)

//Первый вариант - с помощью горутин и WaitGroup
func or(chans ...<-chan interface{}) <-chan interface{} {
	out := make(chan interface{})

	var wg sync.WaitGroup
	wg.Add(len(chans))

	for _, c := range chans {
		go func(c <-chan interface{}) {
			for v := range c {
				out <- v
			}
			wg.Done()
		}(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

// Второй вариант - с помощью reflect.Select
//func or(chans ...<-chan interface{}) <-chan interface{} {
//	out := make(chan interface{})
//
//	go func() {
//		defer close(out)
//		var cases []reflect.SelectCase
//		for _, c := range chans {
//			cases = append(
//				cases, reflect.SelectCase{
//					Dir:  reflect.SelectRecv,
//					Chan: reflect.ValueOf(c),
//				},
//			)
//		}
//
//		for len(cases) > 0 {
//			i, v, ok := reflect.Select(cases)
//			if !ok {
//				cases = append(cases[:i], cases[i+1:]...)
//				continue
//			}
//			out <- v.Interface()
//		}
//	}()
//
//	return out
//}

func main() {
	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}

	start := time.Now()
	<-or(
		sig(2*time.Second),
		sig(5*time.Second),
		sig(1*time.Second),
		sig(2*time.Second),
		sig(1*time.Second),
		sig(1*time.Second),
	)

	fmt.Printf("fone after %v", time.Since(start))
}
