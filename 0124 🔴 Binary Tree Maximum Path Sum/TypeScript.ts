function maxPathSum(root: TreeNode | null): number {
	let result: number = Number.MIN_SAFE_INTEGER;

	const dfs = (root: TreeNode | null): number => {
		if (root === null) {
			return 0;
		}

		const left: number = Math.max(0, dfs(root.left));
		const right: number = Math.max(0, dfs(root.right));

		result = Math.max(result, root.val + left + right);

		return root.val + Math.max(left, right);
	}

	dfs(root);

	return result;
};
