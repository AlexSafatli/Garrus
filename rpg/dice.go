package rpg

import (
	"github.com/AlexSafatli/Garrus/rpg/dice"
)

type DiceResult struct {
	dice.Formula `json:"formula"`
	dice.Roll    `json:"roll"`
	Result       []int `json:"result"`
}

func Roll(f dice.Formula) (d DiceResult, err error) {
	roll, err := dice.Parse(f)
	if err != nil {
		return
	}
	d = DiceResult{
		Formula: f,
		Roll:    roll,
		Result:  roll.Roll(),
	}
	return
}
