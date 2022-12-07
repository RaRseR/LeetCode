func findPeakElement(n []int) int {
    if len(n) == 1 {
        return 0
    }

    l := 0
    r := len(n) - 1

    for l < r {
        m := (l + r) / 2
        if n[m] > n[m + 1] {
            r = m
        } else {
            l = m + 1
        }
    }

    return l
}
