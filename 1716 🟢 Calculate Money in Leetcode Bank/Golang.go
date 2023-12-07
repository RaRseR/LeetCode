func totalMoney(n int) int {
    var weeks int = n / 7
    var days int = n % 7

    return 7 * ((weeks - 1) * weeks) / 2 + 28 * weeks + (days + days * days) / 2 + weeks * days
}
