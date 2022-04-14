package services

import (
	"nuclei-assignment-2/models"
	"sort"
)

type userslist []models.User

func (a userslist) Len() int {
	return len(a)
}
func (a userslist) Less(i, j int) bool {
	return a[i].Fullname < a[j].Fullname
}
func (a userslist) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func SortData(users userslist) {
	sort.Sort(userslist(users))
}
