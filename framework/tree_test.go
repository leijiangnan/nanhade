package framework

import (
	"fmt"
	"testing"
)

var root *node = &node{
	isLast:  false,
	segment: ":id",
	handler: nil,
	childs: []*node{
		{
			isLast:  true,
			segment: "name",
			handler: nil,
			childs:  nil,
		},
		{
			isLast:  true,
			segment: "age",
			handler: nil,
			childs:  nil,
		},
	},
}

func Test_node_filterChildNodes(t *testing.T) {

	nodes := root.filterChildNodes(":id")
	fmt.Println(len(nodes))
	for _, n := range nodes {
		fmt.Println(n.segment)
	}
}

func Test_node_matchNode(t *testing.T) {
	node := root.matchNode(":id")
	fmt.Println(node.segment)
}
