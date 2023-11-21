impl Solution {
    pub fn search_insert(nums: Vec<i32>, target: i32) -> i32 {
        let mut left: usize = 0;
        let mut right: usize = nums.len();

        while left < right {
            let middle: usize = left + (right - left) / 2;

            if nums[middle] == target {
                return middle as i32;
            }
            if target < nums[middle] {
                right = middle;
            } else {
                left = middle + 1;
            }
        }

        return left as i32;
    }   
}
