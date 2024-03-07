package main

import "sync"

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	s := make(chan struct{})

	go func() {
		<-s
		s <- struct{}{}
	}()

	go func() {
		<-s
		s <- struct{}{}
	}()

	wg.Wait()
}
