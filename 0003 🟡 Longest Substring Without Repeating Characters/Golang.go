func lengthOfLongestSubstring(s string) (l int) {
    if len(s) == 0 {
        return
    }

    hash := map[byte]int{}

    for j, i := 0, 0; j < len(s); j++ {
        if v, ok := hash[s[j]]; ok && v > i {
            i = v
        }
        if j - i + 1 > l {
            l = j - i + 1
        }
        hash[s[j]] = j + 1
    }

    return
}
