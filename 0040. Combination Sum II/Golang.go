func combinationSum2(candidates []int, target int) [][]int {
    sort.Ints(candidates)
    
    result := [][]int{}
    
    helper(candidates, target, []int{}, &result, 0)
    
    return result
}

func helper(n []int, t int, p []int, r *[][]int, j int){
    if t < 0 {
        return
    }
    
    if t == 0 {  
        *r = append(*r, append([]int{},  p...))
    }
    
    for i := j; i < len(n); i++ {

        if i > j && n[i] == n[i - 1] {
            continue
        }
        
        helper(n, t - n[i], append(p, n[i]), r, i + 1)
    }
}
