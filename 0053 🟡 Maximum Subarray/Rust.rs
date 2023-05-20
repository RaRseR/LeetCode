impl Solution {
    pub fn max_sub_array(nums: Vec<i32>) -> i32 {
        let (mut local, mut result) = (nums[0], nums[0]);

        for i in 1..nums.len() {
            if local < 0 {
                local = nums[i];
            } else {
                local += nums[i];
            }

            result = result.max(local);
        }

        result
    }
}
