package controllers

import (
	"fmt"
	"nuclei-assignment-3/models"
)

func GETImmediateParents(nid int, tree []models.Node) {
	var targetNode models.Node
	for _, n := range tree {
		if n.NodeId == nid {
			targetNode = n
		}
	}
	if targetNode.NodeId == 0 {
		fmt.Println("node id not found")
	} else if targetNode.IsRoot {
		fmt.Println("root nodes have no parents")
	} else {
		fmt.Println(targetNode.Parents)
	}
}

func GETImmediateChilds(nid int, tree []models.Node) {
	var targetNode models.Node
	for _, n := range tree {
		if n.NodeId == nid {
			targetNode = n
		}
	}
	if targetNode.NodeId == 0 {
		fmt.Println("node id not found")
	} else {
		fmt.Println(targetNode.Childs)
	}
}

func GETDecendants(nid int, tree []models.Node) {
	var targetNode models.Node
	for _, n := range tree {
		if n.NodeId == nid {
			targetNode = n
		}
	}
	if targetNode.NodeId == 0 {
		fmt.Println("node id not found")
	}
	if len(targetNode.Childs) > 0 {
		fmt.Println(targetNode.Childs)
		for _, val := range targetNode.Childs {
			GETDecendants(val, tree)
		}
	}
}

func GETAncestors(nid int, tree []models.Node) {
	var targetNode models.Node
	for _, n := range tree {
		if n.NodeId == nid {
			targetNode = n
		}
	}
	if targetNode.NodeId == 0 {
		fmt.Println("node id not found")
	}
	if len(targetNode.Parents) > 0 {
		fmt.Println(targetNode.Parents)
		for _, val := range targetNode.Parents {
			GETAncestors(val, tree)
		}
	}
}
