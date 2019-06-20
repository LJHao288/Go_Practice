package main

import (
	"fmt"
	"net/http"
)

//JobQueue and WorkerPool
var WorkerPool = make(chan int, 10)
var JobQueue = make(chan string, 100)

func main() {

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

	//dispatch
	dispatch()
}

//keeping send links to JobQueue
func sendLink(l []string) {
	for _, link := range l {
		JobQueue <- link
	}
}

//Dispatch. Generate worker to work
func dispatch() {
	for {
		select {
		case WorkerPool <- 1:
			for {
				select {
				case link := <-JobQueue:
					WorkerPool <- 1
					go worker(link)
				default:
				}
			}
		default:
		}
	}
}

//Worker
func worker(link string) {
	defer func() { <-WorkerPool }()
	checkLink(link)
}

//Check link.
func checkLink(link string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		return
	}
	fmt.Println(link, "is up!")
}
