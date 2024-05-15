class Solution {
    /***
     * 1846. Maximum Element After Decreasing and Rearranging
     */
    public int maximumElementAfterDecrementingAndRearranging(int[] arr) {
        Arrays.sort(arr);
        int res = 1;
        for (int i = 1; i < arr.length; ++i) {
            if (arr[i] > res) {
                ++res;
            }
        }
        return res;
    }
}