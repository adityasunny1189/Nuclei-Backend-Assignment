package handlers

import (
	"bufio"
	"fmt"
	"nuclei-assignment-3/controllers"
	"nuclei-assignment-3/models"
	"nuclei-assignment-3/services"
	"nuclei-assignment-3/utils"
)

const (
	readnodeidparents   = "enter the node id to get its parents: "
	readnodeidchild     = "enter the node id to get its childs: "
	readnodeidancestor  = "enter the node id to get its ancestors: "
	readnodeiddecendant = "enter the node id to get its descendants: "
)

func GetParents(tree []models.Node, in *bufio.Scanner) {
	op := services.ReadInput(readnodeidparents, in)
	nodeid, err := utils.ValidateInput(op)
	if err != nil {
		fmt.Println(err, nodeiderr)
		return
	}
	controllers.GETImmediateParents(nodeid, tree)
}

func GetChilds(tree []models.Node, in *bufio.Scanner) {
	op := services.ReadInput(readnodeidchild, in)
	nodeid, err := utils.ValidateInput(op)
	if err != nil {
		fmt.Println(err, nodeiderr)
		return
	}
	controllers.GETImmediateChilds(nodeid, tree)
}

func GetAncestors(tree []models.Node, in *bufio.Scanner) {
	op := services.ReadInput(readnodeidancestor, in)
	nodeid, err := utils.ValidateInput(op)
	if err != nil {
		fmt.Println(err, nodeiderr)
		return
	}
	controllers.GETAncestors(nodeid, tree)
}

func GetDecendants(tree []models.Node, in *bufio.Scanner) {
	op := services.ReadInput(readnodeiddecendant, in)
	nodeid, err := utils.ValidateInput(op)
	if err != nil {
		fmt.Println(err, nodeiderr)
		return
	}
	controllers.GETDecendants(nodeid, tree)
}
