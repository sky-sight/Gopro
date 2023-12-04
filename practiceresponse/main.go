package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	res, err := http.Get("https://chat.openai.com")
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	defer res.Body.Close()

	fmt.Println("Response Status:", res.Status)

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("error reading body:", err)
		return
	}

	fmt.Println("Response Body:", string(body))

}
