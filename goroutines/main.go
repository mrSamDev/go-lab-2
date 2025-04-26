package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {

	websites := []string{
		"https://www.hotstar.com/",
		"https://gobyexample.com/",
		"https://github.com/",
		"https://go.dev/doc/",
		"https://www.google.com/",
		"https://www.sijosam.in",
	}

	c := make(chan string, len(websites))

	for _, website := range websites {

		go fetch(website, c)

	}

	for l := range c {

		go func() {
			fmt.Println("---------------")
			time.Sleep(time.Second * 10)
			fetch(l, c)
		}()

	}

}

func fetch(url string, c chan string) {

	resp, error := http.Get(url)

	if error != nil {
		fmt.Println(url, " is Down")
		c <- url
		return
	}

	if resp.StatusCode != http.StatusOK {
		fmt.Println(url, " is Down")
		c <- url
		return
	}

	fmt.Println(url, " is Up")
	c <- url
	return

}
