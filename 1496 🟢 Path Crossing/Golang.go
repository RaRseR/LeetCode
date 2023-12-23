func isPathCrossing(path string) bool {
    var visitedPaths map[string]bool = make(map[string]bool)
    visitedPaths["0,0"] = true

    var x, y int = 0, 0

    for _, char := range path {
        if char == 'N' {
            y++
        } else if char == 'S' {
            y--
        } else if char == 'E' {
            x++
        } else {
            x--
        }

        var currentLocation string = fmt.Sprintf("%d,%d", x, y)

        if visitedPaths[currentLocation] {
            return true
        }

        visitedPaths[currentLocation] = true
    }

    return false
}
