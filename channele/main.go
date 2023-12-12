package main

import (
	"fmt"
	"net/http"
)

func main() {

	links := []string{
		"http://www.facebook.com",
		"http://www.amazon.com",
		"http://www.google.com",
		"http://www.stackoverflow.com",
	}

	ch := make(chan string)

	for _, link := range links {
		fmt.Println(link)
		go checklist(link, ch)
	}

	for i := 0; i < len(links); i++ {

		fmt.Println(<-ch)

	}

	//fmt.Println(<-ch)
}

func checklist(link string, channel chan string) {
	_, err := http.Get(link)

	if err != nil {
		fmt.Println(link, "is not working")

		channel <- "Response is Down in Particular Request"
		return
	}

	fmt.Println(link, "is working")

	channel <- "It's ,working fine"

	//c <-
}
