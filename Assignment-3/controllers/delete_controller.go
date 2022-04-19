package controllers

import (
	"fmt"
	"nuclei-assignment-3/models"
	"nuclei-assignment-3/utils"
)

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
}

func DELETENode(pid int, tree *[]models.Node) {

}
