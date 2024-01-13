func minSteps(s string, t string) (result int) {
    count := [26]int{}

    for i := 0; i < len(s); i++ {
        count[s[i] - 'a']++
        count[t[i] - 'a']--
    }

    for i := 0; i < len(count); i++ {
        if count[i] > 0 {
            result += count[i]
        }
    }

    return
}
