package main

import (
	"fmt"
	"net/http"
)

func main() {
	//links example
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

	//buffered channel
	var queue = make(chan string, 2)

	//send links as tasks to buffered channel
	go sendLinks(links, queue)

	//once can read from chanel then generate a new goroutine
	dispatch(queue)

}
func sendLinks(l []string, c chan string) {
	for _, link := range l {
		c <- link
	}
}

//Dispatch. Generate worker goroutine to work
func dispatch(c chan string) {
	for {
		select {
		case link := <-c:
			go worker(link)
		default:
		}
	}
}

func worker(l string) {
	checkLink(l)
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
