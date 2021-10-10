package chat

import (
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
	initialized bool
	msgStrings  []string
	strings.Builder
}

func NewMessageBuilder() *MessageBuilder {
	return &MessageBuilder{}
}

func (b *MessageBuilder) Write(s string) error {
	if b.initialized && b.Len()+len(s) > maxMessageLength {
		b.msgStrings = append(b.msgStrings, b.String())
		b.Builder = strings.Builder{}
		b.initialized = false
	} else {
		b.initialized = true
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

func RandomString(s []string) string {
	if len(s) == 1 {
		return s[0]
	}
	i := rand.Intn(len(s))
	return s[i]
}

func SliceToMessageString(slice []string) (str string) {
	for i, v := range slice {
		str += "`" + v + "`"
		if i < len(slice)-1 {
			str += ", "
		}
	}
	return
}
