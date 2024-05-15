class FairCandySwap {
    /**
     * 888 Fair Candy Swap
     */
    public int[] fairCandySwap(int[] aliceSizes, int[] bobSizes) {
        int sum1 = IntStream.of(aliceSizes).sum();
        int sum2 = IntStream.of(bobSizes).sum();
        int diff=sum2-sum1;
        HashSet<Integer> set = new HashSet<>();
        int[] res = new int[2];
        for(int i=0;i<bobSizes.length;i++){
            set.add(bobSizes[i]);
        }
        for (int i = 0; i < aliceSizes.length; i++) {
            int x = aliceSizes[i];
           int y=(diff+2*x)/2;
           if(set.contains(y)){
            res[0]=x;
            res[1]=y;
           }
            
        }
        return res;
    }
}