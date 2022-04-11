package src

import "fmt"

func DisplayMenu() {
	menu := `-------- OPTIONS --------
1. Add User details.
2. Display User details.
3. Delete User details
4. Save User details.
5. Exit
---------- END ----------`

	fmt.Printf("%s\n", menu)
}
