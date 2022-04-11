package src

import "fmt"

func InputDetails() {
	fmt.Printf("--------------------------\n")
	fmt.Printf(`|Usage: 
|-type <raw || imported || manufactured>
|-price <price of item>
|-quantity <no of items>
`)
	fmt.Printf("--------------------------\n")
}

func InputNameDetails() {
	fmt.Printf("--------------------------------\n")
	fmt.Printf("Usage:\t[-name <product name>]\n")
	fmt.Printf("--------------------------------\n")
}
