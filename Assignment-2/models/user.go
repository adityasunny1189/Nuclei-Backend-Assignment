package models

type User struct {
	Fullname string `json:"full name"`
	Age      int    `json:"age"`
	Address  string `json:"address"`
	Rollno   int    `json:"roll no"`
	Courses  []byte `json:"courses"`
}

func (u *User) Setter(
	name string,
	age int,
	addr string,
	rollno int,
	courses []byte) {
	u.Fullname = name
	u.Age = age
	u.Address = addr
	u.Rollno = rollno
	u.Courses = courses
}
