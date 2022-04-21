package services

import "nuclei-assignment-2/services/user"

func fillRollNo(users []user.User) map[int]bool {
	rnl := make(map[int]bool)
	for _, user := range users {
		rnl[user.Rollno] = true
	}
	return rnl
}

func LoadApp() (users []user.User, rollnolist map[int]bool) {
	users = user.DeserializeJSON()
	rollnolist = fillRollNo(users)
	return
}
