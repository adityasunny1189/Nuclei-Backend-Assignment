package models

type Node struct {
	IsRoot     bool
	NodeId     int
	Name       string
	Info       map[string][]string
	Parents    []Node
	Ancestor   []Node
	Childs     []Node
	Decendants []Node
}
