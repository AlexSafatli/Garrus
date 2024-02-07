// Code generated from DiceLexer.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"sync"
	"unicode"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type DiceLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var DiceLexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	ChannelNames           []string
	ModeNames              []string
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func dicelexerLexerInit() {
	staticData := &DiceLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "", "", "", "", "'('", "')'",
	}
	staticData.SymbolicNames = []string{
		"", "DSEPARATOR", "DIGIT", "ADDOPERATOR", "MULTOPERATOR", "LPAREN",
		"RPAREN", "WS",
	}
	staticData.RuleNames = []string{
		"DSEPARATOR", "DIGIT", "ADDOPERATOR", "MULTOPERATOR", "LPAREN", "RPAREN",
		"ADD", "SUB", "MULT", "DIV", "WS",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 7, 57, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 1, 0, 1, 0, 1, 1, 4, 1, 27, 8, 1, 11, 1, 12, 1, 28, 1, 2, 1,
		2, 3, 2, 33, 8, 2, 1, 3, 1, 3, 3, 3, 37, 8, 3, 1, 4, 1, 4, 1, 5, 1, 5,
		1, 6, 1, 6, 1, 7, 1, 7, 1, 8, 1, 8, 1, 9, 1, 9, 1, 10, 4, 10, 52, 8, 10,
		11, 10, 12, 10, 53, 1, 10, 1, 10, 0, 0, 11, 1, 1, 3, 2, 5, 3, 7, 4, 9,
		5, 11, 6, 13, 0, 15, 0, 17, 0, 19, 0, 21, 7, 1, 0, 2, 2, 0, 68, 68, 100,
		100, 2, 0, 9, 10, 13, 13, 56, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5,
		1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 21,
		1, 0, 0, 0, 1, 23, 1, 0, 0, 0, 3, 26, 1, 0, 0, 0, 5, 32, 1, 0, 0, 0, 7,
		36, 1, 0, 0, 0, 9, 38, 1, 0, 0, 0, 11, 40, 1, 0, 0, 0, 13, 42, 1, 0, 0,
		0, 15, 44, 1, 0, 0, 0, 17, 46, 1, 0, 0, 0, 19, 48, 1, 0, 0, 0, 21, 51,
		1, 0, 0, 0, 23, 24, 7, 0, 0, 0, 24, 2, 1, 0, 0, 0, 25, 27, 2, 48, 57, 0,
		26, 25, 1, 0, 0, 0, 27, 28, 1, 0, 0, 0, 28, 26, 1, 0, 0, 0, 28, 29, 1,
		0, 0, 0, 29, 4, 1, 0, 0, 0, 30, 33, 3, 13, 6, 0, 31, 33, 3, 15, 7, 0, 32,
		30, 1, 0, 0, 0, 32, 31, 1, 0, 0, 0, 33, 6, 1, 0, 0, 0, 34, 37, 3, 17, 8,
		0, 35, 37, 3, 19, 9, 0, 36, 34, 1, 0, 0, 0, 36, 35, 1, 0, 0, 0, 37, 8,
		1, 0, 0, 0, 38, 39, 5, 40, 0, 0, 39, 10, 1, 0, 0, 0, 40, 41, 5, 41, 0,
		0, 41, 12, 1, 0, 0, 0, 42, 43, 5, 43, 0, 0, 43, 14, 1, 0, 0, 0, 44, 45,
		5, 45, 0, 0, 45, 16, 1, 0, 0, 0, 46, 47, 5, 42, 0, 0, 47, 18, 1, 0, 0,
		0, 48, 49, 5, 47, 0, 0, 49, 20, 1, 0, 0, 0, 50, 52, 7, 1, 0, 0, 51, 50,
		1, 0, 0, 0, 52, 53, 1, 0, 0, 0, 53, 51, 1, 0, 0, 0, 53, 54, 1, 0, 0, 0,
		54, 55, 1, 0, 0, 0, 55, 56, 6, 10, 0, 0, 56, 22, 1, 0, 0, 0, 5, 0, 28,
		32, 36, 53, 1, 6, 0, 0,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// DiceLexerInit initializes any static state used to implement DiceLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewDiceLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func DiceLexerInit() {
	staticData := &DiceLexerLexerStaticData
	staticData.once.Do(dicelexerLexerInit)
}

// NewDiceLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewDiceLexer(input antlr.CharStream) *DiceLexer {
	DiceLexerInit()
	l := new(DiceLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &DiceLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "DiceLexer.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// DiceLexer tokens.
const (
	DiceLexerDSEPARATOR   = 1
	DiceLexerDIGIT        = 2
	DiceLexerADDOPERATOR  = 3
	DiceLexerMULTOPERATOR = 4
	DiceLexerLPAREN       = 5
	DiceLexerRPAREN       = 6
	DiceLexerWS           = 7
)
