package main

import (
	"fmt"
	"github.com/panjf2000/ants/v2"
	"os"
	"strconv"
	"sync"
	"time"
)

func main() {
	args := os.Args
	num := 0
	if len(args) >= 2 {
		num, _ = strconv.Atoi(args[1])
	}

	if num == 0 {
		num = 10
	}

	nonblocking(num)
	blocking(num)
}

func blocking(num int) {
	poolSize := num / 2
	waitSize := num / 2
	p, _ := ants.NewPool(poolSize, ants.WithMaxBlockingTasks(waitSize))
	defer p.Release()

	var wg sync.WaitGroup
	for i := 1; i <= num; i++ {
		wg.Add(1)
		go func(i int) {
			err := p.Submit(wrapper(i, &wg))
			if err != nil {
				fmt.Printf("task:%d err:%v\n", i, err)
				wg.Done()
			}
		}(i)
	}
	wg.Wait()
}

func nonblocking(num int) {
	poolSize := num / 2
	p, _ := ants.NewPool(poolSize, ants.WithNonblocking(true))
	defer p.Release()

	var wg sync.WaitGroup
	for i := 1; i <= num; i++ {
		wg.Add(1)
		err := p.Submit(wrapper(i, &wg))
		if err != nil {
			fmt.Printf("task:%d err:%v\n", i, err)
			wg.Done()
		}
	}

	wg.Wait()
}

func wrapper(i int, wg *sync.WaitGroup) func() {
	return func() {
		fmt.Printf("hello from task:%d\n", i)
		time.Sleep(1 * time.Second)
		wg.Done()
	}
}
