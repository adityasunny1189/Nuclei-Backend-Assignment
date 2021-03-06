package app

import (
	"bufio"
	"fmt"
	"nuclei-assignment-2/services"
	"nuclei-assignment-2/services/user"
	pkg "nuclei-assignment-2/services/user/pkg"
	"os"
)

const (
	choice = "Enter your choice: "
)

func StartApp() {
	users, rollnolist := services.LoadApp()
	pkg.DisplayMenu()
	in := bufio.NewScanner(os.Stdin)
	fmt.Printf("%s", choice)
start:
	for in.Scan() {
		ch, err := pkg.ValidateInputChoice(in.Text())
		if err != nil {
			fmt.Println(err)
			fmt.Printf("%s", choice)
			continue start
		}
		var u user.User
		var db user.DBMethods = &u
		switch ch {
		case 1: // Add user details
			u.AddUserDetails(in, rollnolist, &users)
		case 2: // Show user details
			user.ShowDetails(db, users, in) // interface
		case 3: // Delete user details
			u.DeleteUserDetails(in, &users, rollnolist)
		case 4: // Save user details
			user.SaveDetails(db, users) // interface
		case 5: // Exit
			break start
		}
		pkg.DisplayMenu()
		fmt.Printf("%s", choice)
	}
}
