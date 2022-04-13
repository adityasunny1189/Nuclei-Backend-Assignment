package handler

import (
	"bufio"
	"fmt"
	"nuclei-assignment-2/models"
	"nuclei-assignment-2/services/src"
	"nuclei-assignment-2/utils"
)

const (
	delmsg = "enter the user's rollno to delete the user"
)

func DeleteUserDetails(in *bufio.Scanner, u []models.User, rnl map[int]bool) (usrs []models.User, rn int) {
	// read the user rollno input
	// validate the input
	// delete the user
	// return the rollno and remaining users
	usrs = u
	fmt.Printf("%s\n", delmsg)
	r := src.ReadInput("rollno", in)
	op, err := utils.ValidateRollNumberDelOp(r, rnl)
	if err != nil {
		fmt.Println(err)
		return
	}
	rn = op
	var rni int
	for i, usr := range u {
		if usr.Rollno == rn {
			rni = i
		}
	}
	usrs = append(usrs[:rni], usrs[rni+1:]...)
	return
}
