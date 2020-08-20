package classic

import "log"

func CreateTree() TreeNode {
	nodeG := TreeNode{Val: "g", Left: nil, Right: nil}
	nodeF := TreeNode{Val: "f", Left: &nodeG, Right: nil}
	nodeE := TreeNode{Val: "e", Left: nil, Right: nil}
	nodeD := TreeNode{Val: "d", Left: &nodeE, Right: nil}
	nodeC := TreeNode{Val: "c", Left: nil, Right: nil}
	nodeB := TreeNode{Val: "b", Left: &nodeD, Right: &nodeF}
	nodeA := TreeNode{Val: "a", Left: &nodeB, Right: &nodeC}
	return nodeA
}

func PreOrderTraversal(root *TreeNode) {
	if root == nil {
		return
	}
	print(root.Val, "---> ")
	PreOrderTraversal(root.Left)
	PreOrderTraversal(root.Right)
}

func InOrderTraversal(root *TreeNode) {
	if root == nil {
		return
	}
	InOrderTraversal(root.Left)
	print(root.Val, "---> ")
	InOrderTraversal(root.Right)
}

func PostOrderTraversal(root *TreeNode) {
	if root == nil {
		return
	}
	PostOrderTraversal(root.Left)
	PostOrderTraversal(root.Right)
	print(root.Val, "---> ")
}

func Bfs(root *TreeNode) {
	if root == nil {
		return
	}
	log.Println(root.Val)
	Bfs(root.Left)
	Bfs(root.Right)
}

func Dfs(root *TreeNode) {
	if root == nil {
		return
	}
	log.Println(root.Val)
	if root.Left != nil {
		Dfs(root.Left)
	}

	if root.Right != nil {
		Dfs(root.Right)
	}
}
