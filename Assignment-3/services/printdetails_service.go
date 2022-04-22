package services

import "fmt"

const (
	menu = `Operations:
1. Get the immediate parents of a node.
2. Get the immediate children of a node.
3. Get the ancestors of a node.
4. Get the descendants of a node.
5. Delete dependency from a tree, passing parent node id and child node id.
6. Delete a node from a tree, passing node id as input parameter. This should delete all the dependencies of the node.
7. Add a new dependency to a tree, passing parent node id and child node id. This should check for cyclic dependencies.
8. Add a new node to tree. This node will have no parents and children.
9. Exit.`
)

func PrintDetails() {
	fmt.Println(menu)
}
