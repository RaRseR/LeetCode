class MyHashMap {
    buckets: number;
    entries: Array<Entry>;
    
    constructor() {
        this.buckets = 10_000;
        this.entries = [];
    }

    put(key: number, value: number): void {
        this.remove(key);
        const h = this.hash(key);
        
        this.entries[h] = new Entry(key, value, this.entries[h]);
    }

    get(key: number): number {
        const h = this.hash(key);
        let current = this.entries[h];

        while (current != null) {
            if (current.key == key) {
                return current.value;
            }

            current = current.next;
        }

        return -1;
    }

    remove(key: number): void {
        const h = this.hash(key);
        let current = this.entries[h];

        if (current == null) {
            return;
        }

        if (current.key == key) {
            this.entries[h] = current.next
            return;
        }

        while (current.next != null) {
            if (current.next.key == key) {
                current.next = current.next.next
                return;
            }

            current = current.next
        }
    }

    hash(key: number): number {
        return key % this.buckets
    }
}

class Entry{
    key: number;
    value: number;
    next: Entry;

    constructor(key, value, next){
        this.key = key;
        this.value = value;
        this.next = next;
    }
}
