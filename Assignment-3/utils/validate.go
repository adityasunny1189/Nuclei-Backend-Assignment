package utils

import (
	"nuclei-assignment-3/models"
	"strconv"
)

func ValidateInput(ch string) (op int, err error) {
	res, err := strconv.Atoi(ch)
	if err != nil {
		return
	}
	op = res
	return
}

func ValidNode(nid int, tree []models.Node) bool {
	for _, node := range tree {
		if node.NodeId == nid {
			return true
		}
	}
	return false
}
