package user

import "bufio"

type DBMethods interface {
	ShowUserDetails([]User, *bufio.Scanner)
	SaveUserDetails([]User)
}

func ShowDetails(db DBMethods, u []User, in *bufio.Scanner) {
	db.ShowUserDetails(u, in)
}

func SaveDetails(db DBMethods, u []User) {
	db.SaveUserDetails(u)
}
