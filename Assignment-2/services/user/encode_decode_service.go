package user

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"sort"
)

const (
	path = "../Assignment-2/data/users_data.json"
)

func DeserializeJSON() (users []User) {
	// read the data from the file
	// store it in users
	// return the users
	abspath, ferr := filepath.Abs(path)
	if ferr != nil {
		fmt.Println(ferr, "error finding file")
		return
	}
	data, err := ioutil.ReadFile(abspath)
	if err != nil {
		fmt.Println(err, "error reading from file")
	}
	uerr := json.Unmarshal(data, &users)
	if uerr != nil {
		fmt.Println(uerr, "error unmarshaling data")
	}
	return
}

func SerializeJSON(u []User) {
	// sort the users
	// SortData(u)
	data, err := json.MarshalIndent(u, "", "\t")
	if err != nil {
		fmt.Println(err)
		return
	}
	absPath, patherr := filepath.Abs(path)
	if patherr != nil {
		fmt.Println(patherr)
		return
	}
	er := ioutil.WriteFile(absPath, data, 0644)
	if er != nil {
		fmt.Println(er)
	}
}

type byName []User
type byAge []User
type byRollno []User
type byAddress []User

func (a byName) Len() int           { return len(a) }
func (a byName) Less(i, j int) bool { return a[i].Fullname < a[j].Fullname }
func (a byName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func (a byAge) Len() int           { return len(a) }
func (a byAge) Less(i, j int) bool { return a[i].Age < a[j].Age }
func (a byAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func (a byRollno) Len() int           { return len(a) }
func (a byRollno) Less(i, j int) bool { return a[i].Rollno < a[j].Rollno }
func (a byRollno) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func (a byAddress) Len() int           { return len(a) }
func (a byAddress) Less(i, j int) bool { return a[i].Address < a[j].Address }
func (a byAddress) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func SortDataByName(users byName) {
	sort.Sort(byName(users))
}

func SortDatabyAge(users byAge) {
	sort.Sort(byAge(users))
}

func SortDatabyRollno(users byRollno) {
	sort.Sort(byRollno(users))
}

func SortDatabyAddress(users byAddress) {
	sort.Sort(byAddress(users))
}
