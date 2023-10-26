type LRUCache struct {
    capacity int
    cache map[int]*Node
    head *Node
    tail *Node
}

type Node struct {
    key int
    val int
    prev *Node
    next *Node
}

func Constructor(capacity int) LRUCache {
    LRU := LRUCache{
        capacity: capacity,
        cache: make(map[int]*Node),
        head: &Node{key: 0, val: 0},
        tail: &Node{key: 0, val: 0},
    }

    LRU.head.next, LRU.tail.prev = LRU.tail, LRU.head

    return LRU
}

func (this *LRUCache) Get(key int) int {
    if node, ok := this.cache[key]; ok {
        this.Remove(node)
        this.Insert(node)
        return node.val
    }
    return -1
}

func (this *LRUCache) Put(key int, value int)  {
    if node, ok := this.cache[key]; ok {
        this.Remove(node)
    }

    this.cache[key] = &Node{key: key, val: value}
    this.Insert(this.cache[key])
    
    if len(this.cache) > this.capacity {
        lru := this.tail.prev
        this.Remove(lru)
        delete(this.cache, lru.key)
    }
}

func (this *LRUCache) Remove(node *Node) {
    prev, next := node.prev, node.next
    prev.next, next.prev = next, prev
}

func (this *LRUCache) Insert(node *Node) {
    prev, next := this.head, this.head.next
    next.prev, prev.next = node, node
    node.next, node.prev = next, prev
}
