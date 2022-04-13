package src

import "fmt"

const (
	InputDetailsContent = `--------------------------
|Usage: 
|-type <raw || imported || manufactured>
|-price <price of item>
|-quantity <no of items>
--------------------------
`
	InputNameDetailsContent = `--------------------------------
Usage: [-name <product name>]
--------------------------------
`
)

func PrintInputDetails() {
	fmt.Print(InputDetailsContent)

}

func PrintInputNameDetails() {
	fmt.Print(InputNameDetailsContent)
}
