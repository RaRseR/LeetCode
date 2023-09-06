func twoSum(numbers []int, target int) []int {
    pointer1, pointer2 := 0, len(numbers) - 1
    
    for pointer1 < pointer2 {
        n := numbers[pointer1] + numbers[pointer2]
        if n == target {
            return []int{pointer1 + 1, pointer2 + 1}
        }
        
        if n > target {
            pointer2--
        } else {
            pointer1++
        }
    }
    
    return []int{0, 0}
}
