// Code generated from Dice.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // Dice

import "github.com/antlr4-go/antlr/v4"

// DiceListener is a complete listener for a parse tree produced by DiceParser.
type DiceListener interface {
	antlr.ParseTreeListener

	// EnterNotation is called when entering the notation production.
	EnterNotation(c *NotationContext)

	// EnterAdd is called when entering the add production.
	EnterAdd(c *AddContext)

	// EnterMult is called when entering the mult production.
	EnterMult(c *MultContext)

	// EnterOperand is called when entering the operand production.
	EnterOperand(c *OperandContext)

	// EnterDice is called when entering the dice production.
	EnterDice(c *DiceContext)

	// EnterNumber is called when entering the number production.
	EnterNumber(c *NumberContext)

	// ExitNotation is called when exiting the notation production.
	ExitNotation(c *NotationContext)

	// ExitAdd is called when exiting the add production.
	ExitAdd(c *AddContext)

	// ExitMult is called when exiting the mult production.
	ExitMult(c *MultContext)

	// ExitOperand is called when exiting the operand production.
	ExitOperand(c *OperandContext)

	// ExitDice is called when exiting the dice production.
	ExitDice(c *DiceContext)

	// ExitNumber is called when exiting the number production.
	ExitNumber(c *NumberContext)
}
