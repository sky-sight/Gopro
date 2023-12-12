//sudo dpkg -i docker-desktop-2.5.0-x86_64.deb
package main

import (
	"fmt"
)

type equity interface {
	mutualFund()
	stock()
}

type MAHABANK struct {
	Promoter_Name string
}

type AXISBANK struct {
	Promoter_Name string
}

func (m MAHABANK) mutualFund() {
	fmt.Println("The details of the user are :->")
	fmt.Println("Majority Stake Promoter Name in MAHABANK:", m.Promoter_Name)
}

func (a AXISBANK) mutualFund() {
	fmt.Println("The details of the user are :->")
	fmt.Println("Shareholding in AXISBANK:", a.Promoter_Name)
}

func (m MAHABANK) stock() {
	fmt.Println("MAHABANK stock details")
}

func (a AXISBANK) stock() {
	fmt.Println("AXISBANK stock details")
}

func main() {
	User1 := MAHABANK{"MD"}
	User2 := AXISBANK{"MANISH DWIVEDI"}

	var e equity
	e = User1
	e.mutualFund()
	e.stock()

	e = User2
	e.mutualFund()
	e.stock()

}


//need to understand this