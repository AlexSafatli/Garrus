package dice

import (
	"math/rand"
	"time"
)

var rng = rand.New(rand.NewSource(time.Now().Unix()))

type Formula string

type Parser interface {
	Parse(f Formula) (Rollable, error)
}

type parser struct {
}

func (p parser) Parse(f Formula) (_ Rollable, err error) {
	return parse(f)
}

func New() Parser {
	return parser{}
}

func Parse(f Formula) (r Rollable, err error) {
	p := New()
	r, err = p.Parse(f)
	return
}
