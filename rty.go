package main

import (
	"log"
	"time"
)

type RowInfo struct {
	id int64
}

var QueueChan chan RowInfo

func main() {
	QueueChan := make(chan RowInfo)
	workerChan := make(chan RowInfo)
	exitChan := make(chan int64)

	go dispatcher(QueueChan, workerChan, exitChan)

	// Start WorkerCount number of workers
	workerCount := 4
	for i := 0; i < workerCount; i++ {
		go worker(workerChan, exitChan)
	}

	// Send test data
	for i := 0; i < 12; i++ {
		QueueChan <- RowInfo{id: int64(i)}
	}

	// Prevent app close
	for {
		time.Sleep(1 * time.Second)
	}
}

func dispatcher(queueChan, workerChan chan RowInfo, exitChan chan int64) {
	state := make(map[int64]bool)

	for {
		select {
		case job := <-QueueChan:
			if state[job.id] == true {
				continue
			}
			workerChan <- job
		case result := <-exitChan:
			state[result] = false
		}
	}
}

func worker(workerChan chan RowInfo, exitChan chan int64) {
	for job := range workerChan {
		log.Printf("Doing work on job rowInfo ID: %d", job.id)

		// Finish job
		exitChan <- job.id
	}
}
