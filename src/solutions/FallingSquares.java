package solutions;

import java.util.ArrayList;
import java.util.List;

/**
 * 699 Falling Squares
 */
public class FallingSquares {
    public List<Integer> fallingSquares(int[][] positions) {
        int[] qans = new int[positions.length];

        for(int i = 0; i < positions.length; i++) {
            int left = positions[i][0];
            int size = positions[i][1];
            int right = left + size;

            qans[i] += size;

            for(int j = i + 1; j < positions.length; j++) {
                int left2 = positions[j][0];
                int size2 = positions[j][1];
                int right2 = left2 + size2;
                //落左半边， 落右半边
                if(left2 < right && left < right2) {
                    qans[j] = Math.max(qans[j], qans[i]);
                }
            }
        }
        List<Integer> ans = new ArrayList<>();

        int cur = -1;
        for(int x : qans) {
            cur = Math.max(cur, x);
            ans. add(cur);
        }

        return ans;
    }
}
