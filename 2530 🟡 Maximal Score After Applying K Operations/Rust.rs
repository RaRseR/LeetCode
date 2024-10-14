use std::collections::BinaryHeap;

impl Solution {
    pub fn max_kelements(nums: Vec<i32>, k: i32) -> i64 {
        let mut result: i64 = 0;

        let mut heap: BinaryHeap<i32> = BinaryHeap::from(nums);
        
        let mut value: i32 = 0;
        
        for i in 0..k {
            value = heap.pop().unwrap();

            result += value as i64;

            heap.push((value + 2) / 3);
        } 

        return result;
    }
}
