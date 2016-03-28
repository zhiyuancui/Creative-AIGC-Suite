package solutions;

import util.ListNode;

public class IsPalindrome {
    
	/**
	 * Determine whether an integer is a palindrome. 
	 * Do this without extra space.
	 * @param x
	 * @return
	 */
	public boolean isPalindrome(int x) {
        
		
        if( x < 0 )
        {
            return false;
        }
        
        if( x < 10 )
        {
            return true;
        }
        
        String number = Integer.toString(x);
        
        int front = 0;
        int end = number.length() - 1;
        
        while( front < end )
        {
            if( number.charAt(front) != number.charAt(end) )
            {
                return false;
            }
            else
            {
                front++;
                end--;
            }
            
        }
        
        return true;
    }
	
	
	/**
	 * Palindrome Linked List
	 * @param head
	 * @return
	 */
	public boolean isPalindrome(ListNode head) {
        if( head == null || head.next == null ){
            return true;
        }
        
        ListNode cur = head;
        int len = 0;
        while( cur != null ){
            len++;
            cur = cur.next;
        }
        
        int mid = len / 2;
        
        cur = head;
        ListNode newHead = null;
        for(int i = 0; i < mid; i++){
            ListNode next = cur.next;
            cur.next = newHead;
            newHead =cur;
            cur = next;
        }
        
        if( len % 2 == 1 ){
            cur = cur.next;
        }
         
        for(int i = 0; i < mid; i++ ){
            if( newHead.val != cur.val ){
                return false;
            }
            newHead = newHead.next;
            cur = cur.next;
        }
        return true;
    }
	
	
	/**
	 * Valid Palindrome
	 * @param s
	 * @return
	 */
	public boolean isPalindrome(String s) {
        if( s== null || s.length() == 0 ){
            return true;
        }
        
        s = s.toLowerCase();
        
        int len = s.length();
        int front = 0;
        int end = len - 1;
        
        while( front < end ){
            while( front < s.length() && !isLetter( s.charAt(front) ) ){
                front++;
            }
            
            if( front == s.length() ){
                return true;
            }
            
            while( end >= 0 && !isLetter( s.charAt(end) ) ){
                end--;
            }
            
            if(  s.charAt( front) != s.charAt(end) ){
               break;
            }else {
                front++;
                end--;
            }
        }
        
        return end <= front;
    }
    
    
    private boolean isLetter(Character c){
        return Character.isLetter(c) || Character.isDigit(c);
    }
	
}
