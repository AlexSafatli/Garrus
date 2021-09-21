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

func (t *LowercaseTrie) Add(s string) bool {
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
		return child.Add(s[1:])
	}
}

func (t *LowercaseTrie) GetNode(s string) *LowercaseTrie {
	var node = t
	for i := 0; i < len(s); i++ {
		index := sliceIndex(rune(s[0]), lowercaseAlphabet)
		if index == -1 {
			return nil // bad character
		}
		var child = node.children[index]
		node = &child
	}
	return node
}

func (t *LowercaseTrie) IsPrefix(s string) bool {
	var n = t.GetNode(s)
	return n != nil && n.numChildren > 0
}

func (t *LowercaseTrie) GetWordWithPrefix(s string) string {
	var out = s
	n := t.GetNode(s)
	if n == nil || n.numChildren <= 0 {
		return ""
	}
	for n.numChildren > 0 {
		var largest *LowercaseTrie
		var l rune
		for i := 0; i < len(lowercaseAlphabet); i++ {
			c := lowercaseAlphabet[i]
			child := n.children[i]
			if child.isWord {
				return out + string(c)
			}
			if largest == nil || largest.numChildren < child.numChildren {
				largest = &child
				l = c
			}
		}
		if largest == nil {
			return ""
		}
		n = largest
		out += string(l)
	}
	return ""
}

func (t *LowercaseTrie) Contains(s string) bool {
	n := t.GetNode(s)
	return n != nil && n.isWord
}

func (t *LowercaseTrie) HasChildren() bool {
	return t.numChildren > 0
}
