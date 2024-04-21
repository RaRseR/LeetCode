func searchRange(nums []int, target int) []int {
    result := []int{-1, -1}

    if (len(nums) == 0) {
        return result
    }

    left := binarySearch(&nums, target, true)

    if left != -1 {
        result[0] = left
        result[1] = binarySearch(&nums, target, false)
    }

    return result
}

func binarySearch(nums *[]int, target int, isLeft bool) int {
    result := -1

    left, mid, right := 0, 0, len(*nums)

    for left < right {
        mid = left + (right - left) / 2

        if ((*nums)[mid] == target) {
            if result == -1 {
                result = mid
            } else {
                if isLeft {
                    if mid < result {
                        result = mid
                    }
                } else {
                    if mid > result {
                        result = mid
                    }
                }
            }

            if isLeft {
                right = mid
            } else {
                left = mid + 1
            }
        } else if (*nums)[mid] > target {
            right = mid
        } else {
            left = mid + 1
        }
    }

    return result
}
