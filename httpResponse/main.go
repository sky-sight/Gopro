package main

import (
	"fmt"
	"net/http"
)

const url = "https://pkg.go.dev/net/http"

func main() {
	fmt.Println("Web Request is coming:")

	response, err := http.Get(url)

	if err != nil {
		panic(err)
	} else {
		defer response.Body.Close()

		fmt.Printf("Response is of type: %T\n", response)
	}
}
