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
	readparentnodeid = "enter the parent node id: "
	readchildnodeid  = "enter the child node id: "
	nodeiderr        = "invalid node id"
)

func DeleteDependancy(tree []models.Node, in *bufio.Scanner) []models.Node {
	t := tree
	pop := services.ReadInput(readparentnodeid, in)
	pnodeid, perr := utils.ValidateInput(pop)
	if perr != nil {
		fmt.Println(perr, nodeiderr)
		return t
	}
	cop := services.ReadInput(readchildnodeid, in)
	cnodeid, cerr := utils.ValidateInput(cop)
	if cerr != nil {
		fmt.Println(cerr, nodeiderr)
		return t
	}
	t = controllers.DELETEDependancy(pnodeid, cnodeid, tree)
	return t
}

func DeleteNode(tree *[]models.Node, in *bufio.Scanner) {
	pop := services.ReadInput(readparentnodeid, in)
	pnodeid, perr := utils.ValidateInput(pop)
	if perr != nil {
		fmt.Println(perr, nodeiderr)
	}
	controllers.DELETENode(pnodeid, tree)
}
