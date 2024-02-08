// Code generated from Dice.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // Dice

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type DiceParser struct {
	*antlr.BaseParser
}

var DiceParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func diceParserInit() {
	staticData := &DiceParserStaticData
	staticData.LiteralNames = []string{
		"", "", "", "'('", "')'", "','", "' '",
	}
	staticData.SymbolicNames = []string{
		"", "D", "SIGN", "LPAREN", "RPAREN", "COMMA", "SPACE", "WS", "Integer",
		"Id", "StringLiteral",
	}
	staticData.RuleNames = []string{
		"notation", "count", "sides", "modifier",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 10, 23, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 1, 0, 3,
		0, 10, 8, 0, 1, 0, 1, 0, 3, 0, 14, 8, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 3,
		1, 3, 1, 3, 1, 3, 0, 0, 4, 0, 2, 4, 6, 0, 0, 20, 0, 9, 1, 0, 0, 0, 2, 15,
		1, 0, 0, 0, 4, 17, 1, 0, 0, 0, 6, 19, 1, 0, 0, 0, 8, 10, 3, 2, 1, 0, 9,
		8, 1, 0, 0, 0, 9, 10, 1, 0, 0, 0, 10, 11, 1, 0, 0, 0, 11, 13, 3, 4, 2,
		0, 12, 14, 3, 6, 3, 0, 13, 12, 1, 0, 0, 0, 13, 14, 1, 0, 0, 0, 14, 1, 1,
		0, 0, 0, 15, 16, 5, 8, 0, 0, 16, 3, 1, 0, 0, 0, 17, 18, 5, 9, 0, 0, 18,
		5, 1, 0, 0, 0, 19, 20, 5, 2, 0, 0, 20, 21, 5, 8, 0, 0, 21, 7, 1, 0, 0,
		0, 2, 9, 13,
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

// DiceParserInit initializes any static state used to implement DiceParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewDiceParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func DiceParserInit() {
	staticData := &DiceParserStaticData
	staticData.once.Do(diceParserInit)
}

// NewDiceParser produces a new parser instance for the optional input antlr.TokenStream.
func NewDiceParser(input antlr.TokenStream) *DiceParser {
	DiceParserInit()
	this := new(DiceParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &DiceParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "Dice.g4"

	return this
}

// DiceParser tokens.
const (
	DiceParserEOF           = antlr.TokenEOF
	DiceParserD             = 1
	DiceParserSIGN          = 2
	DiceParserLPAREN        = 3
	DiceParserRPAREN        = 4
	DiceParserCOMMA         = 5
	DiceParserSPACE         = 6
	DiceParserWS            = 7
	DiceParserInteger       = 8
	DiceParserId            = 9
	DiceParserStringLiteral = 10
)

// DiceParser rules.
const (
	DiceParserRULE_notation = 0
	DiceParserRULE_count    = 1
	DiceParserRULE_sides    = 2
	DiceParserRULE_modifier = 3
)

// INotationContext is an interface to support dynamic dispatch.
type INotationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Sides() ISidesContext
	Count() ICountContext
	Modifier() IModifierContext

	// IsNotationContext differentiates from other interfaces.
	IsNotationContext()
}

type NotationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNotationContext() *NotationContext {
	var p = new(NotationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DiceParserRULE_notation
	return p
}

func InitEmptyNotationContext(p *NotationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DiceParserRULE_notation
}

func (*NotationContext) IsNotationContext() {}

func NewNotationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NotationContext {
	var p = new(NotationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DiceParserRULE_notation

	return p
}

func (s *NotationContext) GetParser() antlr.Parser { return s.parser }

func (s *NotationContext) Sides() ISidesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISidesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISidesContext)
}

func (s *NotationContext) Count() ICountContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICountContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICountContext)
}

func (s *NotationContext) Modifier() IModifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IModifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IModifierContext)
}

func (s *NotationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NotationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NotationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DiceListener); ok {
		listenerT.EnterNotation(s)
	}
}

func (s *NotationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DiceListener); ok {
		listenerT.ExitNotation(s)
	}
}

func (p *DiceParser) Notation() (localctx INotationContext) {
	localctx = NewNotationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, DiceParserRULE_notation)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(9)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == DiceParserInteger {
		{
			p.SetState(8)
			p.Count()
		}

	}
	{
		p.SetState(11)
		p.Sides()
	}
	p.SetState(13)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == DiceParserSIGN {
		{
			p.SetState(12)
			p.Modifier()
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICountContext is an interface to support dynamic dispatch.
type ICountContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Integer() antlr.TerminalNode

	// IsCountContext differentiates from other interfaces.
	IsCountContext()
}

type CountContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCountContext() *CountContext {
	var p = new(CountContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DiceParserRULE_count
	return p
}

func InitEmptyCountContext(p *CountContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DiceParserRULE_count
}

func (*CountContext) IsCountContext() {}

func NewCountContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CountContext {
	var p = new(CountContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DiceParserRULE_count

	return p
}

func (s *CountContext) GetParser() antlr.Parser { return s.parser }

func (s *CountContext) Integer() antlr.TerminalNode {
	return s.GetToken(DiceParserInteger, 0)
}

func (s *CountContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CountContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CountContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DiceListener); ok {
		listenerT.EnterCount(s)
	}
}

func (s *CountContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DiceListener); ok {
		listenerT.ExitCount(s)
	}
}

func (p *DiceParser) Count() (localctx ICountContext) {
	localctx = NewCountContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, DiceParserRULE_count)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(15)
		p.Match(DiceParserInteger)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISidesContext is an interface to support dynamic dispatch.
type ISidesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Id() antlr.TerminalNode

	// IsSidesContext differentiates from other interfaces.
	IsSidesContext()
}

type SidesContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySidesContext() *SidesContext {
	var p = new(SidesContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DiceParserRULE_sides
	return p
}

func InitEmptySidesContext(p *SidesContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DiceParserRULE_sides
}

func (*SidesContext) IsSidesContext() {}

func NewSidesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SidesContext {
	var p = new(SidesContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DiceParserRULE_sides

	return p
}

func (s *SidesContext) GetParser() antlr.Parser { return s.parser }

func (s *SidesContext) Id() antlr.TerminalNode {
	return s.GetToken(DiceParserId, 0)
}

func (s *SidesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SidesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SidesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DiceListener); ok {
		listenerT.EnterSides(s)
	}
}

func (s *SidesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DiceListener); ok {
		listenerT.ExitSides(s)
	}
}

func (p *DiceParser) Sides() (localctx ISidesContext) {
	localctx = NewSidesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, DiceParserRULE_sides)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(17)
		p.Match(DiceParserId)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IModifierContext is an interface to support dynamic dispatch.
type IModifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SIGN() antlr.TerminalNode
	Integer() antlr.TerminalNode

	// IsModifierContext differentiates from other interfaces.
	IsModifierContext()
}

type ModifierContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyModifierContext() *ModifierContext {
	var p = new(ModifierContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DiceParserRULE_modifier
	return p
}

func InitEmptyModifierContext(p *ModifierContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DiceParserRULE_modifier
}

func (*ModifierContext) IsModifierContext() {}

func NewModifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ModifierContext {
	var p = new(ModifierContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DiceParserRULE_modifier

	return p
}

func (s *ModifierContext) GetParser() antlr.Parser { return s.parser }

func (s *ModifierContext) SIGN() antlr.TerminalNode {
	return s.GetToken(DiceParserSIGN, 0)
}

func (s *ModifierContext) Integer() antlr.TerminalNode {
	return s.GetToken(DiceParserInteger, 0)
}

func (s *ModifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ModifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ModifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DiceListener); ok {
		listenerT.EnterModifier(s)
	}
}

func (s *ModifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DiceListener); ok {
		listenerT.ExitModifier(s)
	}
}

func (p *DiceParser) Modifier() (localctx IModifierContext) {
	localctx = NewModifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, DiceParserRULE_modifier)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(19)
		p.Match(DiceParserSIGN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(20)
		p.Match(DiceParserInteger)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}
