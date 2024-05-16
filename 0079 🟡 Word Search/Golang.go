func exist(board [][]byte, word string) bool {
    for row := 0; row < len(board); row++ {
        for col := 0; col < len(board[0]); col++ {
            if board[row][col] == word[0] && helper(board, row, col, 0, word) {
                return true
            }
        }
    }

    return false
}

func helper(board [][]byte, row, col, pos int, word string) bool {
    if row < 0 || row >= len(board) || col < 0 || col >= len(board[0]) || board[row][col] == '.' {
        return false
    }

    if word[pos] != board[row][col] {
        return false
    }

    if pos == len(word) - 1 {
        return true
    }

    pos++
    tmp := board[row][col]
    board[row][col] = '.'

    result := helper(board, row - 1, col, pos, word) ||
        helper(board, row + 1, col, pos, word) ||
        helper(board, row, col - 1, pos, word) ||
        helper(board, row, col + 1, pos, word)

    board[row][col] = tmp

    return result
}
