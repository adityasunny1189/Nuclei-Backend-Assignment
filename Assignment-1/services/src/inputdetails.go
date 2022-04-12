package src

import "fmt"

func PrintInputDetails() {
	fmt.Printf(`--------------------------
|Usage: 
|-type <raw || imported || manufactured>
|-price <price of item>
|-quantity <no of items>
--------------------------
`)

}

func PrintInputNameDetails() {
	fmt.Printf(`--------------------------------
Usage: [-name <product name>]
--------------------------------
`)
}
