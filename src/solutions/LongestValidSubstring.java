class LongestValidSubstring {
    public int longestValidSubstring(String word, List<String> forbidden) {
        Set<String> fb = new HashSet<>();

        fb.addAll(forbidden);

        int ans = 0, left = 0, n = word.length();

        for(int right = 0; right < n; right++) {
            for(int i = right; i >= left && i > right - 10; i--) {
                if(fb.contains(word.substring(i, right+1))) {
                    left = i +1;
                    break;
                }
            }
            ans = Math.max(ans, right - left + 1);
        }

        return ans;
    }
}