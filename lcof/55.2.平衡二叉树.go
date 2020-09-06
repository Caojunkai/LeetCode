package lcof

var (
	isBalance = true
)

func isBalanced(root *TreeNode) bool {
	height(root)
	return isBalance
}

func height(root *TreeNode) int {
	if root == nil || !isBalance {
		return 0
	}
	leftHeight := height(root.Left)
	rightHeight := height(root.Right)
	if Abs(leftHeight-rightHeight) > 1 {
		isBalance = false
	}

	return Max(leftHeight, rightHeight) + 1
}

func Max(left, right int) int {
	if left > right {
		return left
	}

	return right
}

func Abs(v int) int {
	if v > 0 {
		return v
	}

	return -v
}
