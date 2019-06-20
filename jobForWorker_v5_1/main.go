package main

import (
	"fmt"
	"net/http"
	"runtime"
	"strconv"
	"time"
)

//JobQueue and WorkerPool
var MaxWorker = 10
var WorkerPool = make(chan chan string, MaxWorker)
var JobQueue = make(chan string, 100)

type Worker struct {
	id         int
	WorkerPool chan chan string
	JobChannel chan string
}

func main() {
	//CPU num setting
	fmt.Println(runtime.NumCPU())
	result := runtime.GOMAXPROCS(4)
	if result < 1 {
		fmt.Println("Setting failed")
	} else {
		fmt.Println("Setting success")
	}
	time.Sleep(time.Second * 5)

	//links example. each link is a job
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}
	//add more fake links
	var numOfLink = 1000
	for i := 0; i < numOfLink; i++ {
		fakeLink := fmt.Sprintf("http://fake%d.com", i)
		links = append(links, fakeLink)
	}

	//send links to JobQueue
	go sendLink(links)

	//dispatch work
	dispatch()
}

//keeping send links to JobQueue
func sendLink(l []string) {
	for _, link := range l {
		JobQueue <- link
	}
}

func dispatch() {
	//create max number of worker
	for i := 0; i < MaxWorker; i++ {
		worker := NewWorker(WorkerPool, i)
		worker.Start()
	}
	//only dispacht the link to worker
	for {
		select {
		case link := <-JobQueue:
			// a job request has been received
			go func(l string) {
				// try to obtain a worker job channel that is available.
				// this will block until a worker is idle
				jobChannel := <-WorkerPool
				// dispatch the job to the worker job channel
				jobChannel <- l
			}(link)
		}
	}
}

//Check link.
func checkLink(link string) string {
	var result string
	_, err := http.Get(link)
	if err != nil {
		result = link + "might be down!"
	}

	result = link + "is up!"
	return result
}

//Worker
func NewWorker(workerPool chan chan string, i int) Worker {
	return Worker{
		id:         i,
		WorkerPool: workerPool,
		JobChannel: make(chan string)}
}

func (w Worker) Start() {
	go func() {
		for {
			// register the current worker into the worker queue.
			w.WorkerPool <- w.JobChannel
			select {
			case link := <-w.JobChannel:
				//do job
				fmt.Println("Worker " + strconv.Itoa(w.id) + ": " + checkLink(link))
			default:
			}
		}
	}()
}
