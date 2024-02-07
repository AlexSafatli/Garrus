package dice

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

type Roll struct {
	Num      int
	Sides    int
	Modifier int
	Exts     map[string][]string
}

var rng = rand.New(rand.NewSource(time.Now().Unix()))

type Formula string

func (r *Roll) Roll() (rolls []int) {
	for i := 0; i < r.Num; i++ {
		rolls = append(rolls, int(rng.Int31n(int32(r.Sides))+1))
	}
	return
}

func FromRoll(r Roll) (f Formula) {
	var count, modifier string
	if r.Num > 1 {
		count = strconv.Itoa(r.Num)
	}
	if r.Modifier != 0 {
		if r.Modifier > 0 {
			modifier = fmt.Sprintf("+%d", r.Modifier)
		} else {
			modifier = strconv.Itoa(r.Modifier)
		}
	}
	return Formula(fmt.Sprintf("%sd%d%s", count, r.Sides, modifier))
}

type Parser interface {
	Parse(f Formula) (Roll, error)
}

type parser struct {
}

func (p parser) Parse(f Formula) (_ Roll, err error) {
	return parse(f)
}

func (p parser) atoi(s string, def int) int {
	if len(s) == 0 {
		return def
	}
	i, err := strconv.Atoi(s)
	if err != nil {
		return def
	}
	return i
}

func New() Parser {
	return parser{}
}

func Parse(f Formula) (r Roll, err error) {
	p := New()
	r, err = p.Parse(f)
	return
}
