func canCompleteCircuit(gas []int, cost []int) int {
    var fuel int = 0
    var allFuel int = 0
    var start int = 0

    for i := 0; i < len(gas); i++ {
        allFuel += gas[i] - cost[i]
        fuel += gas[i] - cost[i]

        if fuel < 0 {
            start = i + 1
            fuel = 0
        }
    }

    if allFuel < 0 {
        return -1
    }

    return start
}
