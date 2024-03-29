function buildTree(preorder: number[], inorder: number[]): TreeNode | null {
    if (preorder.length == 0 || inorder.length == 0) {
        return null;
    }  

    const mid: number = inorder.indexOf(preorder[0]);

    return new TreeNode(
        preorder[0],
        buildTree(preorder.slice(1, mid + 1), inorder.slice(0, mid)),
        buildTree(preorder.slice(mid + 1), inorder.slice(mid + 1)),
    );
};
