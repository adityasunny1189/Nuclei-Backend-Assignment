package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"nuclei-assignment-2/models"
	"path/filepath"
)

func SaveUserDetails(u []models.User) {
	// sort the users
	data, err := json.MarshalIndent(u, "", "\t")
	if err != nil {
		fmt.Println(err)
		return
	}
	absPath, patherr := filepath.Abs("../Assignment-2/data/users_data.json")
	if patherr != nil {
		fmt.Println(patherr)
		return
	}
	er := ioutil.WriteFile(absPath, data, 0644)
	if er != nil {
		fmt.Println(er)
	}
}
