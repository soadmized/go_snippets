package main

import (
	"fmt"
	"sync"
)

const (
	buffer     = 5
	tasksCount = 100
)

func main() {
	sem := semaphore{ch: make(chan struct{}, buffer)}
	res := make(chan string)
	input := make(chan int)

	go produceTasks(input, tasksCount)
	go processTasks(input, res, sem)

	for result := range res {
		fmt.Println(result)
	}
}

type semaphore struct {
	ch chan struct{}
}

func (s *semaphore) acquire() {
	s.ch <- struct{}{}
}

func (s *semaphore) release() {
	<-s.ch
}

func produceTasks(ch chan int, tasks int) {
	for i := range tasks {
		ch <- i
	}

	close(ch)
}

func processTasks(inputChan chan int, resChan chan string, sem semaphore) {
	wg := sync.WaitGroup{}
	wg.Add(tasksCount)

	for task := range inputChan {
		go func(task int) {
			sem.acquire()
			defer func() {
				sem.release()
				wg.Done()
			}()

			res := fmt.Sprintf("task = %d", task)
			resChan <- res
		}(task)
	}

	wg.Wait()
	close(resChan)
}
