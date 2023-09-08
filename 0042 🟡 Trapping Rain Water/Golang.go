func trap(height []int) (result int) {
    if len(height) < 1 {
        return
    }

    left, right := 0, len(height) - 1
    leftMax, rightMax := height[left], height[right]

    for left < right {
        if leftMax < rightMax {
            left += 1

            if leftMax < height[left] {
                leftMax = height[left]
            }

            result += leftMax - height[left]
        } else {
            right -= 1

            if rightMax < height[right] {
                rightMax = height[right]
            }

            result += rightMax - height[right]
        }
    }

    return
}
