// Code generated from Dice.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // Dice

import "github.com/antlr4-go/antlr/v4"

// DiceListener is a complete listener for a parse tree produced by DiceParser.
type DiceListener interface {
	antlr.ParseTreeListener

	// EnterNotation is called when entering the notation production.
	EnterNotation(c *NotationContext)

	// EnterCount is called when entering the count production.
	EnterCount(c *CountContext)

	// EnterSides is called when entering the sides production.
	EnterSides(c *SidesContext)

	// EnterModifier is called when entering the modifier production.
	EnterModifier(c *ModifierContext)

	// ExitNotation is called when exiting the notation production.
	ExitNotation(c *NotationContext)

	// ExitCount is called when exiting the count production.
	ExitCount(c *CountContext)

	// ExitSides is called when exiting the sides production.
	ExitSides(c *SidesContext)

	// ExitModifier is called when exiting the modifier production.
	ExitModifier(c *ModifierContext)
}
