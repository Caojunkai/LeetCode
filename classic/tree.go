package classic

type TreeNode struct {
	val   string
	left  *TreeNode
	right *TreeNode
	count int
}

// func NewBstTree(v int) *TreeNode {
// 	return &TreeNode{
// 		val: v,
// 	}
// }

// func (t *TreeNode) Insert(num int) {
// 	if num < t.val {
// 		if t.left == nil {
// 			t.left = &TreeNode{
// 				val: num,
// 			}
// 		} else {
// 			t.left.Insert(num)
// 		}
// 		return
// 	}
//
// 	if num > t.val {
// 		if t.right == nil {
// 			t.right = &TreeNode{
// 				val: num,
// 			}
// 		} else {
// 			t.right.Insert(num)
// 		}
// 		return
// 	}
//
// 	if num == t.val {
// 		t.count++
// 	}
//
// }
//
// func (t *TreeNode) Dfs() {
// 	log.Print(t.val)
// 	if t.left != nil {
// 		t.left.Dfs()
// 	}
// 	if t.right != nil {
// 		t.right.Dfs()
// 	}
// }
