package rpg

import (
	"github.com/AlexSafatli/Garrus/rpg/dice"
	"testing"
)

func TestRoll(t *testing.T) {
	f := dice.Formula("4d20 + 2")
	roll, err := Roll(f)
	if err != nil {
		t.Logf("unexpected error while rolling %s: %s", f, err)
	}

	// verify that the total was modified and within our expected range.
	if roll.Result < 6 || roll.Result > 84 {
		t.Logf("expected the total to be gt 6 and lt 84 with a roll of %s, instead received %d",
			string(f), roll.Result)
		t.Fail()
	}

	if t.Failed() {
		t.Logf("Roll: %#v", roll)
	}
}
