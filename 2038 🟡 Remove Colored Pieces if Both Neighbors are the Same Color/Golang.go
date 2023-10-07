func winnerOfGame(colors string) bool {
    countA := 0
    countB := 0

    for i := 0; i < len(colors); i++ {
        c := colors[i]
        count := 0

        for i < len(colors) && colors[i] == c {
            i++
            count++
        }

        current := count - 2

        if current > 0 {
            if c == 'A' {
                countA += current
            } else {
                countB += current
            }
        }

        i--
    }

    return countA > countB
}
