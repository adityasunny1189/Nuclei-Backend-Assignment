package controllers

import (
	"bufio"
	"fmt"
	"nuclei-assignment-3/models"
	"nuclei-assignment-3/utils"
	"strconv"
)

const (
	invalidparent    = "invalid parent node id"
	invalidchild     = "invalid child node id"
	cyclicdependancy = "this dependancy will result in cyclic dependancy, not allowed"
)

func POSTDependancy(pid, cid int, tree []models.Node) []models.Node {
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
	if ok := utils.CheckCyclicDependancy(pid, cid, tree); ok {
		fmt.Println(cyclicdependancy)
		t = tree
		return t
	}

	for i, node := range tree {
		t[i] = node
		if node.NodeId == pid {
			t[i].Childs = append(node.Childs, cid)
		} else if node.NodeId == cid {
			t[i].IsRoot = false
			t[i].Parents = append(t[i].Parents, pid)
		}
	}
	return t
}

func readNode(in *bufio.Scanner, tree *[]models.Node) (n models.Node) {
	fmt.Println("enter the node details:")
	fmt.Printf("node id: ")
	in.Scan()
	op := in.Text()
	nid, err := strconv.Atoi(op)
	if err != nil {
		fmt.Println(err)
		return
	}
	ok := utils.ValidNode(nid, *tree)
	if ok {
		fmt.Println("nodeid already present")
		return
	}
	n.NodeId = nid
	fmt.Printf("enter node name: ")
	in.Scan()
	n.Name = in.Text()
	n.IsRoot = true
	return
}

func POSTNode(in *bufio.Scanner, tree *[]models.Node) {
	node := readNode(in, tree)
	if node.NodeId != 0 {
		*tree = append(*tree, node)
	}
}
