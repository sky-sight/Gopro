package main

import (
	"fmt"
)

type numbers interface {
	add(float64, float64) float64
	sub(float64, float64) float64
	div(float64, float64) (float64, error)
	mul(float64, float64) float64
}

func main() {
	var num1, num2 float64

	fmt.Println("Enter Two Numbers:-")

	fmt.Scanln()
	_, err1 := fmt.Scanln(&num1)
	if err1 != nil {
		fmt.Println("Error in input 1:", err1)
		return
	}
	fmt.Println("The First number is:", num1)

	_, err2 := fmt.Scanln(&num2)
	if err2 != nil {
		fmt.Println("Error in input 2:", err2)
		return
	}
	fmt.Println("The Second number is:", num2)

	basicCalc := BasicCalculator{} // Create an instance of BasicCalculator

	result := addition(basicCalc, num1, num2)

	mult := substraction(basicCalc, num1, num2)

	fmt.Println("Addition Result:", result)
	fmt.Println("Multiplication Result:", mult)
}

type BasicCalculator struct {
}

func (bc BasicCalculator) add(a, b float64) float64 {
	return a + b
}

func (bc BasicCalculator) sub(a, b float64) float64 {
	return a - b
}

func (bc BasicCalculator) mul(a, b float64) float64 {
	return a * b
}

func (bc BasicCalculator) div(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("Zero Division error")
	}
	return a / b, nil
}

func addition(n numbers, a, b float64) float64 {
	//fmt.Println(n.add(a, b))
	return n.add(a, b)
}

func substraction(n numbers, a, b float64) float64 {
	return n.mul(a, b)
}
