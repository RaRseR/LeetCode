func numberOfArithmeticSlices(nums []int) int {
    res := 0
    l := len(nums)
    m := make(map[int]map[int]int, l)
    
    for i, a := range nums {
        m[i] = make(map[int]int, i)
        
        for j, b := range nums[:i]  {
            diff := a - b
            
            res += m[j][diff]
            
			      m[i][diff] += m[j][diff] + 1
        }
    }
	
	
	return res
}
