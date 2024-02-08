package dice

import "fmt"

type Rollable interface {
	String() string
	Roll() int
}

type RollableDice struct {
	Num      int
	Sides    int
	Modifier int
}

func (r RollableDice) String() string {
	return fmt.Sprintf("%dd%d", r.Num, r.Sides)
}

func (r RollableDice) Roll() (result int) {
	result = r.Modifier
	if r.Num == 0 || r.Sides == 0 {
		return
	}
	for i := 0; i < r.Num; i++ {
		result += rng.Intn(r.Sides-1) + 1
	}
	return
}
