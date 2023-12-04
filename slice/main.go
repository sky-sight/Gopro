/*package main

import (
	"fmt"
)

func main() {
	cards :=[]string{"MutualFund", Greeting(), "Finance"}
	fmt.Println(cards)
}

func Greeting() string {
	return "Hello, Manish!!"
}*/

package main

import (
	"fmt"
)

func main() {
	cards := []string{"MutualFund"}
	cards = append(cards, Greeting())
	cards = append(cards, "Finance")

	fmt.Println(cards)
}

func Greeting() string {
	return "Hello, Manish!!"
}

/*package main

import (
	"fmt"
)

func main() {
	// Initialize an empty slice of strings
	var cards []string

	// Append string literals and the result of the Greeting function
	cards = append(cards, "MutualFund")
	cards = append(cards, Greeting())
	cards = append(cards, "Finance")

	fmt.Println(cards)

	for i, card := range cards {

		fmt.Println(i, card)
	}
}

func Greeting() string {
	return "Hello, Manish!!"
}*/
