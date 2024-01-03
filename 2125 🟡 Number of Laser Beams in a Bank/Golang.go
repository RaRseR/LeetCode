func numberOfBeams(bank []string) (result int) {
    var previous int = 0
    var count int = 0

    for _, row := range bank {
        count = 0;

        for _, char := range row {
            if char == 49 {
                count++
            }
        }
        
        if count > 0 {
            result += previous * count
            previous = count
        }
    }
    
    return
}
