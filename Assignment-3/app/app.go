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
	readnodeiddecendant = "enter the node id to get its descendants: "
	readparentnodeid    = "enter the parent node id: "
	readchildnodeid     = "enter the child node id: "
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

		case 4: // GET Descendants of a node
			op := services.ReadInput(readnodeiddecendant, in)
			nodeid, err := utils.ValidateInput(op)
			if err != nil {
				fmt.Println(err, nodeiderr)
				fmt.Printf("%s", menuchoice)
				continue loop
			}
			controllers.GETDecendants(nodeid, tree)

		case 5: // DELETE Dependency passing parent nodeid and child nodeid
			pop := services.ReadInput(readparentnodeid, in)
			pnodeid, perr := utils.ValidateInput(pop)
			if perr != nil {
				fmt.Println(err, nodeiderr)
				fmt.Printf("%s", menuchoice)
				continue loop
			}
			cop := services.ReadInput(readchildnodeid, in)
			cnodeid, cerr := utils.ValidateInput(cop)
			if cerr != nil {
				fmt.Println(err, nodeiderr)
				fmt.Printf("%s", menuchoice)
				continue loop
			}
			tree = controllers.DELETEDependancy(pnodeid, cnodeid, tree)

		case 6: // DELETE node from the tree and all its dependency
			pop := services.ReadInput(readparentnodeid, in)
			pnodeid, perr := utils.ValidateInput(pop)
			if perr != nil {
				fmt.Println(err, nodeiderr)
				fmt.Printf("%s", menuchoice)
				continue loop
			}
			controllers.DELETENode(pnodeid, &tree)

		case 7: // POST new dependency passing parent nodeid and child nodeid, check cyclic dependency
			pop := services.ReadInput(readparentnodeid, in)
			pnodeid, perr := utils.ValidateInput(pop)
			if perr != nil {
				fmt.Println(err, nodeiderr)
				fmt.Printf("%s", menuchoice)
				continue loop
			}
			cop := services.ReadInput(readchildnodeid, in)
			cnodeid, cerr := utils.ValidateInput(cop)
			if cerr != nil {
				fmt.Println(err, nodeiderr)
				fmt.Printf("%s", menuchoice)
				continue loop
			}
			tree = controllers.POSTDependancy(pnodeid, cnodeid, tree)

		case 8: // POST new root node
			controllers.POSTNode(in, &tree)

		case 9: // EXIT
			break loop
		default:
			fmt.Println(invalidchioiceerr)
			fmt.Printf("%s", menuchoice)
			continue loop
		}
		fmt.Println(tree)
		fmt.Printf("%s", menuchoice)
	}
}
