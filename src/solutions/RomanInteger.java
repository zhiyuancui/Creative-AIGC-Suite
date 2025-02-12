package solutions;

import java.util.*;

public class RomanInteger {
	
	/**
	 * #13 Roman to Integer
	 * @param s
	 * @return
	 */
	public int romanToInt(String s) {
		if( s == null || s.length() == 0 ) {
			return 0;
		} 
		
		HashMap<Character, Integer> map = new HashMap<Character, Integer>();
		
		map.put('I',1);
		map.put('V',5);
		map.put('X',10);
		map.put('L',50);
		map.put('C',100);
		map.put('D',500);
		map.put('M',1000);
		
		int len = s.length();
		
		int result = 0;
		
		for(int i = 0; i < len;i++) {
			int val1 = map.get( s.charAt(i)   );
			int val2 = i + 1 < len ? map.get( s.charAt(i+1) ) : 0;
			
			if( val1 < val2 ) {
				result += val2 - val1;
				i++;
			}
			else{
				result += val1;
			}
		}
		
		return result;  
	}
}
