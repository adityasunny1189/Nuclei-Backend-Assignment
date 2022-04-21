package app

import (
	"bufio"
	"fmt"
	"nuclei-assignment-3/handlers"
	"nuclei-assignment-3/models"
	"nuclei-assignment-3/services"
	"nuclei-assignment-3/utils"
	"os"
)

const (
	menuchoice        = "Enter your choice: "
	choicestringerr   = "choice can only be a integer"
	invalidchioiceerr = "Invalid choice, choose between 1 to 9"
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
			handlers.GetParents(tree, in)

		case 2: // GET Immediate child of a node
			handlers.GetChilds(tree, in)

		case 3: // GET Ancestors of a node
			handlers.GetAncestors(tree, in)

		case 4: // GET Descendants of a node
			handlers.GetDecendants(tree, in)

		case 5: // DELETE Dependency passing parent nodeid and child nodeid
			tree = handlers.DeleteDependancy(tree, in)

		case 6: // DELETE node from the tree and all its dependency
			handlers.DeleteNode(&tree, in)

		case 7: // POST new dependency passing parent nodeid and child nodeid, check cyclic dependency
			tree = handlers.AddDependancy(tree, in)

		case 8: // POST new root node
			handlers.AddNode(&tree, in)

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
