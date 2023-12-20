func buyChoco(prices []int, money int) int {
    var firstChoco, secondChoco int = 101, 101

    for i := 0; i < len(prices); i++ {
        if firstChoco > prices[i] {
            firstChoco, secondChoco = prices[i], firstChoco
        } else if secondChoco > prices[i] {
            secondChoco = prices[i]
        }
    }

    if (firstChoco + secondChoco) <= money {
        return money - (firstChoco + secondChoco)
    }

    return money
}
