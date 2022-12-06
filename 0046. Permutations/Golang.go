func permute(nums []int) (result [][]int) {
    sort.Ints(nums)
    
    helper(nums, []int{}, &result)

    return result
}

func helper(n []int, p []int, r *[][]int){
    if (len(n) == 0){
        *r = append(*r, p)
        return
    } 
        
    for i, v := range n {
        helper(append(append([]int{}, n[:i]...), n[i + 1:]...), append(p, v), r)
    }   
}
