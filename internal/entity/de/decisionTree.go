package de

// DecisionTree represents the main decision tree structure
type DecisionTree struct {
	Name       string `json:"name"`
	UUID       string `json:"uuid"`
	ClientCode string `json:"client_code"`
	FormUUID   string `json:"form_uuid"`
	Active     bool   `json:"active"`
	Root       *Node  `json:"root"`
}

// Node represents a decision tree node (used recursively)
type Node struct {
	Name        string  `json:"name"`
	Title       string  `json:"title,omitempty"`
	Evaluation  string  `json:"evaluation,omitempty"`
	SubTreeUUID string  `json:"sub_tree_uuid,omitempty"`
	Nodes       []*Node `json:"nodes,omitempty"`
	Fallback    *Node   `json:"fallback,omitempty"`
}
