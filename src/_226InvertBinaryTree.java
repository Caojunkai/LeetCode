/* Created by Fitz on 2017/3/23 */

import util.TreeNode;

public class _226InvertBinaryTree {
    public static void main(String[] args) {
        TreeNode treeNode_1 = new TreeNode(1);
        TreeNode treeNode_2 = new TreeNode(2);
        TreeNode treeNode_3 = new TreeNode(3);
        TreeNode treeNode_4 = new TreeNode(4);
        TreeNode treeNode_5 = new TreeNode(5);
        TreeNode treeNode_6 = new TreeNode(6);
        TreeNode treeNode_7 = new TreeNode(7);
        treeNode_1.left = treeNode_2;
        treeNode_1.right = treeNode_3;
        treeNode_2.right = treeNode_5;
        treeNode_3.left = treeNode_6;
        treeNode_3.right = treeNode_7;
        solution(treeNode_1);
        System.out.println(treeNode_1.left.val);
        System.out.println(treeNode_1.right.val);
        System.out.println(treeNode_2.left.val);
    }

    public static void solution(TreeNode treeNode){
        if (treeNode == null){
            return;
        }
        TreeNode temp = treeNode.left;
        treeNode.left = treeNode.right;
        treeNode.right = temp;
        solution(treeNode.left);
        solution(treeNode.right);
    }
}
