package handler

import (
	"fmt"
	"nuclei-assignment-2/models"
)

const (
	startmenu = `-----------------------------------------------------------------------------------------------------------------------
%-20s %-20s %-20s %-20s %20s
-----------------------------------------------------------------------------------------------------------------------
`
	endmenu = `-----------------------------------------------------------------------------------------------------------------------`
)

func ParseCourse(c []byte) (crs string) {
	var cl []byte
	for _, v := range c {
		cl = append(cl, v)
		cl = append(cl, ',')
	}
	crs = string(cl)
	return
}

func ShowUserDetails(users []models.User) {
	fmt.Printf(startmenu, "name", "rollno", "age", "address", "courses")
	for _, user := range users {
		courses := ParseCourse(user.Courses)
		fmt.Printf("%-20s %-20d %-20d %-20s %20v\n",
			user.Fullname, user.Rollno, user.Age, user.Address, courses)
	}
	fmt.Println(endmenu)
}
