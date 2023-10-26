type ListNodeItem = {
    key: number,
    value: number,
    previous: ListNodeItem | null,
    next: ListNodeItem | null
}

class LRUCache {
    capacity: number;
    cache: Map<number, ListNodeItem>;
    head: ListNodeItem;
    tail: ListNodeItem;

    constructor(capacity: number) {
        this.capacity = capacity;
        this.cache = new Map<number, ListNodeItem>();
        this.head = {key: 0, value: 0, previous: null, next: null};
        this.tail = {key: 0, value: 0, previous: null, next: null};

        this.head.next = this.tail;
        this.tail.previous = this.head;
    }

    get(key: number): number {
        if (!this.cache.has(key)) {
            return -1;
        }

        const node = this.cache.get(key);

        this.remove(node);
        this.insert(node);

        return node.value;
    }

    put(key: number, value: number): void {
        if (this.cache.has(key)) {
            this.remove(this.cache.get(key));
        }

        const node = {key, value, previous: null, next: null};

        this.cache.set(key, node);
        this.insert(node);

        if (this.cache.size > this.capacity) {
            const lru = this.tail.previous;

            this.remove(lru);

            this.cache.delete(lru.key);
        }
    }

    remove(node: ListNodeItem): void {
        const previous = node.previous;
        const next = node.next;

        next.previous = previous;
        previous.next = next;
    }

    insert(node: ListNodeItem): void {
        const previous = this.head;
        const next = this.head.next;

        next.previous = node;
        previous.next = node;

        node.previous = previous;
        node.next = next;
    }
}
