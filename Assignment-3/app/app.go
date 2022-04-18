package app

import (
	"bufio"
	"fmt"
	"nuclei-assignment-3/controllers"
	"nuclei-assignment-3/models"
	"nuclei-assignment-3/services"
	"nuclei-assignment-3/utils"
	"os"
)

const (
	menuchoice          = "Enter your choice: "
	choicestringerr     = "choice can only be a integer"
	invalidchioiceerr   = "Invalid choice, choose between 1 to 9"
	readnodeidparents   = "enter the node id to get its parents: "
	readnodeidchild     = "enter the node id to get its childs: "
	readnodeidancestor  = "enter the node id to get its ancestors: "
	readnodeiddecendant = "enter the node id to get its decendants: "
	nodeiderr           = "invalid node id"
)

func StartApp() {
	var tree []models.Node
	services.PrintDetails()
	in := bufio.NewScanner(os.Stdin)
	fmt.Printf("%s", menuchoice)
loop:
	for in.Scan() {
		ch, err := utils.ValidateInput(in.Text())
		if err != nil {
			fmt.Println(err, choicestringerr)
			fmt.Printf("%s", menuchoice)
			continue
		}
		switch ch {
		case 1: // GET Immediate Parents of a node
			op := services.ReadInput(readnodeidparents, in)
			nodeid, err := utils.ValidateInput(op)
			if err != nil {
				fmt.Println(err, nodeiderr)
				fmt.Printf("%s", menuchoice)
				continue loop
			}
			controllers.GETImmediateParents(nodeid, tree)

		case 2: // GET Immediate child of a node
			op := services.ReadInput(readnodeidchild, in)
			nodeid, err := utils.ValidateInput(op)
			if err != nil {
				fmt.Println(err, nodeiderr)
				fmt.Printf("%s", menuchoice)
				continue loop
			}
			controllers.GETImmediateChilds(nodeid, tree)

		case 3: // GET Ancestors of a node
			op := services.ReadInput(readnodeidancestor, in)
			nodeid, err := utils.ValidateInput(op)
			if err != nil {
				fmt.Println(err, nodeiderr)
				fmt.Printf("%s", menuchoice)
				continue loop
			}
			controllers.GETAncestors(nodeid, tree)

		case 4: // GET Decendants of a node
			op := services.ReadInput(readnodeiddecendant, in)
			nodeid, err := utils.ValidateInput(op)
			if err != nil {
				fmt.Println(err, nodeiderr)
				fmt.Printf("%s", menuchoice)
				continue loop
			}
			controllers.GETDecendants(nodeid, tree)

		case 5: // DELETE Dependancy passing parent nodeid and child nodeid
		case 6: // DELETE node from the tree and all its dependancy
		case 7: // POST new depandancy passing parent nodeid and child nodeid, check cyclic dependancy
		case 8: // POST new root node
		case 9: // EXIT
			break loop
		default:
			fmt.Println(invalidchioiceerr)
			fmt.Printf("%s", menuchoice)
			continue loop
		}
		fmt.Printf("%s", menuchoice)
	}
}
