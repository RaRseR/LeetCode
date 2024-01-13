func isSubsequence(s string, t string) bool {
    if len(s) > len(t) {
        return false
    }

    if len(s) == 0 {
        return true
    }

    frstPntr := 0
    scndPntr := 0

    for frstPntr < len(s) && scndPntr < len(t) {
        if s[frstPntr] == t[scndPntr] {
            frstPntr++
        }

        if frstPntr == len(s) {
            return true
        }

        scndPntr++
    }

    return false
}
