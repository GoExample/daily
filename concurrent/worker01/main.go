package main

import "C"
import (
    "fmt"
    "log"
    "runtime"
    "sync"
    "time"
)

func initLog() {
    log.SetFlags(log.Lshortfile | log.LstdFlags)
}

func main() {
    initLog()
    log.Printf("Have %d goroutine running\n", runtime.NumGoroutine())

    tasks := []*Task{
        NewTask(func() error { return nil }),
        NewTask(func() error { return nil }),
        NewTask(func() error { return nil }),
    }

    p := NewPool(tasks, 50)
    p.Run()
    var numErrors int
    for _, task := range p.Tasks {
        if task.Err != nil {
            log.Println(task.Err)
            numErrors++
        }
        if numErrors >= 10 {
            log.Println("Too many errors.")
            break
        }
    }
    log.Printf("Have %d goroutine running\n", runtime.NumGoroutine())
    jobs := make(chan int, 100)
    results := make(chan int, 100)
    errors := make(chan error, 100)

    var wg sync.WaitGroup
    for w := 1; w <= 3; w++ {
        go worker(w, &wg, jobs, results, errors)
    }
    log.Printf("Have %d goroutine running\n", runtime.NumGoroutine())

    for j := 1; j <= 9; j++ {
        jobs <- j
        wg.Add(1)
    }
    close(jobs)

    wg.Wait()
    log.Printf("Have %d goroutine running\n", runtime.NumGoroutine())

    for i := 1; i <= 100; i++ {
        select {
        case err := <-errors:
            log.Println("finished with error: ", err)
        case result := <-results:
            log.Println("result is ", result)
        default:
        }
    }
    log.Printf("Have %d goroutine running\n", runtime.NumGoroutine())

    close(errors)
    close(results)
    log.Printf("Have %d goroutine running\n", runtime.NumGoroutine())

}

func worker(id int, wg *sync.WaitGroup, jobs <-chan int, results chan<- int, errors chan<- error) {
    for j := range jobs {
        log.Println("worker", id, "processing job", j)
        time.Sleep(time.Second)
        if j%2 == 0 {
            results <- j * 2
        } else {
            errors <- fmt.Errorf("error on job %v", j)
        }
        wg.Done()
    }
}
