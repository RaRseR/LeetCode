function kthSmallest(root: TreeNode | null, k: number): number {
    const stack: Array<TreeNode> = new Array(k);

    while (true) {
        while (root !== null) {
            stack.push(root);
            root = root.left;
        }

        root = stack.pop();

        k--

        if (k == 0) {
            return root.val;
        }

        root = root.right;
    }
};
