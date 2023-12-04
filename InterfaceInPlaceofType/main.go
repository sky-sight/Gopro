/*package main

func Anything(anything interface{}) {

}

func man() {

	Anything(struct{}{})
	Anything(3.148)

	m := make(map[string]interface{})

	m["Manish"] = "ReligareBroking"

}*/

package main

import "fmt"

func Anything(anything interface{}) {

}

func main() {
	Anything(struct{}{})
	Anything(3.148)

	m := make(map[string]interface{})
	m["Manish"] = "ReligareBroking"

	// Print the map value to showcase its content
	fmt.Println(m["Manish"])
}
