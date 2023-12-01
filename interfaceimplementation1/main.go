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
	Shareholding string
}

type AXISBANK struct {
	Shareholding string
}

func (m MAHABANK) mutualFund() {
	fmt.Println("The details of the user are :->")
	fmt.Println("Shareholding in MAHABANK:", m.Shareholding)
}

func (a AXISBANK) mutualFund() {
	fmt.Println("The details of the user are :->")
	fmt.Println("Shareholding in AXISBANK:", a.Shareholding)
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

// YASH DWIVEDI will always be in progress and will live long
