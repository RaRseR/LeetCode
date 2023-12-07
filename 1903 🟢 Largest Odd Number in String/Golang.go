func largestOddNumber(num string) string {
    var i int = len(num) - 1

    for i >= 0 && num[i] & 1 == 0 {
        i--
    }

    return num[:i + 1]
}
