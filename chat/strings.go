package chat

import (
	"hash/fnv"
	"math/rand"
)

const (
	separator = " \u2014 "
)

var (
	Whats = []string{"What?", "Nani?", "Huh?", "なんてこったい？", "Que?"}
)

func ToColor(s string) int {
	h := fnv.New32a()
	_, err := h.Write([]byte(s))
	if err != nil {
		return 0
	}
	return int(h.Sum32())
}

func RandomString(s []string) string {
	if len(s) == 1 {
		return s[0]
	}
	i := rand.Intn(len(s))
	return s[i]
}
