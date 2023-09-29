func characterReplacement(s string, k int) (res int) {
    count := make(map[byte]int, 26)

    leftPointer := 0
    currentMax := 0

    for rightPointer, _ := range s {
        count[s[rightPointer]]++

        if count[s[rightPointer]] > currentMax {
            currentMax = count[s[rightPointer]]
        }

        if rightPointer - leftPointer + 1 - currentMax > k{
            count[s[leftPointer]]--
            leftPointer++
        }

        windowLength := rightPointer - leftPointer + 1

        if windowLength > res {
            res = windowLength
        }
    }
    
    return   
}
