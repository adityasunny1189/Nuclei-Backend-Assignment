package services

import (
	"bufio"
	"fmt"
)

const (
	menu = `-------- OPTIONS --------
1. Add User details.
2. Display User details.
3. Delete User details
4. Save User details.
5. Exit
---------- END ----------`
)

func DisplayMenu() {
	fmt.Printf("%s\n", menu)
}

func ReadInput(op string, in *bufio.Scanner) string {
	fmt.Printf("enter your %s: ", op)
	in.Scan()
	return in.Text()
}
