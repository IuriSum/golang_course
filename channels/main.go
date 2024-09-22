package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://go.dev",
		"http://amazon.com",
	}
	// channel
	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	// Channel blocker
	// this "for" format makes visible that channels emit values then iterate
	for l := range c {
		// forces the Main Routine to pause and execute routine
		// time.Sleep(time.Duration(2000 * time.Millisecond))

		// forces the channel to maintain main function running
		// <-c makes the channel give string data to function
		// go checkLink(l, c)

		// literal function
		go (func(link string) {
			time.Sleep(5000 * time.Millisecond)
			checkLink(link, c)
		}(l) )// l is a memory copy of the string that is passing trouch channel
	}
}

func checkLink(link string, c chan string) {

	// the response will not be used
	_, err := http.Get(link)

	if err != nil {
		fmt.Println("Error in ", link)
		fmt.Println(err)

		// signals to channel a value
		c <- link

		return
	}

	fmt.Println(link, "Status is ok")
	c <- link
}
