package solutions.amazon;

/**
 * 2405. Optimal Partition of String
 */
class OptimalPartitionString {
    public int partitionString(String s) {
        StringBuilder lookup = new StringBuilder();
        int ans = 0;
        for(int i = 0; i < s.length(); i++) {
            char c = s.charAt(i);
            if(lookup.indexOf(c+"") != -1) {
                ans++;
                lookup = new StringBuilder();
            }
            lookup.append(c);
        }

        if(lookup.length() == s.length()) {
            return 1;
        }

        if(lookup.length() != 0) {
            ans += 1;
        }

        return ans;
    }
}