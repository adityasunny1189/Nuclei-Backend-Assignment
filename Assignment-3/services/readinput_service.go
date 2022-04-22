package services

import (
	"bufio"
	"fmt"
)

func ReadInput(str string, in *bufio.Scanner) string {
	fmt.Printf("%s", str)
	in.Scan()
	return in.Text()
}
