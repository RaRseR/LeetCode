type TimeMap struct {
  Store map[string][]Data
}

type Data struct {
	Value string
	Timestamp int
}


func Constructor() TimeMap {
    return TimeMap{
		Store: make(map[string][]Data),
	}
}


func (this *TimeMap) Set(key string, value string, timestamp int)  {
    if _, ok := this.Store[key]; !ok {
		this.Store[key] = []Data{}
	}

	this.Store[key] = append(this.Store[key], Data{
		Value: value,
		Timestamp: timestamp,
	})
}


func (this *TimeMap) Get(key string, timestamp int) string {
	var result string

	if values, ok := this.Store[key]; ok {
		var left int = 0
		var right int = len(values) - 1

		for left <= right {
			var middle int = (left + right) / 2

			if values[middle].Timestamp <= timestamp {
				result = values[middle].Value
				left = middle + 1
			} else {
				right = middle - 1
			}
		}
	}

	return result
}
