package classic

type TreeNode struct {
	Val   string
	Left  *TreeNode
	Right *TreeNode
	Count int
}

func preorderTravel(root *TreeNode) {
	if root == nil {
		return
	}
	println(root.Val)
	preorderTravel(root.Left)
	preorderTravel(root.Right)
}

func inorderTraversal(root *TreeNode) {
	if root == nil {
		return
	}
	inorderTraversal(root.Left)
	println(root.Val)
	inorderTraversal(root.Right)
}

func postorderTraversal(root *TreeNode) {
	if root == nil {
		return
	}
	postorderTraversal(root.Left)
	postorderTraversal(root.Right)
	println(root.Val)
}

func dfs(root *TreeNode) {
	if root == nil {
		return
	}
	println(root.Val)
	dfs(root.Left)
	dfs(root.Right)
}
