class Solution {

    static double helper(int arr1[], int arr2[], int arr3[]){
        double x1 = (double)arr1[0];
        double y1 = (double)arr1[1];

        double x2 = (double)arr2[0];
        double y2 = (double)arr2[1];

        double x3 = (double)arr3[0];
        double y3 = (double)arr3[1];

        return Math.abs((x1*(y2-y3) + x2*(y3-y1) + x3*(y1-y2)))/2;
    }
    public double largestTriangleArea(int[][] arr) {
        

        double ans = 0;

        for(int i=0; i<arr.length; i++){
            for(int j=i+1; j<arr.length; j++){
                for(int k=j+1; k<arr.length; k++){
                    ans = Math.max(ans,helper(arr[i],arr[j],arr[k]));
                }
            }
        }
        return ans;
    }
}