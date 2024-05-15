class PalindromicPath {
    /**
     * 1457. Pseudo-Palindromic Paths in a Binary Tree
     */
    public int pseudoPalindromicPaths (TreeNode root) {
        res = 0;
        int arr[] = new int[10];

        countPalindrome(root, arr);
        return res;
    }

    public void countPalindrome(TreeNode root , int arr[]) {
        if(root == null) {
            return ; 
        }
        arr[root.val]++;
        if(root.left == null && root.right == null) {
            int count = 0 ; 
            boolean flag = true; 
            for(int i : arr) {
                if(i % 2  != 0) {
                    count++; 
                }
                if(count > 1 ) {
                    flag =false ; 
                    break ; 
                }
               
            }
            if(flag) {
                res++; 
            }
        }
        countPalindrome(root.left , arr);
        countPalindrome(root.right , arr); 
        arr[root.val]--; 
        return ; 
    }
}