package src

import (
	"bufio"
	"fmt"
)

func ReadInput(op string, in *bufio.Scanner) string {
	fmt.Printf("enter your %s: ", op)
	in.Scan()
	return in.Text()
}
