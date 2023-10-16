package main

import (
	"log"
	"time"
)

type Job interface {
	Run()
}

type JobChan chan Job

type Worker struct {
	id    int
	JobCh JobChan   // 用于将任务分配给相关的worker
	quit  chan bool // 表明退出
}

func NewWorker(index int) *Worker {
	return &Worker{
		id:    index,
		JobCh: make(JobChan),
		quit:  make(chan bool),
	}
}

func (w *Worker) Start(wp *WorkerPool) {
	go func() {
		for {
			wp.WorkerQueue <- w // if current worker finish last job then it will register itself into workerpool
			select {
			case job := <-w.JobCh:
				log.Printf("Worker%d is running...\n", w.id)
				job.Run()
			case <-w.quit:
				return
			default:
				w.quit <- true
			}
		}
	}()
}

type WorkerPool struct {
	workersize  int
	JobQueue    JobChan
	WorkerQueue chan *Worker
}

func NewWorkerPool(size int, jobsize int) *WorkerPool {
	return &WorkerPool{
		workersize:  size,
		JobQueue:    make(JobChan, jobsize),
		WorkerQueue: make(chan *Worker, size),
	}
}

func (wp *WorkerPool) Start() {
	// create and register
	for i := 0; i < wp.workersize; i++ {
		worker := NewWorker(i)
		worker.Start(wp)
	}

	go func() {
		for {
			// allocate job to worker
			select {
			case job := <-wp.JobQueue:
				w := <-wp.WorkerQueue
				w.JobCh <- job
			}
		}
	}()
}

type Task struct {
	Number int
}

func (t *Task) Run() {
	time.Sleep(time.Second * 2)
	log.Printf("job %d is done\n", t.Number)
}

func main() {
	workSize := 100
	jobChanSize := 1000
	jobSize := 100 * 20
	// create workerpool
	wp := NewWorkerPool(workSize, jobChanSize)

	wp.Start()

	for i := 0; i < jobSize; i++ {
		go func(i int) {
			task := &Task{Number: i}
			wp.JobQueue <- task
		}(i)
	}

	time.Sleep(time.Second * 3)
}
