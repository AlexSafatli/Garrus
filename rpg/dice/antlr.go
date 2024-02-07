package dice

import (
	"fmt"
	fp "github.com/AlexSafatli/Garrus/rpg/dice/parser"
	"github.com/antlr4-go/antlr/v4"
	"github.com/golang-collections/collections/stack"
	"github.com/pkg/errors"
)

type ListenerError struct {
	antlr.ErrorListener
}

// DiceParserListener implements the BaseDiceListener
type DiceParserListener struct {
	roll  Roll
	nodes *stack.Stack
	fp.BaseDiceListener
}

func (s *DiceParserListener) Roll() Roll { return s.roll }

func (s *DiceParserListener) ExitAdd(ctx *fp.AddContext) {
	var operators []string
	for _, operator := range ctx.AllADDOPERATOR() {
		operators = append(operators, operator.GetText())
	}

}

func binaryOperation(operators []string) string {
	var operands *stack.Stack
	operands = stack.New()

	for _, op := range operators {
		operands.Push()
	}
}

func parse(f Formula) (_ Roll, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = errors.Wrap(r.(error), fmt.Sprintf("failed to parse formula \"%s\"", f))
		}
	}()

	input := antlr.NewInputStream(string(f))
	lexer := fp.NewDiceLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	listener := new(DiceParserListener)

	dp := fp.NewDiceParser(stream)
	dp.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	dp.BuildParseTrees = true

	antlr.ParseTreeWalkerDefault.Walk(listener, dp.Formula())
	return listener.Roll(), err
}
