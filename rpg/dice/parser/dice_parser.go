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
		"", "", "", "", "", "'('", "')'",
	}
	staticData.SymbolicNames = []string{
		"", "DSEPARATOR", "DIGIT", "ADDOPERATOR", "MULTOPERATOR", "LPAREN",
		"RPAREN", "WS",
	}
	staticData.RuleNames = []string{
		"notation", "add", "mult", "operand", "dice", "number",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 7, 56, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7, 4,
		2, 5, 7, 5, 1, 0, 1, 0, 1, 0, 3, 0, 16, 8, 0, 1, 1, 1, 1, 1, 1, 5, 1, 21,
		8, 1, 10, 1, 12, 1, 24, 9, 1, 1, 2, 1, 2, 1, 2, 5, 2, 29, 8, 2, 10, 2,
		12, 2, 32, 9, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 3, 3, 40, 8, 3, 1,
		4, 3, 4, 43, 8, 4, 1, 4, 3, 4, 46, 8, 4, 1, 4, 1, 4, 1, 4, 1, 5, 3, 5,
		52, 8, 5, 1, 5, 1, 5, 1, 5, 0, 0, 6, 0, 2, 4, 6, 8, 10, 0, 0, 58, 0, 15,
		1, 0, 0, 0, 2, 17, 1, 0, 0, 0, 4, 25, 1, 0, 0, 0, 6, 39, 1, 0, 0, 0, 8,
		42, 1, 0, 0, 0, 10, 51, 1, 0, 0, 0, 12, 16, 3, 8, 4, 0, 13, 16, 3, 10,
		5, 0, 14, 16, 3, 2, 1, 0, 15, 12, 1, 0, 0, 0, 15, 13, 1, 0, 0, 0, 15, 14,
		1, 0, 0, 0, 16, 1, 1, 0, 0, 0, 17, 22, 3, 4, 2, 0, 18, 19, 5, 3, 0, 0,
		19, 21, 3, 4, 2, 0, 20, 18, 1, 0, 0, 0, 21, 24, 1, 0, 0, 0, 22, 20, 1,
		0, 0, 0, 22, 23, 1, 0, 0, 0, 23, 3, 1, 0, 0, 0, 24, 22, 1, 0, 0, 0, 25,
		30, 3, 6, 3, 0, 26, 27, 5, 4, 0, 0, 27, 29, 3, 6, 3, 0, 28, 26, 1, 0, 0,
		0, 29, 32, 1, 0, 0, 0, 30, 28, 1, 0, 0, 0, 30, 31, 1, 0, 0, 0, 31, 5, 1,
		0, 0, 0, 32, 30, 1, 0, 0, 0, 33, 40, 3, 8, 4, 0, 34, 40, 3, 10, 5, 0, 35,
		36, 5, 5, 0, 0, 36, 37, 3, 0, 0, 0, 37, 38, 5, 6, 0, 0, 38, 40, 1, 0, 0,
		0, 39, 33, 1, 0, 0, 0, 39, 34, 1, 0, 0, 0, 39, 35, 1, 0, 0, 0, 40, 7, 1,
		0, 0, 0, 41, 43, 5, 3, 0, 0, 42, 41, 1, 0, 0, 0, 42, 43, 1, 0, 0, 0, 43,
		45, 1, 0, 0, 0, 44, 46, 5, 2, 0, 0, 45, 44, 1, 0, 0, 0, 45, 46, 1, 0, 0,
		0, 46, 47, 1, 0, 0, 0, 47, 48, 5, 1, 0, 0, 48, 49, 5, 2, 0, 0, 49, 9, 1,
		0, 0, 0, 50, 52, 5, 3, 0, 0, 51, 50, 1, 0, 0, 0, 51, 52, 1, 0, 0, 0, 52,
		53, 1, 0, 0, 0, 53, 54, 5, 2, 0, 0, 54, 11, 1, 0, 0, 0, 7, 15, 22, 30,
		39, 42, 45, 51,
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
	DiceParserEOF          = antlr.TokenEOF
	DiceParserDSEPARATOR   = 1
	DiceParserDIGIT        = 2
	DiceParserADDOPERATOR  = 3
	DiceParserMULTOPERATOR = 4
	DiceParserLPAREN       = 5
	DiceParserRPAREN       = 6
	DiceParserWS           = 7
)

// DiceParser rules.
const (
	DiceParserRULE_notation = 0
	DiceParserRULE_add      = 1
	DiceParserRULE_mult     = 2
	DiceParserRULE_operand  = 3
	DiceParserRULE_dice     = 4
	DiceParserRULE_number   = 5
)

// INotationContext is an interface to support dynamic dispatch.
type INotationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Dice() IDiceContext
	Number() INumberContext
	Add() IAddContext

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

func (s *NotationContext) Dice() IDiceContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDiceContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDiceContext)
}

func (s *NotationContext) Number() INumberContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INumberContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INumberContext)
}

func (s *NotationContext) Add() IAddContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAddContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAddContext)
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
	p.SetState(15)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 0, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(12)
			p.Dice()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(13)
			p.Number()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(14)
			p.Add()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
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

// IAddContext is an interface to support dynamic dispatch.
type IAddContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllMult() []IMultContext
	Mult(i int) IMultContext
	AllADDOPERATOR() []antlr.TerminalNode
	ADDOPERATOR(i int) antlr.TerminalNode

	// IsAddContext differentiates from other interfaces.
	IsAddContext()
}

type AddContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAddContext() *AddContext {
	var p = new(AddContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DiceParserRULE_add
	return p
}

func InitEmptyAddContext(p *AddContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DiceParserRULE_add
}

func (*AddContext) IsAddContext() {}

func NewAddContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AddContext {
	var p = new(AddContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DiceParserRULE_add

	return p
}

func (s *AddContext) GetParser() antlr.Parser { return s.parser }

func (s *AddContext) AllMult() []IMultContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IMultContext); ok {
			len++
		}
	}

	tst := make([]IMultContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IMultContext); ok {
			tst[i] = t.(IMultContext)
			i++
		}
	}

	return tst
}

func (s *AddContext) Mult(i int) IMultContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMultContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMultContext)
}

func (s *AddContext) AllADDOPERATOR() []antlr.TerminalNode {
	return s.GetTokens(DiceParserADDOPERATOR)
}

func (s *AddContext) ADDOPERATOR(i int) antlr.TerminalNode {
	return s.GetToken(DiceParserADDOPERATOR, i)
}

func (s *AddContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AddContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AddContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DiceListener); ok {
		listenerT.EnterAdd(s)
	}
}

func (s *AddContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DiceListener); ok {
		listenerT.ExitAdd(s)
	}
}

func (p *DiceParser) Add() (localctx IAddContext) {
	localctx = NewAddContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, DiceParserRULE_add)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(17)
		p.Mult()
	}
	p.SetState(22)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == DiceParserADDOPERATOR {
		{
			p.SetState(18)
			p.Match(DiceParserADDOPERATOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(19)
			p.Mult()
		}

		p.SetState(24)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
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

// IMultContext is an interface to support dynamic dispatch.
type IMultContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllOperand() []IOperandContext
	Operand(i int) IOperandContext
	AllMULTOPERATOR() []antlr.TerminalNode
	MULTOPERATOR(i int) antlr.TerminalNode

	// IsMultContext differentiates from other interfaces.
	IsMultContext()
}

type MultContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMultContext() *MultContext {
	var p = new(MultContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DiceParserRULE_mult
	return p
}

func InitEmptyMultContext(p *MultContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DiceParserRULE_mult
}

func (*MultContext) IsMultContext() {}

func NewMultContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MultContext {
	var p = new(MultContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DiceParserRULE_mult

	return p
}

func (s *MultContext) GetParser() antlr.Parser { return s.parser }

func (s *MultContext) AllOperand() []IOperandContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IOperandContext); ok {
			len++
		}
	}

	tst := make([]IOperandContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IOperandContext); ok {
			tst[i] = t.(IOperandContext)
			i++
		}
	}

	return tst
}

func (s *MultContext) Operand(i int) IOperandContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOperandContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOperandContext)
}

func (s *MultContext) AllMULTOPERATOR() []antlr.TerminalNode {
	return s.GetTokens(DiceParserMULTOPERATOR)
}

func (s *MultContext) MULTOPERATOR(i int) antlr.TerminalNode {
	return s.GetToken(DiceParserMULTOPERATOR, i)
}

func (s *MultContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MultContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MultContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DiceListener); ok {
		listenerT.EnterMult(s)
	}
}

func (s *MultContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DiceListener); ok {
		listenerT.ExitMult(s)
	}
}

func (p *DiceParser) Mult() (localctx IMultContext) {
	localctx = NewMultContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, DiceParserRULE_mult)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(25)
		p.Operand()
	}
	p.SetState(30)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == DiceParserMULTOPERATOR {
		{
			p.SetState(26)
			p.Match(DiceParserMULTOPERATOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(27)
			p.Operand()
		}

		p.SetState(32)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
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

// IOperandContext is an interface to support dynamic dispatch.
type IOperandContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Dice() IDiceContext
	Number() INumberContext
	LPAREN() antlr.TerminalNode
	Notation() INotationContext
	RPAREN() antlr.TerminalNode

	// IsOperandContext differentiates from other interfaces.
	IsOperandContext()
}

type OperandContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOperandContext() *OperandContext {
	var p = new(OperandContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DiceParserRULE_operand
	return p
}

func InitEmptyOperandContext(p *OperandContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DiceParserRULE_operand
}

func (*OperandContext) IsOperandContext() {}

func NewOperandContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OperandContext {
	var p = new(OperandContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DiceParserRULE_operand

	return p
}

func (s *OperandContext) GetParser() antlr.Parser { return s.parser }

func (s *OperandContext) Dice() IDiceContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDiceContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDiceContext)
}

func (s *OperandContext) Number() INumberContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INumberContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INumberContext)
}

func (s *OperandContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(DiceParserLPAREN, 0)
}

func (s *OperandContext) Notation() INotationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INotationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INotationContext)
}

func (s *OperandContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(DiceParserRPAREN, 0)
}

func (s *OperandContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OperandContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OperandContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DiceListener); ok {
		listenerT.EnterOperand(s)
	}
}

func (s *OperandContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DiceListener); ok {
		listenerT.ExitOperand(s)
	}
}

func (p *DiceParser) Operand() (localctx IOperandContext) {
	localctx = NewOperandContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, DiceParserRULE_operand)
	p.SetState(39)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(33)
			p.Dice()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(34)
			p.Number()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(35)
			p.Match(DiceParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(36)
			p.Notation()
		}
		{
			p.SetState(37)
			p.Match(DiceParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
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

// IDiceContext is an interface to support dynamic dispatch.
type IDiceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DSEPARATOR() antlr.TerminalNode
	AllDIGIT() []antlr.TerminalNode
	DIGIT(i int) antlr.TerminalNode
	ADDOPERATOR() antlr.TerminalNode

	// IsDiceContext differentiates from other interfaces.
	IsDiceContext()
}

type DiceContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDiceContext() *DiceContext {
	var p = new(DiceContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DiceParserRULE_dice
	return p
}

func InitEmptyDiceContext(p *DiceContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DiceParserRULE_dice
}

func (*DiceContext) IsDiceContext() {}

func NewDiceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DiceContext {
	var p = new(DiceContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DiceParserRULE_dice

	return p
}

func (s *DiceContext) GetParser() antlr.Parser { return s.parser }

func (s *DiceContext) DSEPARATOR() antlr.TerminalNode {
	return s.GetToken(DiceParserDSEPARATOR, 0)
}

func (s *DiceContext) AllDIGIT() []antlr.TerminalNode {
	return s.GetTokens(DiceParserDIGIT)
}

func (s *DiceContext) DIGIT(i int) antlr.TerminalNode {
	return s.GetToken(DiceParserDIGIT, i)
}

func (s *DiceContext) ADDOPERATOR() antlr.TerminalNode {
	return s.GetToken(DiceParserADDOPERATOR, 0)
}

func (s *DiceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DiceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DiceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DiceListener); ok {
		listenerT.EnterDice(s)
	}
}

func (s *DiceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DiceListener); ok {
		listenerT.ExitDice(s)
	}
}

func (p *DiceParser) Dice() (localctx IDiceContext) {
	localctx = NewDiceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, DiceParserRULE_dice)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(42)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == DiceParserADDOPERATOR {
		{
			p.SetState(41)
			p.Match(DiceParserADDOPERATOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	p.SetState(45)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == DiceParserDIGIT {
		{
			p.SetState(44)
			p.Match(DiceParserDIGIT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(47)
		p.Match(DiceParserDSEPARATOR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(48)
		p.Match(DiceParserDIGIT)
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

// INumberContext is an interface to support dynamic dispatch.
type INumberContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DIGIT() antlr.TerminalNode
	ADDOPERATOR() antlr.TerminalNode

	// IsNumberContext differentiates from other interfaces.
	IsNumberContext()
}

type NumberContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNumberContext() *NumberContext {
	var p = new(NumberContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DiceParserRULE_number
	return p
}

func InitEmptyNumberContext(p *NumberContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DiceParserRULE_number
}

func (*NumberContext) IsNumberContext() {}

func NewNumberContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NumberContext {
	var p = new(NumberContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DiceParserRULE_number

	return p
}

func (s *NumberContext) GetParser() antlr.Parser { return s.parser }

func (s *NumberContext) DIGIT() antlr.TerminalNode {
	return s.GetToken(DiceParserDIGIT, 0)
}

func (s *NumberContext) ADDOPERATOR() antlr.TerminalNode {
	return s.GetToken(DiceParserADDOPERATOR, 0)
}

func (s *NumberContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumberContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NumberContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DiceListener); ok {
		listenerT.EnterNumber(s)
	}
}

func (s *NumberContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DiceListener); ok {
		listenerT.ExitNumber(s)
	}
}

func (p *DiceParser) Number() (localctx INumberContext) {
	localctx = NewNumberContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, DiceParserRULE_number)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(51)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == DiceParserADDOPERATOR {
		{
			p.SetState(50)
			p.Match(DiceParserADDOPERATOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(53)
		p.Match(DiceParserDIGIT)
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
