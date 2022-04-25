package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var wg sync.WaitGroup
	var catCounter uint64
	var dogCounter uint64
	var fishCounter uint64
	var catCh = make(chan struct{}, 1)
	var dogCh = make(chan struct{}, 1)
	var fishCh = make(chan struct{}, 1)

	wg.Add(3)
	go printCat(&wg, catCounter, catCh, dogCh)
	go printDog(&wg, dogCounter, dogCh, fishCh)
	go printFish(&wg, fishCounter, fishCh, catCh)

	catCh <- struct{}{}
	wg.Wait()
}

func printCat(wg *sync.WaitGroup, counter uint64, catch, dogCh chan struct{}) {
	for {
		if counter >= uint64(100) {
			wg.Done()
			return
		}
		<- catch
		fmt.Println("cat")
		atomic.AddUint64(&counter, 1)
		dogCh <- struct{}{}
	}
}

func printDog(wg *sync.WaitGroup, counter uint64, dogCh, fishCh chan struct{}) {
	for {
		if counter >= uint64(100) {
			wg.Done()
			return
		}
		<- dogCh
		fmt.Println("dog")
		atomic.AddUint64(&counter, 1)
		fishCh <- struct{}{}
	}
}

func printFish(wg *sync.WaitGroup, counter uint64, fishCh, catCh chan struct{}) {
	for {
		if counter >= uint64(100) {
			wg.Done()
			return
		}
		<- fishCh
		fmt.Println("fish")
		atomic.AddUint64(&counter, 1)
		catCh <- struct{}{}
	}
}
