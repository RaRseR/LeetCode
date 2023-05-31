impl Solution {
    pub fn maximum_wealth(accounts: Vec<Vec<i32>>) -> i32 {
        let mut result: i32 = 0;

        for account in accounts.iter() {
            let mut wealth: i32 = 0;

            for money in account.iter() {
                wealth += money;
            }
            
            if (result < wealth) {
                result = wealth;
            }
        }

        result
    }
}
