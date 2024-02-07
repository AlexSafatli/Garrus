package rpg

import (
	"github.com/AlexSafatli/Garrus/rpg/dice"
	"testing"
)

func TestRoll(t *testing.T) {
	f := dice.Formula("4d20+1d4+4")
	roll, err := Roll(f)
	if err != nil {
		t.Logf("unexpected error while rolling %s: %s", f, err)
	}

	// verify that the total was modified and within our expected range.
	if roll.Result[0] < 9 || roll.Result[0] > 88 {
		t.Logf("expected the total to be gt 9 and lt 88 with a roll of %s, instead received %d",
			string(f), roll.Result[0])
		t.Fail()
	}

	if t.Failed() {
		t.Logf("Roll: %#v", roll)
	}
}
