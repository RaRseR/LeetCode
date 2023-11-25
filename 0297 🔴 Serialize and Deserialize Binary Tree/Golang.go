type Codec struct {
}

func Constructor() Codec {
    return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
    var result []string = make([]string, 0)

    var dfs func(node* TreeNode)
    dfs = func(node *TreeNode) {
        if node == nil {
            result = append(result, "N")
            return
        }
        
        result = append(result, strconv.Itoa(node.Val))
        
        dfs(node.Left)
        dfs(node.Right)
    }

    dfs(root);

    return strings.Join(result, ",")
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
    var vals []string = strings.Split(data, ",")
    var i int = 0

    var dfs func() *TreeNode
    dfs = func() *TreeNode {
        if vals[i] == "N" {
            i += 1
            return nil
        }

        val, _ := strconv.Atoi(vals[i])
        i += 1

        var node *TreeNode = &TreeNode{
            Val: val,
            Left: dfs(),
            Right: dfs(),
        }

        return node
    }

    return dfs()
}
