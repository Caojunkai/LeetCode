// https://leetcode-cn.com/problems/cong-shang-dao-xia-da-yin-er-cha-shu-lcof/

/*

从上到下打印出二叉树的每个节点，同一层的节点按照从左到右的顺序打印。



例如:
给定二叉树: [3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7
返回：

[3,9,20,15,7]

*/

package lcof

func levelOrder(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	var node *TreeNode
	rsp, queue := make([]int, 0), make([]*TreeNode, 0)

	queue = append(queue, root)

	for i := 0; i < len(queue); i++ {
		node = queue[i]
		rsp = append(rsp, node.Val)

		if node.Left != nil {
			queue = append(queue, node.Left)
		}

		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}

	return rsp
}
