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

	// ReadString('\n') includes the newline character in the input,
	// causing issues in conversion. Using ReadString('\n') instead.
	input, _ := reader.ReadString('\n')

	// Trim the newline character and spaces
	input = strings.TrimSpace(input)

	fmt.Println("Thanks", input)

	// Check if input is empty after trimming
	if input == "" {
		fmt.Println("No rating provided")
		return
	}

	numRating, err := strconv.ParseFloat(input, 64)

	if err != nil {
		fmt.Println("Invalid input:", err)
	} else {
		newRating := numRating + 1
		fmt.Println("New Rating:", newRating)
	}
}
