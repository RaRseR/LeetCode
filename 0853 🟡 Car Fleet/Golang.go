func carFleet(target int, position []int, speed []int) int {
    var cars [][2]int = make([][2]int, len(position)) 
    
    for i := range cars {
        cars[i] = [2]int{
            position[i],
            speed[i],
        }
    }

    sort.Slice(cars, func(i, j int) bool {
        return cars[i][0] > cars[j][0]
    })

    var time []float64 = make([]float64, 0, len(cars))

    for i := range cars {
        var currentTime float64 = float64(target - cars[i][0]) / float64(cars[i][1])

        if len(time) == 0 || currentTime > time[len(time) - 1] {
            time = append(time, currentTime)
        }
    }

    return len(time)
}
