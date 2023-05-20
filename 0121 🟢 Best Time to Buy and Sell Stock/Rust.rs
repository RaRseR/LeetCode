impl Solution {
    pub fn max_profit(prices: Vec<i32>) -> i32 {
        let (mut min, mut max) = (10000, 0);
        
        for price in prices {
            max = max.max(price - min);
            min = min.min(price);
        }
        
        max
    }
}
