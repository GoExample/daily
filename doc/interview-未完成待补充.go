package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	runtime.GOMAXPROCS(1)
	wg := sync.WaitGroup{}
	wg.Add(20)
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println("i: ", i)
			wg.Done()
		}(i)
	}
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println("fi: ", i)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
