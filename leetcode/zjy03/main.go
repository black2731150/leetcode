package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	s := make(chan struct{})
	wg.Add(2)

	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println("协程1: 1")
			s <- struct{}{}
			<-s
		}
		wg.Done()
	}()

	go func() {
		for i := 0; i < 5; i++ {
			<-s
			fmt.Println("协程2: 2")
			s <- struct{}{}
		}
		wg.Done()
	}()

	wg.Wait()
	close(s)
}
