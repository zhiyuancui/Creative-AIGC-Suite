package solutions;

import java.util.Arrays;

/**
 * #455 Assign Cookies
 */
public class AssignCookies {

	public int findContentChildren(int[] greedy, int[] cookies) {
		 
        int count = 0;
        
        Arrays.sort(greedy);
        Arrays.sort(cookies);
        
        int i = 0;
        int j = 0;
        while( i < greedy.length && j < cookies.length ) {
            if( cookies[j] >= greedy[i] ){
                j++;
                i++;
                count++;
            } else {
                j++;
            }
        }
        
        
        return count;
    }
}
