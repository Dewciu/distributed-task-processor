package main

import (
	"flag"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Task struct {
	Id          int
	Name        string
	Description string
	ProcessTime time.Duration
}

func (t *Task) String() string {
	return fmt.Sprintf("Task with ID: %d, Name: %s, Description %s", t.Id, t.Name, t.Description)
}

type Worker struct {
	Id int
}

func (w *Worker) Process(wg *sync.WaitGroup, chTask <-chan Task) {
	defer wg.Done()

	for {
		select {
		case task, ok := <-chTask:
			if !ok {
				fmt.Printf("Worker %d: Task channel closed. Stopping.\n", w.Id)
				return
			}
			fmt.Printf("Worker %d is processing task %d for %s\n", w.Id, task.Id, task.ProcessTime.String())
			time.Sleep(task.ProcessTime)
			fmt.Printf("Task %d finished by worker %d\n", task.Id, w.Id)
		}
	}
}

func generateTask(wg *sync.WaitGroup, chTask chan<- Task, taskNumber int) {
	defer wg.Done()
	processTime := rand.Intn(5000)
	time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond)
	chTask <- Task{
		Id:          taskNumber,
		Name:        fmt.Sprintf("Task name: %d", taskNumber),
		Description: fmt.Sprintf("Task description: %d", taskNumber),
		ProcessTime: time.Duration(processTime) * time.Millisecond,
	}
}

func main() {
	taskCount := flag.Int("tasks", 20, "number of tasks to generate")
	workerCount := flag.Int("workers", 3, "number of workers to create")
	flag.Parse()

	fmt.Printf("Starting app with %d tasks and %d workers\n", *taskCount, *workerCount)

	var wg sync.WaitGroup
	var wgT sync.WaitGroup

	workers := []Worker{}

	for i := 0; i < *workerCount; i++ {
		workers = append(workers, Worker{Id: i})
	}

	taskChannel := make(chan Task, *taskCount)

	wg.Add(*workerCount + 1)
	wgT.Add(*taskCount)

	for _, worker := range workers {
		go worker.Process(
			&wg, taskChannel,
		)
	}

	go func() {
		wg.Done()
		for i := 0; i < *taskCount; i++ {
			go generateTask(&wgT, taskChannel, i)
		}
		wgT.Wait()
		close(taskChannel)
	}()

	wg.Wait()

	time.Sleep(time.Second)
}
