package test

import (
	"nuclei-assignment-2/utils"
	"testing"
)

var userstestdata = []struct {
	tid      int
	tname    string
	tage     string
	taddr    string
	trollno  string
	tcourses []byte
	expected bool
}{
	{
		tid:      1,
		tname:    "aditya",
		tage:     "21",
		taddr:    "darbhanga",
		trollno:  "3",
		tcourses: []byte{'A', 'B', 'C', 'D'},
		expected: true,
	},
}

func TestUserInput(t *testing.T) {
	for _, data := range userstestdata {
		actualValue := true
		_, nerr := utils.ValidateFullName(data.tname)
		_, aerr := utils.ValidateAge(data.tage)
		_, addrerr := utils.ValidateAddr(data.taddr)
		if nerr != nil || aerr != nil || addrerr != nil {
			actualValue = false
		}
		if data.expected != actualValue {
			t.Errorf("error in test case %d\n", data.tid)
		}
	}
}
