package main

import (
	"fmt"
	"net/http"
	"time"
)

type checkCache []string

func (cc checkCache) print() {
	fmt.Println("Confirming cache...")
	for _, url := range cc {
		fmt.Println(url)
	}
}

func httpRequest(url string, c chan string) {

	startTimer := time.Now()

	res, err := http.Get(url)
	if err != nil {
		fmt.Println("[ERROR]:", err)
		c <- fmt.Sprintf("[ERROR]: %v", err)
	}
	endTimer := time.Since(startTimer)

	fmt.Println("[INFO]: Url:", url, "Status:", res.Status, "Time:", endTimer)

	c <- url
}

func main() {
	fmt.Println("Checking status...")
	cache := checkCache{
		"https://facebook.com",
		"https://amazon.com",
		"https://apple.com",
		"https://netflix.com",
		"https://google.com",
	}

	//Make channel
	c := make(chan string)

	//Make concurrent http request
	for _, url := range cache {
		go httpRequest(url, c)
	}

	//Print requests from channel equal to length of requests
	for u := range c {
		go func(u string) {
			time.Sleep(5 * time.Second)
			httpRequest(u, c)
		}(u)
	}
}
