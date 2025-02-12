package solutions;

import java.util.ArrayList;
import java.util.List;

/**
 * #120 Triangle
 */

public class Triangle {
	public int minimumTotal(List<List<Integer>> triangle) {
        if( triangle == null ){
            return 0;
        }
        
        int[] dp = new int[triangle.size()];
        
        int count = 0;
        for(int i : triangle.get(triangle.size() - 1) ){
            dp[ count ] = i;
            count++;
        }
        
        for(int row = triangle.size() - 2; row >= 0 ; row--){
            
            List<Integer> level = triangle.get(row);
            for(int i = 0; i < level.size(); i++){
                dp[i] = Math.min( dp[i],dp[i+1]) + level.get(i);
            }
        }
        
        
        return dp[0];
    }
	
}
