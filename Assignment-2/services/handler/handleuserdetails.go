package handler

import (
	"bufio"
	"fmt"
	"nuclei-assignment-2/models"
	"nuclei-assignment-2/services/src"
	"nuclei-assignment-2/utils"
)

func AddUserDetails(in *bufio.Scanner, rollnolist map[int]bool) models.User {
	var (
		name    string
		age     int
		addr    string
		rollno  int
		courses []byte
	)
	courselist := make(map[byte]bool)
inputlabel:
	for {
		if name == "" {
			op, err := utils.ValidateFullName(src.ReadInput("name", in))
			if err != nil {
				fmt.Println(err)
				continue inputlabel
			}
			name = op
		}
		if age == 0 {
			op, err := utils.ValidateAge(src.ReadInput("age", in))
			if err != nil {
				fmt.Println(err)
				continue inputlabel
			}
			age = op
		}
		if addr == "" {
			op, err := utils.ValidateAddr(src.ReadInput("address", in))
			if err != nil {
				fmt.Println(err)
				continue inputlabel
			}
			addr = op
		}
		if rollno == 0 {
			op, err := utils.ValidateRollNumber(src.ReadInput("roll no", in), rollnolist)
			if err != nil {
				fmt.Println(err)
				continue inputlabel
			}
			rollno = op
		}
	courselabel:
		for len(courses) != 4 {
			op, err := utils.ValidateCourse(src.ReadInput("course", in), courselist)
			if err != nil {
				fmt.Println(err)
				continue courselabel
			} else {
				courselist[op] = true
				courses = append(courses, op)
			}
		}
		break
	}
	return models.User{
		Fullname: name,
		Age:      age,
		Address:  addr,
		Rollno:   rollno,
		Courses:  courses,
	}
}
