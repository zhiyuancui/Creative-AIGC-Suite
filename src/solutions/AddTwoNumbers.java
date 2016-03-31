package solutions;

import util.ListNode;

public class AddTwoNumbers {
	
	public ListNode addTwoNumbers(ListNode l1, ListNode l2) {
        ListNode head1 = l1;
        ListNode head2 = l2;
        ListNode dummy = new ListNode(0);
        ListNode cur = dummy;
        
        int carry = 0;
        
        while( head1 != null || head2 != null || carry != 0 ){
            int sum = ( head1 != null ? head1.val : 0 )
                    + ( head2 != null ? head2.val : 0 ) + carry;
            carry = sum / 10;
            int num = sum % 10;
            cur.next = new ListNode( num );
            cur = cur.next;
            head1 = ( head1 != null ? head1.next : null );
            head2 = head2 != null ? head2.next : null;
        }
        
        return dummy.next;
    }
	
	public static void main(String[] args){
		AddTwoNumbers a = new AddTwoNumbers();
		ListNode l1 = new ListNode(0);
		ListNode l2 = new ListNode(0);
		a.addTwoNumbers(l1, l2);
	}
}
