package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://golang.org",
		"http://amazon.com",
		"http://google.com",
		"http://facebook.com",
	}

	c := make(chan string)

	for _, link := range links {
		// checkLink(link)
		go checkLink(c, link)
		//fmt.Println(<-c)
	}
	//fmt.Println(<-c)
	//fmt.Println(<-c)
	//fmt.Println(<-c)
	//fmt.Println(<-c)
	for l := range c {
		//fmt.Println(<-c)
		//time.Sleep(5 * time.Second)
		//go checkLink(c, l) //for repeated checking
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(c, link)
		}(l)
	}
	//fmt.Println(<-c)
}

func checkLink(channel chan string, link string) {
	//time.Sleep(5 * time.Second)
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "is down")
		channel <- link
		return
	}
	fmt.Println(link, "is up")
	channel <- link
}
