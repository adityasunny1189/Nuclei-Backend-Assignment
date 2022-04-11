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

func StartScript() {
	var users []models.User
	choice := "Enter your choice: "
	src.DisplayMenu()
	in := bufio.NewScanner(os.Stdin)
	fmt.Printf("%s", choice)
start:
	for in.Scan() {
		ch, ok := utils.ValidateInputChoice(in.Text())
		if !ok {
			fmt.Printf("%s", choice)
			continue start
		}
		switch ch {
		case 1:
			u := handler.AddUserDetails(in)
			users = append(users, u)
			fmt.Println(users)
		case 2:
		case 3:
		case 4:
		case 5:
			break start
		}
		fmt.Printf("%s", choice)
	}
}
