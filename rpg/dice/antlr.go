package dice

import (
	"fmt"
	fp "github.com/AlexSafatli/Garrus/rpg/dice/parser"
	"github.com/antlr4-go/antlr/v4"
	"github.com/pkg/errors"
	"strconv"
)

type ListenerError struct {
	antlr.ErrorListener
}

type DiceParserListener struct {
	roll RollableDice
	fp.BaseDiceListener
}

func (s *DiceParserListener) EnterNotation(_ *fp.NotationContext) {
	s.roll = RollableDice{
		Num:   1,
		Sides: 1,
	}
}

func (s *DiceParserListener) EnterCount(ctx *fp.CountContext) {
	i, err := strconv.Atoi(ctx.GetText())
	if err != nil {
		panic(fmt.Sprintf("failed to cast Count to integer: %s", err))
	}
	s.roll.Num = i
}

func (s *DiceParserListener) EnterSides(ctx *fp.SidesContext) {
	i, err := strconv.Atoi(ctx.GetText()[1:])
	if err != nil {
		panic(fmt.Sprintf("failed to cast Sides to integer: %s", err))
	}
	s.roll.Sides = i
}

func (s *DiceParserListener) EnterModifier(ctx *fp.ModifierContext) {
	var n int
	n = 1
	if ctx.SIGN().GetText() == "-" {
		n = -1
	}

	str := ctx.Integer().GetText()
	i, err := strconv.Atoi(str)
	if err != nil {
		panic(fmt.Sprintf("failed to cast Modifier to integer: %s", err))
	}
	s.roll.Modifier = i * n
}

func parse(f Formula) (_ Rollable, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = errors.Wrap(errors.New(r.(string)), fmt.Sprintf("failed to parse formula \"%s\"", f))
		}
	}()

	input := antlr.NewInputStream(string(f))
	lexer := fp.NewDiceLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	listener := new(DiceParserListener)

	dp := fp.NewDiceParser(stream)
	dp.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	dp.BuildParseTrees = true

	antlr.ParseTreeWalkerDefault.Walk(listener, dp.Notation())
	return listener.roll, err
}
