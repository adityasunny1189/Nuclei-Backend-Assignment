package utils

import "fmt"

func StringError() {
	fmt.Printf("❌expected string\n")
}

func InputNameError() {
	fmt.Printf("❌Input Rejected\n")
	fmt.Printf("Usage:\t[-name <product name>]\n")
}

func InputValueError() {
	fmt.Printf("❌Input Rejected\n")
	fmt.Printf("--------------------------\n")
	fmt.Printf(`|Usage: 
|-type <raw || imported || manufactured>
|-price <price of item>
|-quantity <no of items>
`)
	fmt.Printf("--------------------------\n")
}

func ParseFloatError(err error) {
	fmt.Println(err, "❌error occured in parsing input, please give correct input\n")
}

func ParseIntError(err error) {
	fmt.Println(err, "❌error occured in parsing input, please give correct input\n")
}
