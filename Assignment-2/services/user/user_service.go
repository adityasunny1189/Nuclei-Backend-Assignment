package user

import (
	"bufio"
	"fmt"
	pkg "nuclei-assignment-2/services/user/pkg"
)

const (
	delmsg    = "enter the user's rollno to delete the user"
	startmenu = `-----------------------------------------------------------------------------------------------------------------------
%-20s %-20s %-20s %-20s %20s
-----------------------------------------------------------------------------------------------------------------------
`
	endmenu = `-----------------------------------------------------------------------------------------------------------------------`
	sortmsg = `enter the method you want to sort the data:
1. By Name
2. By Age
3. By Roll no
4. By Address
c`
)

func (u *User) AddUserDetails(in *bufio.Scanner, rollnolist map[int]bool, users *[]User) {
	var (
		name    string
		age     int
		addr    string
		rollno  int
		courses []string
	)
	courselist := make(map[string]bool)
inputlabel:
	for {
		if name == "" {
			op, err := pkg.ValidateFullName(pkg.ReadInput("name", in))
			if err != nil {
				fmt.Println(err)
				continue inputlabel
			}
			name = op
		}
		if age == 0 {
			op, err := pkg.ValidateAge(pkg.ReadInput("age", in))
			if err != nil {
				fmt.Println(err)
				continue inputlabel
			}
			age = op
		}
		if addr == "" {
			op, err := pkg.ValidateAddr(pkg.ReadInput("address", in))
			if err != nil {
				fmt.Println(err)
				continue inputlabel
			}
			addr = op
		}
		if rollno == 0 {
			op, err := pkg.ValidateRollNumber(pkg.ReadInput("roll no", in), rollnolist)
			if err != nil {
				fmt.Println(err)
				continue inputlabel
			}
			rollno = op
		}
	courselabel:
		for len(courses) != 4 {
			op, err := pkg.ValidateCourse(pkg.ReadInput("course", in), courselist)
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
	u.Fullname = name
	u.Age = age
	u.Address = addr
	u.Rollno = rollno
	u.Courses = courses
	rollnolist[u.Rollno] = true
	*users = append(*users, *u)
}

func (u *User) DeleteUserDetails(in *bufio.Scanner, users *[]User, rnl map[int]bool) {
	// read the user rollno input
	// validate the input
	// delete the user
	// return the rollno and remaining users
	fmt.Printf("%s\n", delmsg)
	r := pkg.ReadInput("rollno", in)
	op, err := pkg.ValidateRollNumberDelOp(r, rnl)
	if err != nil {
		fmt.Println(err)
		return
	}
	rn := op
	var rni int
	for i, usr := range *users {
		if usr.Rollno == rn {
			rni = i
		}
	}
	rnl[rn] = false
	*users = append((*users)[:rni], (*users)[rni+1:]...)
}

func ParseCourse(c []string) (crs string) {
	for _, val := range c {
		crs += val + " "
	}
	return
}

func (u *User) ShowUserDetails(users []User, in *bufio.Scanner) {
	fmt.Printf("%s", sortmsg)
	so := pkg.ReadInput("choice", in)
	switch so {
	case "1":
		SortDataByName(users)
	case "2":
		SortDatabyAge(users)
	case "3":
		SortDatabyRollno(users)
	case "4":
		SortDatabyAddress(users)
	default:
		fmt.Println("Invalid Choice")
	}
	fmt.Printf(startmenu, "name", "rollno", "age", "address", "courses")
	for _, user := range users {
		courses := ParseCourse(user.Courses)
		fmt.Printf("%-20s %-20d %-20d %-20s %20v\n",
			user.Fullname, user.Rollno, user.Age, user.Address, courses)
	}
	fmt.Println(endmenu)
}

func (u *User) SaveUserDetails(users []User) {
	SerializeJSON(users)
}
