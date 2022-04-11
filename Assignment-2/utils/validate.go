package utils

import (
	"fmt"
	"strconv"
)

// validate name, should not be (blank/int)
func ValidateFullName(n string) (name string, ok bool) {
	_, err := strconv.Atoi(n)
	if err == nil {
		fmt.Printf("name can't be a number\n")
	} else if n == "" {
		fmt.Printf("name can't be empty\n")
	} else {
		name, ok = n, true
	}
	return
}

// validate age, handle (blank, string, <150)
func ValidateAge(a string) (age int, ok bool) {
	ag, err := strconv.Atoi(a)
	if err != nil {
		fmt.Println(err, "invalid input")
	} else if ag <= 0 {
		fmt.Printf("age can't be 0 or negative\ninvalid input\n")
	} else if ag > 150 {
		fmt.Printf("age can't be that big\ninvalid input\n")
	} else {
		age, ok = ag, true
	}
	return
}

// validate address, should not be (int, blank)
func ValidateAddr(a string) (addr string, ok bool) {
	_, err := strconv.Atoi(a)
	if err == nil {
		fmt.Printf("address can't be integer\ninvalid input\n")
	} else if a == "" {
		fmt.Printf("please provide an address\ninvalid input\n")
	} else {
		addr, ok = a, true
	}
	return
}

// validate roll no, handle (blank, string, unique)
// Todo: handle unique roll numbers
func ValidateRollNumber(r string) (roll int, ok bool) {
	rl, err := strconv.Atoi(r)
	if err != nil {
		fmt.Println(err, "invalid input")
	} else if rl <= 0 {
		fmt.Printf("roll no can't be 0 or negative\ninvalid input\n")
	} else {
		roll, ok = rl, true
	}
	return
}

// validate courses, handle (blank, string, int, len(courses) == 4)
func ValidateCourses(c string) (byte, bool) {
	return 'A', true
}

// validate choice, should not be (string, blank, <= 5)
func ValidateInputChoice(c string) (ch int, ok bool) {
	op, err := strconv.Atoi(c)
	if err != nil {
		fmt.Println(err, "invalid choice")
	} else if op > 5 {
		fmt.Printf("choice can't be greater than 5\ninvalid choice\n")
	} else if op <= 0 {
		fmt.Printf("choice can't be negative or 0\ninvalid choice\n")
	} else {
		ch, ok = op, true
	}
	return
}
