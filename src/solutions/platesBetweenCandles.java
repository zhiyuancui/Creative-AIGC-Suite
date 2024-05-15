class Solution {
    public int[] platesBetweenCandles(String s, int[][] queries) {
        int[] prefixSum = new int[s.length()+1];
        int[] next = new int[s.length()+1];
        int[] prev = new int[s.length()+1];

        Arrays.fill(next, Integer.MAX_VALUE);

        int[] res = new int[queries.length];

        for(int i = 0; i < s.length(); i++) {
            prefixSum[i+1] = prefixSum[i] + (s.charAt(i) == '|' ? 1 : 0);
            prev[i+1] = s.charAt(i) == '|' ? i : prev[i];
        }

        for(int i = s.length() - 1; i >= 0; i--) {
            next[i] = s.charAt(i) == '|' ? i : next[i+1];
        }

        for(int j = 0; j < queries.length; j++) {
            int left = next[queries[j][0]];
            int right = prev[queries[j][1] + 1];
            res[j] = left < right ? right - left - (prefixSum[right] - prefixSum[left]) : 0;
        }

        return res;
    }
}