package main

import (
	"fmt"
)

type fund struct {
	mutualFund int
	stock      int
	Baskets    int
}

func main() {
	// Create an empty map with keys of type string and values of type fund
	user1 := make(map[string]fund)

	// Assign fund values for the user "MD"
	user1["MD"] = fund{
		mutualFund: 10,
		stock:      100,
		Baskets:    1000,
	}

	fmt.Println(user1)
}
