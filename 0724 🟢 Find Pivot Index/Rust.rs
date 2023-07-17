pub fn pivot_index(nums: Vec<i32>) -> i32 {
  let total: i32 = nums.iter().sum();

  let mut left: i32 = 0;
  let mut i = 0;

  while i < nums.len() {
    if left * 2 == total - nums[i] {
      return i as i32;
    }
    
    left += nums[i];
    i += 1;
  }

  return -1;
}
