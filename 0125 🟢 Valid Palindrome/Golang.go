func isPalindrome(s string) bool {
    pointer1, pointer2 := 0, len(s) - 1
    
    for pointer1 < pointer2 {
        left := unicode.ToLower(rune(s[pointer1]))
		    right := unicode.ToLower(rune(s[pointer2]))

        if !unicode.IsLetter(left) && !unicode.IsDigit(left) {
            pointer1++
        } else if !unicode.IsLetter(right) && !unicode.IsDigit(right) {
            pointer2--
        } else if left == right {
			      pointer1++
            pointer2--
		    } else {
            return false
        }
    }

    return true
}
