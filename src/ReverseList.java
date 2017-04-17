/* Created by Fitz on 2017/3/21 */

import util.DListNode;

public class ReverseList {
    //    1->2->3->null  --->  3->2->1->null
    public static void main(String[] args) {

    }

    public ListNode reverse(ListNode listNode){
        ListNode prev = null;
        while (listNode != null){
            ListNode next = listNode.next;
            listNode.next = prev;
            prev = listNode;
            listNode = next;
        }
        return prev;
    }

    public DListNode dListNodeReverse(DListNode dListNode){
        DListNode curr = null;
        while (dListNode != null){
            curr = dListNode;
            dListNode = curr.next;
            curr.next = curr.prev;
            curr.prev = dListNode;
        }
        return curr;
    }
}
