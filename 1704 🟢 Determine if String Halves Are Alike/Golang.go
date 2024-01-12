func halvesAreAlike(s string) bool {
    count := 0
    mid := len(s) / 2

    for i := 0; i < len(s); i++ {
        switch rune(s[i]) {
            case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
                if i < mid {
                    count++
                } else {
                    count--
                }
        }
    }

    return count == 0
}
