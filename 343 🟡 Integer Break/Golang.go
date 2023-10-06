func integerBreak(n int) int {
    if n <= 3 {
        return n-1
    } 
    
    if n % 3 == 0 {
        return power(3, n / 3)
    } 
    
    if n % 3 == 1 {
        return power(3, ((n / 3) - 1)) * 4
    }
    
    if n % 3 == 2 {
        return power(3, (n / 3)) * 2
    }

    return -1
}

func power(base, exponent int) int {
	result := 1
	for i := 0; i < exponent; i++ {
		result *= base
	}
	return result
}
