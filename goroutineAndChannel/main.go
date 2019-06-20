package main

import (
	"fmt"
	"net/http"
)

func main() {

	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	c := make(chan string)

	//goroutines
	for _, link := range links {
		go checkLink(link, c)
	}

	//receive message from
	for i := 0; i < len(links); i++ {
		fmt.Println(<-c)
	}

}
func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		c <- "message from channel: down"
		return
	}

	fmt.Println(link, "is up!")
	c <- "message from channel: up"
}
