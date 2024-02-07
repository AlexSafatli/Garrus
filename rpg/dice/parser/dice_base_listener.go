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

// EnterAdd is called when production add is entered.
func (s *BaseDiceListener) EnterAdd(ctx *AddContext) {}

// ExitAdd is called when production add is exited.
func (s *BaseDiceListener) ExitAdd(ctx *AddContext) {}

// EnterMult is called when production mult is entered.
func (s *BaseDiceListener) EnterMult(ctx *MultContext) {}

// ExitMult is called when production mult is exited.
func (s *BaseDiceListener) ExitMult(ctx *MultContext) {}

// EnterOperand is called when production operand is entered.
func (s *BaseDiceListener) EnterOperand(ctx *OperandContext) {}

// ExitOperand is called when production operand is exited.
func (s *BaseDiceListener) ExitOperand(ctx *OperandContext) {}

// EnterDice is called when production dice is entered.
func (s *BaseDiceListener) EnterDice(ctx *DiceContext) {}

// ExitDice is called when production dice is exited.
func (s *BaseDiceListener) ExitDice(ctx *DiceContext) {}

// EnterNumber is called when production number is entered.
func (s *BaseDiceListener) EnterNumber(ctx *NumberContext) {}

// ExitNumber is called when production number is exited.
func (s *BaseDiceListener) ExitNumber(ctx *NumberContext) {}
