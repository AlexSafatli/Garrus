package chat

import (
	"hash/fnv"
	"math/rand"
	"strings"
)

const (
	Separator        = " \u2014 "
	maxMessageLength = 2000
)

var (
	Whats = []string{"What?", "Nani?", "Huh?", "なんてこったい？", "Que?"}
)

type MessageBuilder struct {
	msgStrings []string
	*strings.Builder
}

func (b *MessageBuilder) Write(s string) error {
	if b.Len()+len(s) > maxMessageLength {
		b.msgStrings = append(b.msgStrings, b.String())
		b.Builder.Reset()
	}
	_, err := b.WriteString(s)
	if err != nil {
		return err
	}
	return nil
}

func (b *MessageBuilder) GetMessageStrings() []string {
	if b.Len() > 0 {
		b.msgStrings = append(b.msgStrings, b.String())
		b.Builder.Reset()
	}
	return b.msgStrings
}

func (b *MessageBuilder) Reset() {
	b.msgStrings = make([]string, 0)
	b.Builder.Reset()
}

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
