package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	var q TreeNode
	d := maxDepth(&q)
	fmt.Println("vim-go ", d)
}

func maxDepth(root *TreeNode) int {
	ldepth := 0
	rdepth := 0
	if root == nil {
		return 0
	} else {
		ldepth += maxDepth(root.Left)
		rdepth += maxDepth(root.Right)

		if ldepth > rdepth {
			return ldepth + 1
		} else {
			return rdepth + 1
		}
	}
}
