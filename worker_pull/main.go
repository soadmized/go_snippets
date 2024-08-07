package main

import (
	"fmt"
	"sync"
)

func main() {
	const workersCount = 5
	const tasksCount = 100

	res := make(chan string)
	input := make(chan int)

	go produceTasks(input, tasksCount)
	go processTasks(input, res, workersCount)

	for result := range res {
		fmt.Println(result)
	}

}

func produceTasks(ch chan int, tasks int) {
	for i := range tasks {
		ch <- i
	}

	close(ch)
}

func processTasks(inputChan chan int, resChan chan string, workersCount int) {
	wg := sync.WaitGroup{}
	wg.Add(workersCount)

	for i := range workersCount {
		go func(i int) {
			defer wg.Done()

			for val := range inputChan {
				res := fmt.Sprintf("task = %d, goroutine = %d", val, i)

				resChan <- res
			}
		}(i)
	}

	wg.Wait()
	close(resChan)
}
