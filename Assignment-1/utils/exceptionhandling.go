package utils

import "fmt"

func StringError() {
	fmt.Printf("❌expected string\n")
}

func InputNameError() {
	fmt.Printf("❌Input Rejected\nUsage:\t[-name <product name>]\n")
}

func InputValueError() {
	fmt.Printf(`❌Input Rejected
--------------------------
|Usage: 
|-type <raw || imported || manufactured>
|-price <price of item>
|-quantity <no of items>
--------------------------
`)
}

func ParseFloatError(err error) {
	fmt.Println(err, "❌error occured in parsing input, please give correct input")
}

func ParseIntError(err error) {
	fmt.Println(err, "❌error occured in parsing input, please give correct input")
}
