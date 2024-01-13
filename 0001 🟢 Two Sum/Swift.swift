class Solution {
    func twoSum(_ nums: [Int], _ target: Int) -> [Int] {
        var map: [Int: Int] = [:]

        for i in 0..<nums.count {
            if let index = map[nums[i]] {
                return [index, i]
            }

            map[target - nums[i]] = i
        }

        return [-1, -1]
    }
}
