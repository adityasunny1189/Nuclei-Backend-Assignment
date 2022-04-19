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
	} else {
		// fmt.Println(targetNode.Decendants)
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
	} else if targetNode.IsRoot {
		fmt.Println("root nodes have no ancestors")
	} else {
		// fmt.Println(targetNode.Ancestor)
	}
}
