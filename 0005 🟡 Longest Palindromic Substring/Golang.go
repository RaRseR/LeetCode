func longestPalindrome(s string) string {
    if len(s) < 2 {
        return s
    }

    expandAroundCenter := func(left, right int) string {
        for left >= 0 && right < len(s) && s[left] == s[right] {
            left--
            right++
        }

        return s[left + 1:right]
    }

    result := ""

    for i := 0; i < len(s); i++ {
        pal := expandAroundCenter(i, i)

        if len(pal) > len(result) {
            result = pal
        }

        pal = expandAroundCenter(i, i + 1)

        if len(pal) > len(result) {
            result = pal
        }
    }

    return result
}
