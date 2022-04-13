package services

import (
	"bufio"
	"fmt"
	"nuclei-assignment-2/models"
	"nuclei-assignment-2/services/handler"
	"nuclei-assignment-2/services/src"
	"nuclei-assignment-2/utils"
	"os"
)

const (
	choice = "Enter your choice: "
)

func StartScript() {
	var users []models.User
	rollnolist := make(map[int]bool)
	src.DisplayMenu()
	in := bufio.NewScanner(os.Stdin)
	fmt.Printf("%s", choice)
start:
	for in.Scan() {
		ch, err := utils.ValidateInputChoice(in.Text())
		if err != nil {
			fmt.Println(err)
			fmt.Printf("%s", choice)
			continue start
		}
		switch ch {
		case 1: // Add user details
			u := handler.AddUserDetails(in, rollnolist)
			rollnolist[u.Rollno] = true
			users = append(users, u)
			fmt.Println(users)
		case 2: // Show user details
		case 3: // Delete user details
		case 4: // Save user details
		case 5: // Exit
			break start
		}
		src.DisplayMenu()
		fmt.Printf("%s", choice)
	}
}
