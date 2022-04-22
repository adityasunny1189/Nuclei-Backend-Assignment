package test

import (
	pkg "nuclei-assignment-2/services/user/pkg"
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
		tage:     "ui",
		taddr:    "darbhanga",
		trollno:  "3",
		tcourses: []byte{'A', 'B', 'C', 'D'},
		expected: false,
	},
	{
		tid:      2,
		tname:    "aditya",
		tage:     "21",
		taddr:    "darbhanga",
		trollno:  "3",
		tcourses: []byte{'A', 'B', 'C', 'D'},
		expected: true,
	},
	{
		tid:      3,
		tname:    "aditya",
		tage:     "-54",
		taddr:    "darbhanga",
		trollno:  "3",
		tcourses: []byte{'A', 'B', 'C', 'D'},
		expected: false,
	},
	{
		tid:      4,
		tname:    "aditya",
		tage:     "234",
		taddr:    "darbhanga",
		trollno:  "3",
		tcourses: []byte{'A', 'B', 'C', 'D'},
		expected: false,
	},
	{
		tid:      5,
		tname:    "kunal",
		tage:     "21",
		taddr:    "patna",
		trollno:  "3",
		tcourses: []byte{'A', 'B', 'C', 'D'},
		expected: true,
	},
	{
		tid:      6,
		tname:    "",
		tage:     "2",
		taddr:    "darbhanga",
		trollno:  "3",
		tcourses: []byte{'A', 'B', 'C', 'D'},
		expected: false,
	},
	{
		tid:      7,
		tname:    "amit",
		tage:     "63",
		taddr:    "madhya pradesh",
		trollno:  "3",
		tcourses: []byte{'A', 'B', 'C', 'D'},
		expected: true,
	},
	{
		tid:      8,
		tname:    "aditya",
		tage:     "ui",
		taddr:    "",
		trollno:  "3",
		tcourses: []byte{'A', 'B', 'C', 'D'},
		expected: false,
	},
	{
		tid:      9,
		tname:    "aditya",
		tage:     "ui",
		taddr:    "34",
		trollno:  "3",
		tcourses: []byte{'A', 'B', 'C', 'D'},
		expected: false,
	},
	{
		tid:      10,
		tname:    "12",
		tage:     "ui",
		taddr:    "darbhanga",
		trollno:  "3",
		tcourses: []byte{'A', 'B', 'C', 'D'},
		expected: false,
	},
}

func TestUserInput(t *testing.T) {
	for _, data := range userstestdata {
		actualValue := true
		_, nerr := pkg.ValidateFullName(data.tname)
		_, aerr := pkg.ValidateAge(data.tage)
		_, addrerr := pkg.ValidateAddr(data.taddr)
		if nerr != nil || aerr != nil || addrerr != nil {
			actualValue = false
		}
		if data.expected != actualValue {
			t.Errorf("error in test case %d\t expected %v got %v\n", data.tid, data.expected, actualValue)
		}
	}
}
