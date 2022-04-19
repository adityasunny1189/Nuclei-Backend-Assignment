package controllers

import (
	"fmt"
	"nuclei-assignment-3/models"
	"nuclei-assignment-3/utils"
)

func getIndex(id int, nodes []int) int {
	for i, val := range nodes {
		if id == val {
			return i
		}
	}
	return -1
}

func DELETEDependancy(pid, cid int, tree []models.Node) []models.Node {
	t := make([]models.Node, len(tree))
	if ok := utils.ValidNode(pid, tree); !ok {
		fmt.Println(invalidparent)
		t = tree
		return t
	}
	if ok := utils.ValidNode(cid, tree); !ok {
		fmt.Println(invalidchild)
		t = tree
		return t
	}

	for i, node := range tree {
		t[i] = node
		if node.NodeId == pid {
			j := getIndex(cid, node.Childs)
			if j == -1 {
				fmt.Println("no child node invalid relationship")
				return tree
			}
			t[i].Childs = append(node.Childs[0:j], node.Childs[j+1:]...)
		} else if node.NodeId == cid {
			j := getIndex(pid, node.Parents)
			t[i].Parents = append(node.Parents[0:j], node.Parents[j+1:]...)
			if len(t[i].Parents) == 0 {
				t[i].IsRoot = true
			}
		}
	}
	return t
}

func DELETENode(pid int, tree *[]models.Node) {
	var parentnode models.Node
	var pindex int
	for i, node := range *tree {
		if node.NodeId == pid {
			parentnode = node
			pindex = i
		}
	}
	for _, c := range parentnode.Childs {
		*tree = DELETEDependancy(pid, c, *tree)
	}
	for _, p := range parentnode.Parents {
		*tree = DELETEDependancy(p, pid, *tree)
	}
	*tree = append((*tree)[:pindex], (*tree)[pindex+1:]...)
}
