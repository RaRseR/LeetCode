type WordDictionary struct {
    nodes   map[rune]*WordDictionary
    isEnd   bool
}

func Constructor() WordDictionary {
    return WordDictionary{make(map[rune]*WordDictionary), false}
}


func (this *WordDictionary) AddWord(word string)  {
    for _, r := range word {
        if _, ok := this.nodes[r]; !ok {
            node := Constructor()
            this.nodes[r] = &node
        }

        this = this.nodes[r]
    }

    this.isEnd = true
}


func (this *WordDictionary) Search(word string) bool {
    for i, r := range word {
        if _, ok := this.nodes[r]; !ok {
            if r != '.' {
                return false
            }

            for _, node := range this.nodes {
                if node.Search(word[i+1:]) {
                    return true
                }
            }

            return false
        }

        this = this.nodes[r]
    }

    return this.isEnd
}
