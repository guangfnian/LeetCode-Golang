package p0208

type Trie struct {
	Next [27]*Trie
	End  bool
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	for _, w := range word {
		if this.Next[w-'a'] == nil {
			this.Next[w-'a'] = &Trie{}
		}
		this = this.Next[w-'a']
	}
	this.End = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	for _, w := range word {
		this = this.Next[w-'a']
		if this == nil {
			return false
		}
	}
	return this.End
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	for _, w := range prefix {
		this = this.Next[w-'a']
		if this == nil {
			return false
		}
	}
	return true
}
