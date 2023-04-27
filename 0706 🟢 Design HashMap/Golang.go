const buckets = 10_000

type MyHashMap struct {
    entries [buckets]*entry
}

type entry struct {
    Key int
    Value int
    Next *entry
}

func Constructor() MyHashMap  {
    return MyHashMap{[buckets]*entry{}}
}

func (this *MyHashMap ) Put(key int, value int)  {
    this.Remove(key)
    h := this.Hash(key)
    
    entry := &entry{
        Key: key,
        Value: value,
        Next: this.entries[h],
    }
    this.entries[h] = entry
}

func (this *MyHashMap ) Get(key int) int {
    h := this.Hash(key)
    curr := this.entries[h]

    for curr != nil {
        if curr.Key == key {
            return curr.Value
        }
        curr = curr.Next
    }

    return -1
}


func (this *MyHashMap ) Remove(key int)  {
    h := this.Hash(key)
    curr := this.entries[h]

    if curr == nil {
        return
    }

    if curr.Key == key {
        this.entries[h] = curr.Next
        return
    }

    for curr.Next != nil {
        if curr.Next.Key == key {
            curr.Next = curr.Next.Next
            return
        }
        curr = curr.Next
    }
}

func (this *MyHashMap) Hash(key int) int {
    return key % buckets
}
