func countNicePairs(nums []int) int {
    if len(nums) <= 1 {
        return 0
    }

    var hashMap map[int]int = make(map[int]int)
    var result int = 0
    var mod int = int(1e9 + 7)

    for i := 0; i < len(nums); i++ {
        var current int = nums[i] - reverseNum(nums[i])
        result = (result + hashMap[current]) % mod
        hashMap[current]++
    }

    return result
}

func reverseNum(num int) (result int) {
    for num != 0 {
        result = result * 10 + num % 10
        num /= 10
    }

    return
}
