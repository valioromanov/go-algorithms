package tree

import (
	"fmt"
	"io"
)

type Node struct {
	data  int
	left  *Node
	right *Node
}

type BinaryTree struct {
	Root *Node
}

func (t *BinaryTree) InsertEl(data int) *BinaryTree {
	if t.Root == nil {
		t.Root = &Node{data, nil, nil}
	} else {
		t.Root.insert(data)
	}

	return t
}

func (n *Node) insert(data int) {
	if n == nil {
		return
	} else if data <= n.data {
		if n.left == nil {
			n.left = &Node{data, nil, nil}
		} else {
			n.left.insert(data)
		}
	} else {
		if n.right == nil {
			n.right = &Node{data, nil, nil}
		} else {
			n.right.insert(data)
		}
	}
}

func PrintTree(w io.Writer, node *Node, ns int, ch rune) {
	if node == nil {
		return
	}

	for i := 0; i < ns; i++ {
		fmt.Fprint(w, " ")
	}
	fmt.Fprintf(w, "%c:%v\n", ch, node.data)
	PrintTree(w, node.left, ns+2, 'L')
	PrintTree(w, node.right, ns+2, 'R')
}
