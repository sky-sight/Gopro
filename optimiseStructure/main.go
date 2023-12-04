package main

import "fmt"

type StockMarket struct {
	Name       string
	Age        int
	ProfitLoss int
}

func (s StockMarket) Println() {
	fmt.Println(s)
}

func (s StockMarket) Travesting() {
	fmt.Printf("user %s is Travesting\n", s.Name)
}

func (s StockMarket) GetName() string {
	return s.Name
}

func main() {
	s := StockMarket{
		Name:       "Manish",
		Age:        24,
		ProfitLoss: 100,
	}

	s.Println()
	s.Travesting()
}

/*package main

import "fmt"

func main() {

	stock := 2

	mutualfund := &stock

	fmt.Println(mutualfund)

}*/
