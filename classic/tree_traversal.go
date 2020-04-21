package classic

import "log"

func CreateTree() TreeNode {
	nodeG := TreeNode{val: "g", left: nil, right: nil}
	nodeF := TreeNode{val: "f", left: &nodeG, right: nil}
	nodeE := TreeNode{val: "e", left: nil, right: nil}
	nodeD := TreeNode{val: "d", left: &nodeE, right: nil}
	nodeC := TreeNode{val: "c", left: nil, right: nil}
	nodeB := TreeNode{val: "b", left: &nodeD, right: &nodeF}
	nodeA := TreeNode{val: "a", left: &nodeB, right: &nodeC}
	return nodeA
}

func PreOrderTraversal(root *TreeNode) {
	if root == nil {
		return
	}
	print(root.val, "---> ")
	PreOrderTraversal(root.left)
	PreOrderTraversal(root.right)
}

func InOrderTraversal(root *TreeNode) {
	if root == nil {
		return
	}
	InOrderTraversal(root.left)
	print(root.val, "---> ")
	InOrderTraversal(root.right)
}

func PostOrderTraversal(root *TreeNode) {
	if root == nil {
		return
	}
	PostOrderTraversal(root.left)
	PostOrderTraversal(root.right)
	print(root.val, "---> ")
}

func Bfs(root *TreeNode) {
	if root == nil {
		return
	}
	log.Println(root.val)
	Bfs(root.left)
	Bfs(root.right)
}

func Dfs(root *TreeNode) {
	if root == nil {
		return
	}
	log.Println(root.val)
	if root.left != nil {
		Dfs(root.left)
	}

	if root.right != nil {
		Dfs(root.right)
	}
}
