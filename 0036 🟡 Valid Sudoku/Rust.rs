use std::collections::HashSet;

impl Solution {
    pub fn is_valid_sudoku(board: Vec<Vec<char>>) -> bool {
        let mut rows = HashSet::new();
        let mut cols = HashSet::new();
        let mut boxs = vec![HashSet::new(); 9];

        for i in 0..9 {
            for j in 0..9 {
                if let Some(digit) = board[i][j].to_digit(10) {
                    if !rows.insert((i, digit)) 
                    || !cols.insert((j, digit))
                    || !boxs[3*(i/3) + j/3].insert(digit)
                    {
                        return false;
                    }
                }
            }
        }

        true
    }
}
