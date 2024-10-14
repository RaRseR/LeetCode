impl Solution {
    pub fn smallest_range(nums: Vec<Vec<i32>>) -> Vec<i32> {
        let k: usize = nums.len();

        let mut ordered_nums: Vec<(i32, usize)> = Vec::with_capacity(k);

        for i in 0..k {
            for n in &nums[i] {
                ordered_nums.push((*n, i));
            }
        }

        ordered_nums.sort_by_key(|x| x.0);

        let mut result: Vec<i32> = vec![ordered_nums[0].0, ordered_nums[ordered_nums.len() - 1].0];

        let mut count: Vec<usize> = vec![0; k];
        let mut unique: usize = 0;

        let mut left: usize = 0;

        for right in 0..ordered_nums.len() {
            count[ordered_nums[right].1] += 1;
            if count[ordered_nums[right].1] == 1 {
                unique += 1;
            }

            if unique == k {
                while left <= right && count[ordered_nums[left].1] > 1 {
                    count[ordered_nums[left].1] -= 1;
                    left += 1;
                }

                if ordered_nums[right].0 - ordered_nums[left].0 < result[1] - result[0] {
                    result[0] = ordered_nums[left].0;
                    result[1] = ordered_nums[right].0;
                }
            }
        }

        return result;
    }
}
