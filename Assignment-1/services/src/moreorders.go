package src

import (
	"bufio"
	"fmt"
	"os"
)

func MoreOrders() bool {
	reader := bufio.NewReader(os.Stdin)
choicelabel:
	fmt.Printf("Do you want to enter details of any other item (y/n): ")
	choice, _, err := reader.ReadRune()
	if err != nil {
		fmt.Println(err)
	}
	switch choice {
	case 'y':
		return true
	case 'n':
		return false
	default:
		fmt.Printf("invalid choice\n")
		goto choicelabel
	}
}
