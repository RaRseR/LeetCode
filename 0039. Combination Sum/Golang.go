func combinationSum(candidates []int, target int) [][]int {
    sort.Ints(candidates)
    
    result := [][]int{}
    
    helper(candidates, target, []int{}, &result)
    
    return result
}


func helper(n []int, t int, p []int, r *[][]int){
    if t == 0 {
        c := make([]int, len(p))
        copy(c, p)
        
        *r = append(*r, c)
        return
    }
        
    for i, v := range n  {
        if v > t {
            return
        }
        
        helper(n[i:], t - v, append(p, v), r)
    }
}
