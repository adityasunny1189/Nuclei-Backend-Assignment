package services

import "nuclei-assignment-2/models"

func FillRollNo(users []models.User) map[int]bool {
	rnl := make(map[int]bool)
	for _, user := range users {
		rnl[user.Rollno] = true
	}
	return rnl
}
