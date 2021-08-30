package structs

type LowercaseTrie struct {
	isWord      bool
	children    []LowercaseTrie
	numChildren int
}

func NewLowercaseTrie(set []string) LowercaseTrie {
	var trie LowercaseTrie
	trie.children = make([]LowercaseTrie, len(lowercaseAlphabet))
	return trie
}

func (t *LowercaseTrie) add(s string) bool {
	first := rune(s[0])
	index := sliceIndex(first, lowercaseAlphabet)
	if index < 0 {
		return false
	}
	var child = &t.children[index]
	if child == nil {
		child = &LowercaseTrie{}
		t.children[index] = *child
		t.numChildren++
	}
	if len(s) == 1 {
		if child.isWord {
			// Word already in trie
			return false
		}
		child.isWord = true
		return true
	} else {
		// Recurse into sub-trie
		return child.add(s[1:])
	}
}
