// https://leetcode-cn.com/problems/cong-shang-dao-xia-da-yin-er-cha-shu-ii-lcof/

/*

从上到下按层打印二叉树，同一层的节点按从左到右的顺序打印，每一层打印到一行。



例如:
给定二叉树: [3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7
返回其层次遍历结果：

[
  [3],
  [9,20],
  [15,7]
]

*/

package lcof

func levelOrder2(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	resp := make([][]int, 0)
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		nextLevelQueue := make([]*TreeNode, 0)
		levelResp := make([]int, 0)

		for _, v := range queue {
			levelResp = append(levelResp, v.Val)
			if v.Left != nil {
				nextLevelQueue = append(nextLevelQueue, v.Left)
			}

			if v.Right != nil {
				nextLevelQueue = append(nextLevelQueue, v.Right)
			}
		}

		resp = append(resp, levelResp)
		queue = nextLevelQueue
	}

	return resp
}
