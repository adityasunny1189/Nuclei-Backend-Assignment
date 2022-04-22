package handlers

import (
	"bufio"
	"fmt"
	"nuclei-assignment-3/controllers"
	"nuclei-assignment-3/models"
	"nuclei-assignment-3/services"
	"nuclei-assignment-3/utils"
)

func AddDependancy(tree []models.Node, in *bufio.Scanner) []models.Node {
	pop := services.ReadInput(readparentnodeid, in)
	pnodeid, err := utils.ValidateInput(pop)
	if err != nil {
		fmt.Println(err, nodeiderr)
		return tree
	}
	cop := services.ReadInput(readchildnodeid, in)
	cnodeid, cerr := utils.ValidateInput(cop)
	if cerr != nil {
		fmt.Println(err, nodeiderr)
		return tree
	}
	tree = controllers.POSTDependancy(pnodeid, cnodeid, tree)
	return tree
}

func AddNode(tree *[]models.Node, in *bufio.Scanner) {
	controllers.POSTNode(in, tree)
}
