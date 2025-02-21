// Code generated from /Users/bytedance/go/src/code.byted.org/lark-approval/ai_sdk/clickhouse_parse/Time.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // Time

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

type TimeParser struct {
	*antlr.BaseParser
}

var TimeParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func timeParserInit() {
	staticData := &TimeParserStaticData
	staticData.LiteralNames = []string{
		"", "'('", "')'", "','", "'-'", "'>='", "'<='", "'='", "'>'", "'<'",
		"", "", "", "", "", "", "", "", "", "'toDateTimeOrZero'", "'date_add'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "EGT", "ELT", "EQ", "GT", "LT", "NUMBER", "Clause",
		"DateType", "YEAR", "MONTH", "DAY", "Date", "TIMEZONECONST", "TIMEZONE",
		"ToDateTimeOrZero", "DateAdd", "WHITESPACE",
	}
	staticData.RuleNames = []string{
		"start", "expression",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 21, 43, 2, 0, 7, 0, 2, 1, 7, 1, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 3, 1, 33,
		8, 1, 1, 1, 1, 1, 1, 1, 5, 1, 38, 8, 1, 10, 1, 12, 1, 41, 9, 1, 1, 1, 0,
		1, 2, 2, 0, 2, 0, 1, 1, 0, 5, 9, 45, 0, 4, 1, 0, 0, 0, 2, 32, 1, 0, 0,
		0, 4, 5, 3, 2, 1, 0, 5, 6, 5, 0, 0, 1, 6, 1, 1, 0, 0, 0, 7, 8, 6, 1, -1,
		0, 8, 33, 5, 11, 0, 0, 9, 33, 5, 16, 0, 0, 10, 11, 5, 19, 0, 0, 11, 12,
		5, 1, 0, 0, 12, 13, 3, 2, 1, 0, 13, 14, 5, 2, 0, 0, 14, 33, 1, 0, 0, 0,
		15, 16, 5, 19, 0, 0, 16, 17, 5, 1, 0, 0, 17, 18, 3, 2, 1, 0, 18, 19, 5,
		3, 0, 0, 19, 20, 5, 18, 0, 0, 20, 21, 5, 2, 0, 0, 21, 33, 1, 0, 0, 0, 22,
		23, 5, 20, 0, 0, 23, 24, 5, 1, 0, 0, 24, 25, 5, 12, 0, 0, 25, 26, 5, 3,
		0, 0, 26, 27, 5, 4, 0, 0, 27, 28, 5, 10, 0, 0, 28, 29, 5, 3, 0, 0, 29,
		30, 3, 2, 1, 0, 30, 31, 5, 2, 0, 0, 31, 33, 1, 0, 0, 0, 32, 7, 1, 0, 0,
		0, 32, 9, 1, 0, 0, 0, 32, 10, 1, 0, 0, 0, 32, 15, 1, 0, 0, 0, 32, 22, 1,
		0, 0, 0, 33, 39, 1, 0, 0, 0, 34, 35, 10, 6, 0, 0, 35, 36, 7, 0, 0, 0, 36,
		38, 3, 2, 1, 7, 37, 34, 1, 0, 0, 0, 38, 41, 1, 0, 0, 0, 39, 37, 1, 0, 0,
		0, 39, 40, 1, 0, 0, 0, 40, 3, 1, 0, 0, 0, 41, 39, 1, 0, 0, 0, 2, 32, 39,
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

// TimeParserInit initializes any static state used to implement TimeParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewTimeParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func TimeParserInit() {
	staticData := &TimeParserStaticData
	staticData.once.Do(timeParserInit)
}

// NewTimeParser produces a new parser instance for the optional input antlr.TokenStream.
func NewTimeParser(input antlr.TokenStream) *TimeParser {
	TimeParserInit()
	this := new(TimeParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &TimeParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "Time.g4"

	return this
}

// TimeParser tokens.
const (
	TimeParserEOF              = antlr.TokenEOF
	TimeParserT__0             = 1
	TimeParserT__1             = 2
	TimeParserT__2             = 3
	TimeParserT__3             = 4
	TimeParserEGT              = 5
	TimeParserELT              = 6
	TimeParserEQ               = 7
	TimeParserGT               = 8
	TimeParserLT               = 9
	TimeParserNUMBER           = 10
	TimeParserClause           = 11
	TimeParserDateType         = 12
	TimeParserYEAR             = 13
	TimeParserMONTH            = 14
	TimeParserDAY              = 15
	TimeParserDate             = 16
	TimeParserTIMEZONECONST    = 17
	TimeParserTIMEZONE         = 18
	TimeParserToDateTimeOrZero = 19
	TimeParserDateAdd          = 20
	TimeParserWHITESPACE       = 21
)

// TimeParser rules.
const (
	TimeParserRULE_start      = 0
	TimeParserRULE_expression = 1
)

// IStartContext is an interface to support dynamic dispatch.
type IStartContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expression() IExpressionContext
	EOF() antlr.TerminalNode

	// IsStartContext differentiates from other interfaces.
	IsStartContext()
}

type StartContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStartContext() *StartContext {
	var p = new(StartContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TimeParserRULE_start
	return p
}

func InitEmptyStartContext(p *StartContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TimeParserRULE_start
}

func (*StartContext) IsStartContext() {}

func NewStartContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StartContext {
	var p = new(StartContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TimeParserRULE_start

	return p
}

func (s *StartContext) GetParser() antlr.Parser { return s.parser }

func (s *StartContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *StartContext) EOF() antlr.TerminalNode {
	return s.GetToken(TimeParserEOF, 0)
}

func (s *StartContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StartContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StartContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TimeListener); ok {
		listenerT.EnterStart(s)
	}
}

func (s *StartContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TimeListener); ok {
		listenerT.ExitStart(s)
	}
}

func (s *StartContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TimeVisitor:
		return t.VisitStart(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TimeParser) Start_() (localctx IStartContext) {
	localctx = NewStartContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, TimeParserRULE_start)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(4)
		p.expression(0)
	}
	{
		p.SetState(5)
		p.Match(TimeParserEOF)
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

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TimeParserRULE_expression
	return p
}

func InitEmptyExpressionContext(p *ExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TimeParserRULE_expression
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TimeParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) CopyAll(ctx *ExpressionContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ToDateTimeOrZeroContext struct {
	ExpressionContext
}

func NewToDateTimeOrZeroContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ToDateTimeOrZeroContext {
	var p = new(ToDateTimeOrZeroContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *ToDateTimeOrZeroContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ToDateTimeOrZeroContext) ToDateTimeOrZero() antlr.TerminalNode {
	return s.GetToken(TimeParserToDateTimeOrZero, 0)
}

func (s *ToDateTimeOrZeroContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ToDateTimeOrZeroContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TimeListener); ok {
		listenerT.EnterToDateTimeOrZero(s)
	}
}

func (s *ToDateTimeOrZeroContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TimeListener); ok {
		listenerT.ExitToDateTimeOrZero(s)
	}
}

func (s *ToDateTimeOrZeroContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TimeVisitor:
		return t.VisitToDateTimeOrZero(s)

	default:
		return t.VisitChildren(s)
	}
}

type CompareContext struct {
	ExpressionContext
	op antlr.Token
}

func NewCompareContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CompareContext {
	var p = new(CompareContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *CompareContext) GetOp() antlr.Token { return s.op }

func (s *CompareContext) SetOp(v antlr.Token) { s.op = v }

func (s *CompareContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CompareContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *CompareContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
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

	return t.(IExpressionContext)
}

func (s *CompareContext) GT() antlr.TerminalNode {
	return s.GetToken(TimeParserGT, 0)
}

func (s *CompareContext) LT() antlr.TerminalNode {
	return s.GetToken(TimeParserLT, 0)
}

func (s *CompareContext) EQ() antlr.TerminalNode {
	return s.GetToken(TimeParserEQ, 0)
}

func (s *CompareContext) EGT() antlr.TerminalNode {
	return s.GetToken(TimeParserEGT, 0)
}

func (s *CompareContext) ELT() antlr.TerminalNode {
	return s.GetToken(TimeParserELT, 0)
}

func (s *CompareContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TimeListener); ok {
		listenerT.EnterCompare(s)
	}
}

func (s *CompareContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TimeListener); ok {
		listenerT.ExitCompare(s)
	}
}

func (s *CompareContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TimeVisitor:
		return t.VisitCompare(s)

	default:
		return t.VisitChildren(s)
	}
}

type ClauseContext struct {
	ExpressionContext
}

func NewClauseContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ClauseContext {
	var p = new(ClauseContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *ClauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ClauseContext) Clause() antlr.TerminalNode {
	return s.GetToken(TimeParserClause, 0)
}

func (s *ClauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TimeListener); ok {
		listenerT.EnterClause(s)
	}
}

func (s *ClauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TimeListener); ok {
		listenerT.ExitClause(s)
	}
}

func (s *ClauseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TimeVisitor:
		return t.VisitClause(s)

	default:
		return t.VisitChildren(s)
	}
}

type ToDateTimeOrZeroWithTimeZoneContext struct {
	ExpressionContext
}

func NewToDateTimeOrZeroWithTimeZoneContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ToDateTimeOrZeroWithTimeZoneContext {
	var p = new(ToDateTimeOrZeroWithTimeZoneContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *ToDateTimeOrZeroWithTimeZoneContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ToDateTimeOrZeroWithTimeZoneContext) ToDateTimeOrZero() antlr.TerminalNode {
	return s.GetToken(TimeParserToDateTimeOrZero, 0)
}

func (s *ToDateTimeOrZeroWithTimeZoneContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ToDateTimeOrZeroWithTimeZoneContext) TIMEZONE() antlr.TerminalNode {
	return s.GetToken(TimeParserTIMEZONE, 0)
}

func (s *ToDateTimeOrZeroWithTimeZoneContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TimeListener); ok {
		listenerT.EnterToDateTimeOrZeroWithTimeZone(s)
	}
}

func (s *ToDateTimeOrZeroWithTimeZoneContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TimeListener); ok {
		listenerT.ExitToDateTimeOrZeroWithTimeZone(s)
	}
}

func (s *ToDateTimeOrZeroWithTimeZoneContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TimeVisitor:
		return t.VisitToDateTimeOrZeroWithTimeZone(s)

	default:
		return t.VisitChildren(s)
	}
}

type DateAddContext struct {
	ExpressionContext
}

func NewDateAddContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DateAddContext {
	var p = new(DateAddContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *DateAddContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DateAddContext) DateAdd() antlr.TerminalNode {
	return s.GetToken(TimeParserDateAdd, 0)
}

func (s *DateAddContext) DateType() antlr.TerminalNode {
	return s.GetToken(TimeParserDateType, 0)
}

func (s *DateAddContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(TimeParserNUMBER, 0)
}

func (s *DateAddContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *DateAddContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TimeListener); ok {
		listenerT.EnterDateAdd(s)
	}
}

func (s *DateAddContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TimeListener); ok {
		listenerT.ExitDateAdd(s)
	}
}

func (s *DateAddContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TimeVisitor:
		return t.VisitDateAdd(s)

	default:
		return t.VisitChildren(s)
	}
}

type DateContext struct {
	ExpressionContext
}

func NewDateContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DateContext {
	var p = new(DateContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *DateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DateContext) Date() antlr.TerminalNode {
	return s.GetToken(TimeParserDate, 0)
}

func (s *DateContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TimeListener); ok {
		listenerT.EnterDate(s)
	}
}

func (s *DateContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TimeListener); ok {
		listenerT.ExitDate(s)
	}
}

func (s *DateContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TimeVisitor:
		return t.VisitDate(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TimeParser) Expression() (localctx IExpressionContext) {
	return p.expression(0)
}

func (p *TimeParser) expression(_p int) (localctx IExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 2
	p.EnterRecursionRule(localctx, 2, TimeParserRULE_expression, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(32)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 0, p.GetParserRuleContext()) {
	case 1:
		localctx = NewClauseContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(8)
			p.Match(TimeParserClause)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		localctx = NewDateContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(9)
			p.Match(TimeParserDate)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		localctx = NewToDateTimeOrZeroContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(10)
			p.Match(TimeParserToDateTimeOrZero)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(11)
			p.Match(TimeParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(12)
			p.expression(0)
		}
		{
			p.SetState(13)
			p.Match(TimeParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		localctx = NewToDateTimeOrZeroWithTimeZoneContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(15)
			p.Match(TimeParserToDateTimeOrZero)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(16)
			p.Match(TimeParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(17)
			p.expression(0)
		}
		{
			p.SetState(18)
			p.Match(TimeParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(19)
			p.Match(TimeParserTIMEZONE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(20)
			p.Match(TimeParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 5:
		localctx = NewDateAddContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(22)
			p.Match(TimeParserDateAdd)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(23)
			p.Match(TimeParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(24)
			p.Match(TimeParserDateType)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(25)
			p.Match(TimeParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(26)
			p.Match(TimeParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(27)
			p.Match(TimeParserNUMBER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(28)
			p.Match(TimeParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(29)
			p.expression(0)
		}
		{
			p.SetState(30)
			p.Match(TimeParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(39)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewCompareContext(p, NewExpressionContext(p, _parentctx, _parentState))
			p.PushNewRecursionContext(localctx, _startState, TimeParserRULE_expression)
			p.SetState(34)

			if !(p.Precpred(p.GetParserRuleContext(), 6)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
				goto errorExit
			}
			{
				p.SetState(35)

				var _lt = p.GetTokenStream().LT(1)

				localctx.(*CompareContext).op = _lt

				_la = p.GetTokenStream().LA(1)

				if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&992) != 0) {
					var _ri = p.GetErrorHandler().RecoverInline(p)

					localctx.(*CompareContext).op = _ri
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}
			{
				p.SetState(36)
				p.expression(7)
			}

		}
		p.SetState(41)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext())
		if p.HasError() {
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
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

func (p *TimeParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 1:
		var t *ExpressionContext = nil
		if localctx != nil {
			t = localctx.(*ExpressionContext)
		}
		return p.Expression_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *TimeParser) Expression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 6)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
