func dailyTemperatures(temperatures []int) []int {
    answer := make([]int, len(temperatures))

    for i := len(temperatures) - 2; i >= 0; i-- {
        j := i + 1
        
        for j < len(temperatures) && temperatures[i] >= temperatures[j] {
            if answer[j] <= 0 {
                break
            }
            j += answer[j]
        } 
        
        if j < len(temperatures) && temperatures[i] < temperatures[j] {
            answer[i] = j - i
        }
    }

    return answer
}
