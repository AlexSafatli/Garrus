package trie

var (
	lowercaseAlphabet = []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k',
		'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}
)

func toInt(str string, alphabet []rune) []int {
	r := make([]int, len(str))
	for i := 0; i < len(r); i++ {
		r[i] = sliceIndex(rune(str[i]), alphabet)
	}
	return r
}

func sliceIndex(ele rune, alphabet []rune) int {
	for k, v := range alphabet {
		if ele == v {
			return k
		}
	}
	return -1
}
