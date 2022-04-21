package user

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"sort"
)

func DeserializeJSON() (users []User) {
	// read the data from the file
	// store it in users
	// return the users
	abspath, ferr := filepath.Abs("../A2/data/users_data.json")
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
	absPath, patherr := filepath.Abs("../A2/data/users_data.json")
	if patherr != nil {
		fmt.Println(patherr)
		return
	}
	er := ioutil.WriteFile(absPath, data, 0644)
	if er != nil {
		fmt.Println(er)
	}
}

type userlist []User

func (a userlist) Len() int {
	return len(a)
}
func (a userlist) Less(i, j int) bool {
	return a[i].Fullname < a[j].Fullname
}
func (a userlist) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func SortData(users userlist) {
	sort.Sort(userlist(users))
}
