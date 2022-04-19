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

func CheckCyclicDependancy(pid, cid int, tree []models.Node) bool {
	// have to check that the child node id(cid) and parent node id will not create any cyclic dependancy
	// no one can be his own parent or ancestor
	return true
}
