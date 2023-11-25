class Solution:
    def buildTree(self, preorder: List[int], inorder: List[int]) -> Optional[TreeNode]:
        if len(preorder) == 0 or len(inorder) == 0:
            return None

        mid = inorder.index(preorder[0])

        return TreeNode(
            preorder[0],
            self.buildTree(preorder[1:mid + 1], inorder[:mid]),
            self.buildTree(preorder[mid + 1:], inorder[mid + 1:])
        )
        
