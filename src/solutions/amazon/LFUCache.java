package solutions.amazon;

/**
 * #460 LFU Cache
 */
class LFUCache {

        private Node head = null;
    private Node tail = null;

    private int capacity = 0;

    private final HashMap<Integer, Integer> valueHash;
    private final HashMap<Integer, Node> nodeHash;

    class Node {
        public int count = 0;
        public LinkedHashSet<Integer> keys;
        public Node prev = null;
        public Node next = null;

        public Node(int count) {
            this.count = count;
            keys = new LinkedHashSet<>();
            prev = next = null;
        }
    }
    
    public LFUCache(int capacity) {
        valueHash = new HashMap<>();
        nodeHash = new HashMap<>();

        this.capacity = capacity;

        head = new Node(-1);
        tail = new Node(-1);

        head.next = tail;
        tail.prev = head;
    }

    public int get(int key) {
        if(valueHash.containsKey(key)) {
            increaseCount(key);
            return valueHash.get(key);
        }

        return -1;
    }

    public void put(int key, int value) {
        if(capacity == 0) {
            return;
        }

        if(valueHash.containsKey(key)) {
            valueHash.put(key, value);
            increaseCount(key);
        } else {
            valueHash.put(key, value);
            if(valueHash.size() > capacity) {
                removeHead();
            }
            addToHead(key);
        }
    }

    private void addToHead(int key) {
        if(head.next.count == 1) {
            Node node = head.next;
            nodeHash.put(key, node);
            node.keys.add(key);
        } else {
            Node node = new Node(1);
            node.keys.add(key);

            node.next = head.next;
            node.next.prev = node;
            head.next = node;
            node.prev = head;

            nodeHash.put(key, node);
        }
    }

    private void removeHead() {
        //printLFU();
        Node least = head.next;
        int key = least.keys.iterator().next();
        least.keys.remove(key);

        if(least.keys.size() == 0) {
            remove(least);
        }

        nodeHash.remove(key);
        valueHash.remove(key);
    }

    private void increaseCount(int key) {
        Node node = nodeHash.get(key);
        node.keys.remove(key);

        if(node.next == tail) {
            node.next = new Node(node.count + 1);
            node.next.keys.add(key);
            node.next.prev = node;
            node.next.next = tail;
        } else if(node.next.count == node.count + 1) {
            node.next.keys.add(key);
        } else {
            Node next = new Node(node.count + 1);
            next.keys.add(key);
            next.next = node.next;
            node.next.prev = next;
            node.next = next;
            next.prev = node;
        }

        nodeHash.put(key, node.next);
        if(node.keys.size() == 0) {
            remove(node);
        }
    }

    private void remove(Node node) {
        Node prev = node.prev;
        Node next = node.next;

        prev.next = next;
        next.prev = prev;
    }
}