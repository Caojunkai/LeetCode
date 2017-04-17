/* Created by Fitz on 2017/3/15 */

public class AddTwoNum {
    public static void main(String[] args) {
        AddTwoNum addTwoNum = new AddTwoNum();
        ListNode listNode_1_0 = new ListNode(3);
        ListNode listNode_1_1 = new ListNode(5);
        ListNode listNode_1_2 = new ListNode(7);
        listNode_1_0.next = listNode_1_1;
        listNode_1_1.next = listNode_1_2;

        ListNode listNode_2_0 = new ListNode(4);
        ListNode listNode_2_1 = new ListNode(6);
        ListNode listNode_2_2 = new ListNode(8);
        listNode_2_0.next = listNode_2_1;
        listNode_2_1.next = listNode_2_2;

        ListNode listNodeResult = addTwoNum.solution(listNode_1_0, listNode_2_0);
    }

    public ListNode solution(ListNode l1, ListNode l2) {
        return helper(l1, l2, 0);
    }

    public ListNode helper(ListNode l1, ListNode l2, int carry) {
        if (l1 == null && l2 == null) {
            return carry == 0 ? null : new ListNode(carry);
        }
        if (l1 == null && l2 != null) {
            l1 = new ListNode(0);
        }
        if (l2 == null && l1 != null) {
            l2 = new ListNode(0);
        }
        int sum = l1.val + l2.val + carry;
        ListNode curr = new ListNode(sum % 10);
        curr.next = helper(l1.next, l2.next, sum / 10);
        return curr;
    }
}

class ListNode {
    int val;
    ListNode next;

    ListNode(int x) {
        val = x;
    }
}
