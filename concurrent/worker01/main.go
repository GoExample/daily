package main

import (
    "fmt"
    "sync"
    "time"
)

func main() {
    jobs := make(chan int, 100)
    results := make(chan int, 100)
    errors := make(chan error, 100)

    var wg sync.WaitGroup
    for w := 1; w <= 3; w++ {
        go worker(w, &wg, jobs, results, errors)
    }

    for j := 1; j <= 9; j++ {
        jobs <- j
        wg.Add(1)
    }
    close(jobs)

    wg.Wait()

    for i := 1; i <= 100; i++ {
        select {
        case err := <-errors:
            fmt.Println("finished with error: ", err)
        case result := <-results:
            fmt.Println("result is ", result)
        default:
        }
    }
    close(errors)
    close(results)
}

func worker(id int, wg *sync.WaitGroup, jobs <-chan int, results chan<- int, errors chan<- error) {
    for j := range jobs {
        fmt.Println("worker", id, "processing job", j)
        time.Sleep(time.Second)
        if j%2 == 0 {
            results <- j * 2
        } else {
            errors <- fmt.Errorf("error on job %v", j)
        }
        wg.Done()
    }
}
