// Code generated from Dice.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // Dice

import "github.com/antlr4-go/antlr/v4"

// BaseDiceListener is a complete listener for a parse tree produced by DiceParser.
type BaseDiceListener struct{}

var _ DiceListener = &BaseDiceListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseDiceListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseDiceListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseDiceListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseDiceListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterNotation is called when production notation is entered.
func (s *BaseDiceListener) EnterNotation(ctx *NotationContext) {}

// ExitNotation is called when production notation is exited.
func (s *BaseDiceListener) ExitNotation(ctx *NotationContext) {}

// EnterCount is called when production count is entered.
func (s *BaseDiceListener) EnterCount(ctx *CountContext) {}

// ExitCount is called when production count is exited.
func (s *BaseDiceListener) ExitCount(ctx *CountContext) {}

// EnterSides is called when production sides is entered.
func (s *BaseDiceListener) EnterSides(ctx *SidesContext) {}

// ExitSides is called when production sides is exited.
func (s *BaseDiceListener) ExitSides(ctx *SidesContext) {}

// EnterModifier is called when production modifier is entered.
func (s *BaseDiceListener) EnterModifier(ctx *ModifierContext) {}

// ExitModifier is called when production modifier is exited.
func (s *BaseDiceListener) ExitModifier(ctx *ModifierContext) {}
