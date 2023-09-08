func maxArea(height []int) (result int) {
    left, right := 0, len(height) - 1

    for left < right {
        area := (right - left)

        if height[left] < height[right] {
            area *= height[left]
            left++
        } else {
            area *= height[right]
            right--
        }

        if area > result {
            result = area
        }
    }

    return
}
