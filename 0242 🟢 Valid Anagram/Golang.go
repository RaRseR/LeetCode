func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }

    var chars []int16 = make([]int16, 26)

    for i := 0; i < len(s); i++ {
        chars[s[i] - 97]++
        chars[t[i] - 97]--
    }

    for i := 0; i < len(chars); i++ {
        if chars[i] != 0 {
            return false
        }
    }

    return true
}
