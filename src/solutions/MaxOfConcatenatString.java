class Solution {
    int max = 0;
    public int maxLength(List<String> arr) {
        backTrack(arr, "", 0);
        return max;
    }

    private void backTrack(List<String> arr, String current, int start) {
        if(max < current.length()) {
            max = current.length();
        }

        for(int i = start; i < arr.size(); i++) {
            if(!isValid(current, arr.get(i))) {
                continue;
            }
            backTrack(arr, current+arr.get(i), i+1);
        }
    }

    private boolean isValid(String current, String newString) {
        Set<Character> set = new HashSet<>();

        for(char ch : newString.toCharArray()) {
            if(set.contains(ch)) {
                return false;
            }

            set.add(ch);

            if(current.contains(String.valueOf(ch))) {
                return false;
            }
        }

        return true;
    }
}