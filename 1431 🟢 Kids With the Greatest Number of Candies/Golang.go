func kidsWithCandies(candies []int, extraCandies int) []bool {    
    res := make([]bool, len(candies))
    
    biggest := candies[0]
    
    for _, candy := range candies[1:] {
        if (candy > biggest) {
            biggest = candy
        }
    }
    
    for i, candy := range candies {
        if (candy + extraCandies >= biggest) {
            res[i] = true
        } else {
            res[i] = false
        }
    }
    
    return res
}
