func minEatingSpeed(piles []int, h int) int {
    left, right := 1, piles[0]

    for _, val := range(piles) {
        if right < val {
            right = val
        }
    }

    for left < right {
        middle := (left + right) / 2
        hours := 0

        for _, pile := range(piles) {
            hours += int(math.Ceil(float64(pile) / float64(middle)))
        }

        if hours <= h {
            right = middle
        } else {
            left = middle + 1
        }
    }

    return right
}
