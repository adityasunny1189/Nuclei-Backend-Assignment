package handler

import (
	"bufio"
	"nuclei-assignment-2/models"
	"nuclei-assignment-2/services/src"
	"nuclei-assignment-2/utils"
)

func AddUserDetails(in *bufio.Scanner) models.User {
	var u models.User
	var (
		name   string
		age    int
		addr   string
		rollno int
		// courses []byte
		ok bool
	)
inputlabel:
	for {
		if name == "" {
			name, ok = utils.ValidateFullName(src.ReadInput("name", in))
			if !ok {
				continue inputlabel
			} else {
				u.Fullname = name
			}
		}
		if age == 0 {
			age, ok = utils.ValidateAge(src.ReadInput("age", in))
			if !ok {
				continue inputlabel
			} else {
				u.Age = age
			}
		}
		if addr == "" {
			addr, ok = utils.ValidateAddr(src.ReadInput("address", in))
			if !ok {
				continue inputlabel
			} else {
				u.Address = addr
			}
		}
		if rollno == 0 {
			rollno, ok = utils.ValidateRollNumber(src.ReadInput("roll no", in))
			if !ok {
				continue inputlabel
			} else {
				u.Rollno = rollno
			}
		}
		break
	}
	// courses := src.ReadInput("courses", in)
	return u
}
