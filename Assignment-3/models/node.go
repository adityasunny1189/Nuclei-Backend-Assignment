package models

type Node struct {
	IsRoot  bool
	NodeId  int
	Name    string
	Info    map[string][]string
	Parents []int
	Childs  []int
}
