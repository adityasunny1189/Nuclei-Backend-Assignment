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
	// have to check that the parent id passed does not have the child id passed as its parent or its ancestor
	var parentnode, childnode models.Node
	for _, node := range tree {
		if node.NodeId == pid {
			parentnode = node
		} else if node.NodeId == cid {
			childnode = node
		}
	}
	if len(parentnode.Parents) == 0 {
		return false
	}
	for _, parent := range parentnode.Parents {
		if childnode.NodeId == parent {
			return true
		}
		res := CheckCyclicDependancy(parent, cid, tree)
		if res {
			return res
		}
	}
	return false
}
