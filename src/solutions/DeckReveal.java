public class DeckReveal {
    /**
     * 950. Reveal Cards In Increasing Order
     * 先排序，然后把后面的插入
     */
    public int[] deckRevealedIncreasing(int[] deck) {
        int N = deck.length;
        int[] result = new int[N];

        Arrays.sort(deck);

        return everyOther(deck, result, 0, 0, false);
    }

    private int[] everyOther(int[] deck, int[] result, int indexInDeck, int indexInResult, boolean skip) {
        int N = deck.length;

        if (indexInDeck == N) {
            return result;
        }

        while (indexInResult < N) {
            if (result[indexInResult] == 0) {
                if (!skip) {
                    result[indexInResult] = deck[indexInDeck];
                    indexInDeck++;
                }
                skip = !skip;
            }
            indexInResult++;
        }

        return everyOther(deck, result, indexInDeck, 0, skip);
    }
}