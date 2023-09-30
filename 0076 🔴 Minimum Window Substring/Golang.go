func minWindow(s string, t string) string {
    if len(t) > len(s) {
        return ""
    }

    targetCharacterFrequency := make(map[uint8]int)
    currentCharacterFrequency := make(map[uint8]int)

    for index := range t {
		targetCharacterFrequency[t[index]]++
    }

    leftPointer, rightPointer := 0, 0
    distinctCharacterCount := 0
    result := ""

    for rightPointer < len(s) {
        currentCharacterFrequency[s[rightPointer]]++
        if targetCharacterFrequency[s[rightPointer]] != 0 &&
	    	targetCharacterFrequency[s[rightPointer]] == currentCharacterFrequency[s[rightPointer]] {
				distinctCharacterCount++
	    }

        for distinctCharacterCount == len(targetCharacterFrequency) {
            if result == "" || rightPointer - leftPointer + 1 < len(result) {
                result = s[leftPointer:rightPointer + 1]
            }
            
			currentCharacterFrequency[s[leftPointer]]--
			if currentCharacterFrequency[s[leftPointer]] < targetCharacterFrequency[s[leftPointer]] {
				distinctCharacterCount--
			}
			
			leftPointer++
        }
        
        rightPointer++
    }

    return result
}
