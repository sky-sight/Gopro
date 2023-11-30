package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Kindly Rate")

	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n') // Reading until newline (\n)

	fmt.Println("Thanks", strings.TrimSpace(input))

	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println(err)
	} else {
		newRating := numRating + 1
		fmt.Println("New Rating:", newRating)
	}
}
