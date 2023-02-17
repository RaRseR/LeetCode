function minDiffInBST(root: TreeNode | null): number {
    let min = 100_000;
    let prev = -1;

    const dfs = (node: TreeNode | null) => {
        if (node == null) {
            return;
        }

        dfs(node.left);

        if (prev !== -1 && node.val - prev < min){
            min = node.val - prev;
        }

        prev = node.val;

        dfs(node.right);
    }

    dfs(root);

    return min;
};
