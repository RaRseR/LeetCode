func wordPattern(pattern string, s string) bool {
    words := strings.Fields(s)

    if len(words) != len(pattern) {
        return false
    }

    codes := make(map[string]byte)
    used := make(map[byte]bool, 26)

    for i := 0; i < len(pattern); i++ {
        if code, ok := codes[words[i]]; ok {
            if code != pattern[i] {
                return false
            }

            continue
        } else {
            if used[pattern[i]] {
                return false
            }

            used[pattern[i]] = true
            codes[words[i]] = pattern[i]
        }
    }

    return true
}
