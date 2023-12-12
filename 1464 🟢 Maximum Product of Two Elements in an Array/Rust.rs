impl Solution {
    pub fn max_product(nums: Vec<i32>) -> i32 {
        let mut first_num: i32 = 0;
        let mut second_num: i32 = 0;

        for num in nums {
            if num > first_num {
                let swap: i32 = first_num;

                first_num = num;
                second_num = swap;
            } else if num > second_num {
                second_num = num
            }
        }

        return (first_num - 1) * (second_num - 1)
    }
}
