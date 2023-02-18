function maxDepth(root: TreeNode | null): number {
    if(!root) return 0;

    const maxL = maxDepth(root.left);
    const maxR = maxDepth(root.right);

    return Math.max(maxL, maxR) + 1;
};
