/**
 * 
1485. Clone Binary Tree With Random Pointer
 */
class CopyRandomBinaryTree {
    public NodeCopy copyRandomBinaryTree(Node root) {
        if(root == null) {
            return null;
        }    
        Map<Node, NodeCopy> map = new HashMap<>();

        Stack<Node> stack = new Stack<>();
        Node cur = root;

        while(!stack.isEmpty() || cur != null) {
            while(cur != null) {
                stack.push(cur);
                cur = cur.left;
            }
            cur = stack.pop();
            NodeCopy copy = new NodeCopy(cur.val);
            map.put(cur, copy);
            cur = cur.right;
        }

        for(Node key: map.keySet()) {
            NodeCopy copy = map.get(key);

            copy.left = key.left == null ? null : map.get(key.left);
            copy.right = key.right == null ? null : map.get(key.right);
            copy.random = key.random == null ? null : map.get(key.random);
        }

        return map.get(root);
    }
}