package utils

import (
	"errors"
	"strconv"
)

const (
	GreaterChoiceErr   = "choice can't be greater than 5"
	NegativeChoiceErr  = "choice can't be negative or 0"
	NameEmptyErr       = "name can't be empty"
	NameIntErr         = "name can't be integer"
	NegativeAgeErr     = "age can't be negative"
	BigAgeErr          = "age too big to be real"
	AddrIntErr         = "address can't be integer value"
	EmptyAddrErr       = "address can't be empty"
	NegativeRollnoErr  = "negative roll no is not allowed"
	DuplicateRollnoErr = "roll no already in use"
	CourseIntErr       = "course can't be integer"
	EmptyCourseErr     = "course can't be left empty"
	InvalidCourseErr   = "course diffrent then prescribed"
	DuplicateCourseErr = "course already subscribed"
)

// validate name, should not be (blank/int)
func ValidateFullName(n string) (name string, err error) {
	_, er := strconv.Atoi(n)
	if er == nil {
		err = errors.New(NameIntErr)
	} else if n == "" {
		err = errors.New(NameEmptyErr)
	} else {
		name = n
	}
	return
}

// validate age, handle (blank, string, <150)
func ValidateAge(a string) (age int, err error) {
	ag, er := strconv.Atoi(a)
	if er != nil {
		err = er
	} else if ag <= 0 {
		err = errors.New(NegativeAgeErr)
	} else if ag > 150 {
		err = errors.New(BigAgeErr)
	} else {
		age = ag
	}
	return
}

// validate address, should not be (int, blank)
func ValidateAddr(a string) (addr string, err error) {
	_, er := strconv.Atoi(a)
	if er == nil {
		err = errors.New(AddrIntErr)
	} else if a == "" {
		err = errors.New(EmptyAddrErr)
	} else {
		addr = a
	}
	return
}

// validate roll no, handle (blank, string, unique)
func ValidateRollNumber(r string, rnlist map[int]bool) (roll int, err error) {
	rl, er := strconv.Atoi(r)
	if er != nil {
		err = er
	} else if rnlist[rl] {
		err = errors.New(DuplicateRollnoErr)
	} else if rl <= 0 {
		err = errors.New(NegativeRollnoErr)
	} else {
		roll = rl
	}
	return
}

// validate courses, handle (blank, string, int, len(courses) == 4)
func ValidateCourse(c string, cl map[byte]bool) (course byte, err error) {
	_, er := strconv.Atoi(c)
	if er == nil {
		err = errors.New(CourseIntErr)
	} else if c == "" {
		err = errors.New(EmptyCourseErr)
	} else if len(c) != 1 {
		err = errors.New(InvalidCourseErr)
	} else {
		byteArr := []byte(c)
		if cl[byteArr[0]] {
			err = errors.New(DuplicateCourseErr)
			return
		}
		switch byteArr[0] {
		case 'A', 'B', 'C', 'D', 'E', 'F':
			course = byteArr[0]
		default:
			err = errors.New(InvalidCourseErr)
		}
	}
	return
}

// validate choice, should not be (string, blank, <= 5)
func ValidateInputChoice(c string) (ch int, err error) {
	op, er := strconv.Atoi(c)
	if er != nil {
		err = er
	} else if op > 5 {
		err = errors.New(GreaterChoiceErr)
	} else if op <= 0 {
		err = errors.New(NegativeChoiceErr)
	} else {
		ch = op
	}
	return
}
