// https://leetcode-cn.com/problems/shu-de-zi-jie-gou-lcof/

package lcof

func isSubStructure(A *TreeNode, B *TreeNode) bool {

	if A == nil || B == nil {
		return false
	}

	return isSubStructureWithRoot(A, B) || isSubStructure(A.Left, B) || isSubStructure(A.Right, B)
}

func isSubStructureWithRoot(root, target *TreeNode) bool {
	if target == nil {
		return true
	}

	if root == nil || root.Val != target.Val {
		return false
	}

	return isSubStructureWithRoot(root.Left, target.Left) && isSubStructureWithRoot(root.Right, target.Right)
}
