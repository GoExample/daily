package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(100)

	exit := false
	go func() {
		wg.Add(1)
		defer wg.Done()

		for i := 0; i < 100; i++ {
			fmt.Println("fuck")
			time.Sleep(100)
		}
		exit = true
	}()

	wg.Add(1)
	//time.Sleep(100000)
	wg.Wait()
	//for !exit {
	//	return
	//}
}
