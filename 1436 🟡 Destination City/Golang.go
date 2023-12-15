func destCity(paths [][]string) string {
    var hashMap map[string]string = make(map[string]string)

    for _, path := range paths {
        hashMap[path[0]] = path[1]
    }

    for _, city := range hashMap {
        if _, exists := hashMap[city]; !exists {
            return city
        }
    }

    return ""
}
