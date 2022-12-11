import "math"

func reverse(x int) (y int) {
    for x != 0 {
        y = y * 10 + x % 10
        x /= 10
    }

    if y < math.MinInt32 || y > math.MaxInt32 {
        return 0
    }

    return
}
