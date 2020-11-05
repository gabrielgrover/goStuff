package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://goggle.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	for l := range c {
		go func(link string) {
			time.Sleep(time.Second)
			checkLink(link, c)
		}(l)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		c <- "Might be down"
		return
	}

	fmt.Println(link, "is up!")
	// msg := fmt.Sprintf("%s is up!", link)
	// c <- "it's up"
	c <- link
}
