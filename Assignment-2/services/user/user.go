package user

type User struct {
	Fullname string   `json:"full name"`
	Age      int      `json:"age"`
	Address  string   `json:"address"`
	Rollno   int      `json:"roll no"`
	Courses  []string `json:"courses"`
}
