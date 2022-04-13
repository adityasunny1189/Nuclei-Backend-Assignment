package utils

import "fmt"

const (
	ExpectedStringContent = `❌expected string`
	InputNameErrContent   = `❌Input Rejected
Usage: [-name <product name>]`
	InputValueErrContent = `❌Input Rejected
--------------------------
|Usage: 
|-type <raw || imported || manufactured>
|-price <price of item>
|-quantity <no of items>
--------------------------`
	ParseErrContent = `❌error occurred in parsing input, please give correct input`
)

func StringError() {
	fmt.Println(ExpectedStringContent)
}

func InputNameError() {
	fmt.Println(InputNameErrContent)
}

func InputValueError(err error) {
	fmt.Println(err, InputValueErrContent)
}

func ParseFloatError(err error) {
	fmt.Println(err, ParseErrContent)
}

func ParseIntError(err error) {
	fmt.Println(err, ParseErrContent)
}
