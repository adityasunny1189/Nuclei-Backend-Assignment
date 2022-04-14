package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"nuclei-assignment-2/models"
	"path/filepath"
)

func DeserializeJSON() (users []models.User) {
	// read the data from the file
	// store it in users
	// return the users
	abspath, ferr := filepath.Abs("../Assignment-2/data/users_data.json")
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
