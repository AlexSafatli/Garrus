// Code generated from Dice.g4 by ANTLR 4.13.1. DO NOT EDIT.

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
		"", "", "", "'('", "')'", "','", "' '",
	}
	staticData.SymbolicNames = []string{
		"", "D", "SIGN", "LPAREN", "RPAREN", "COMMA", "SPACE", "WS", "Integer",
		"Id", "StringLiteral",
	}
	staticData.RuleNames = []string{
		"D", "SIGN", "LPAREN", "RPAREN", "COMMA", "SPACE", "WS", "Integer",
		"Id", "StringLiteral",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 10, 56, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 1,
		0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 5, 1, 5, 1,
		6, 1, 6, 1, 6, 1, 6, 1, 7, 4, 7, 39, 8, 7, 11, 7, 12, 7, 40, 1, 8, 1, 8,
		4, 8, 45, 8, 8, 11, 8, 12, 8, 46, 1, 9, 1, 9, 4, 9, 51, 8, 9, 11, 9, 12,
		9, 52, 1, 9, 1, 9, 0, 0, 10, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7,
		15, 8, 17, 9, 19, 10, 1, 0, 7, 2, 0, 68, 68, 100, 100, 2, 0, 43, 43, 45,
		45, 2, 0, 9, 10, 13, 13, 1, 0, 48, 57, 2, 0, 65, 90, 97, 122, 3, 0, 48,
		57, 65, 90, 97, 122, 4, 0, 10, 10, 13, 13, 39, 39, 92, 92, 58, 0, 1, 1,
		0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1,
		0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17,
		1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 1, 21, 1, 0, 0, 0, 3, 23, 1, 0, 0, 0, 5,
		25, 1, 0, 0, 0, 7, 27, 1, 0, 0, 0, 9, 29, 1, 0, 0, 0, 11, 31, 1, 0, 0,
		0, 13, 33, 1, 0, 0, 0, 15, 38, 1, 0, 0, 0, 17, 42, 1, 0, 0, 0, 19, 48,
		1, 0, 0, 0, 21, 22, 7, 0, 0, 0, 22, 2, 1, 0, 0, 0, 23, 24, 7, 1, 0, 0,
		24, 4, 1, 0, 0, 0, 25, 26, 5, 40, 0, 0, 26, 6, 1, 0, 0, 0, 27, 28, 5, 41,
		0, 0, 28, 8, 1, 0, 0, 0, 29, 30, 5, 44, 0, 0, 30, 10, 1, 0, 0, 0, 31, 32,
		5, 32, 0, 0, 32, 12, 1, 0, 0, 0, 33, 34, 7, 2, 0, 0, 34, 35, 1, 0, 0, 0,
		35, 36, 6, 6, 0, 0, 36, 14, 1, 0, 0, 0, 37, 39, 7, 3, 0, 0, 38, 37, 1,
		0, 0, 0, 39, 40, 1, 0, 0, 0, 40, 38, 1, 0, 0, 0, 40, 41, 1, 0, 0, 0, 41,
		16, 1, 0, 0, 0, 42, 44, 7, 4, 0, 0, 43, 45, 7, 5, 0, 0, 44, 43, 1, 0, 0,
		0, 45, 46, 1, 0, 0, 0, 46, 44, 1, 0, 0, 0, 46, 47, 1, 0, 0, 0, 47, 18,
		1, 0, 0, 0, 48, 50, 5, 34, 0, 0, 49, 51, 8, 6, 0, 0, 50, 49, 1, 0, 0, 0,
		51, 52, 1, 0, 0, 0, 52, 50, 1, 0, 0, 0, 52, 53, 1, 0, 0, 0, 53, 54, 1,
		0, 0, 0, 54, 55, 5, 34, 0, 0, 55, 20, 1, 0, 0, 0, 4, 0, 40, 46, 52, 1,
		6, 0, 0,
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
	l.GrammarFileName = "Dice.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// DiceLexer tokens.
const (
	DiceLexerD             = 1
	DiceLexerSIGN          = 2
	DiceLexerLPAREN        = 3
	DiceLexerRPAREN        = 4
	DiceLexerCOMMA         = 5
	DiceLexerSPACE         = 6
	DiceLexerWS            = 7
	DiceLexerInteger       = 8
	DiceLexerId            = 9
	DiceLexerStringLiteral = 10
)
