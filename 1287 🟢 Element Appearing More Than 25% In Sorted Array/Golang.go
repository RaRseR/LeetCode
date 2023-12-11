func findSpecialInteger(arr []int) int {
    if len(arr) == 1 {
        return arr[0]
    }

    var required int = len(arr) / 4

    for i := 0; i < len(arr) - required; i++ {
        if arr[i] == arr[i + required] {
            return arr[i]
        }
    }

    return -1
}
