package main

import "fmt"

// TreeNode is a binary tree node.
type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

// LeftView returns the first node of every level, top to bottom — what an
// observer standing to the tree's left would see. DFS visiting left first:
// the first node to reach each new depth is that level's leftmost.
func LeftView(root *TreeNode) []int {
	view := []int{}
	var walk func(n *TreeNode, depth int)
	walk = func(n *TreeNode, depth int) {
		if n == nil {
			return
		}
		if depth == len(view) {
			view = append(view, n.Val) // first arrival at this depth
		}
		walk(n.Left, depth+1)
		walk(n.Right, depth+1)
	}
	walk(root, 0)
	return view
}

func tn(v int, l, r *TreeNode) *TreeNode { return &TreeNode{Val: v, Left: l, Right: r} }

func main() {
	root := tn(1, tn(2, nil, tn(4, nil, nil)), tn(3, nil, nil))
	fmt.Println(LeftView(root))
}
