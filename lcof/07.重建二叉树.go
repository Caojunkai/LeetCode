package lcof

func reConstructBinaryTree(pre []int, vin []int) *TreeNode {
	if len(pre) == 0 {
		return nil
	}

	root := &TreeNode{
		Val: pre[0],
	}

	if len(pre) == 1 {
		return root
	}

	mid := findMiddle(vin, pre[0])

	root.Left = reConstructBinaryTree(pre[1:mid+1], vin[:mid])
	root.Right = reConstructBinaryTree(pre[mid+1:], vin[mid+1:])
	return root
}

func findMiddle(vin []int, target int) int {
	for _, v := range vin {
		if v == target {
			return v
		}
	}
	return 0
}
