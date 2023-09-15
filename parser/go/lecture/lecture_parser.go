// Code generated from Lecture.g4 by ANTLR 4.13.1. DO NOT EDIT.

package lecture // Lecture
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

type LectureParser struct {
	*antlr.BaseParser
}

var LectureParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func lectureParserInit() {
	staticData := &LectureParserStaticData
	staticData.LiteralNames = []string{
		"", "'okay, hear me out'", "'i rest my case'", "'let's say'", "'then we have'",
		"", "'is'", "'literally'", "'plus'", "' '",
	}
	staticData.SymbolicNames = []string{
		"", "START_CLAUSE", "END_CLAUSE", "DECLARATION_STATEMENT_INTRO", "PRINT_STATEMENT_INTRO",
		"TERMINATOR", "IS", "LITERALLY", "PLUS", "SPACE", "ALPHANUMERICSTRING",
		"INTEGER", "WS",
	}
	staticData.RuleNames = []string{
		"lecture", "program", "startClause", "endClause", "statement", "declarationStatement",
		"printStatement", "identifier", "valueClause", "value", "literalClause",
		"literal", "operator", "number",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 12, 88, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 1, 0, 1, 0, 1, 0, 1, 1, 1,
		1, 4, 1, 34, 8, 1, 11, 1, 12, 1, 35, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 3,
		1, 3, 1, 3, 1, 4, 1, 4, 3, 4, 48, 8, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1,
		5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 8, 1,
		8, 1, 8, 1, 8, 1, 8, 1, 8, 3, 8, 72, 8, 8, 1, 9, 1, 9, 3, 9, 76, 8, 9,
		1, 10, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 12, 1, 12, 1, 13, 1, 13, 1,
		13, 0, 0, 14, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 0, 0,
		77, 0, 28, 1, 0, 0, 0, 2, 31, 1, 0, 0, 0, 4, 39, 1, 0, 0, 0, 6, 42, 1,
		0, 0, 0, 8, 47, 1, 0, 0, 0, 10, 51, 1, 0, 0, 0, 12, 59, 1, 0, 0, 0, 14,
		63, 1, 0, 0, 0, 16, 65, 1, 0, 0, 0, 18, 75, 1, 0, 0, 0, 20, 77, 1, 0, 0,
		0, 22, 81, 1, 0, 0, 0, 24, 83, 1, 0, 0, 0, 26, 85, 1, 0, 0, 0, 28, 29,
		3, 2, 1, 0, 29, 30, 5, 0, 0, 1, 30, 1, 1, 0, 0, 0, 31, 33, 3, 4, 2, 0,
		32, 34, 3, 8, 4, 0, 33, 32, 1, 0, 0, 0, 34, 35, 1, 0, 0, 0, 35, 33, 1,
		0, 0, 0, 35, 36, 1, 0, 0, 0, 36, 37, 1, 0, 0, 0, 37, 38, 3, 6, 3, 0, 38,
		3, 1, 0, 0, 0, 39, 40, 5, 1, 0, 0, 40, 41, 5, 5, 0, 0, 41, 5, 1, 0, 0,
		0, 42, 43, 5, 2, 0, 0, 43, 44, 5, 5, 0, 0, 44, 7, 1, 0, 0, 0, 45, 48, 3,
		10, 5, 0, 46, 48, 3, 12, 6, 0, 47, 45, 1, 0, 0, 0, 47, 46, 1, 0, 0, 0,
		48, 49, 1, 0, 0, 0, 49, 50, 5, 5, 0, 0, 50, 9, 1, 0, 0, 0, 51, 52, 5, 3,
		0, 0, 52, 53, 5, 9, 0, 0, 53, 54, 3, 14, 7, 0, 54, 55, 5, 9, 0, 0, 55,
		56, 5, 6, 0, 0, 56, 57, 5, 9, 0, 0, 57, 58, 3, 16, 8, 0, 58, 11, 1, 0,
		0, 0, 59, 60, 5, 4, 0, 0, 60, 61, 5, 9, 0, 0, 61, 62, 3, 16, 8, 0, 62,
		13, 1, 0, 0, 0, 63, 64, 5, 10, 0, 0, 64, 15, 1, 0, 0, 0, 65, 71, 3, 18,
		9, 0, 66, 67, 5, 9, 0, 0, 67, 68, 3, 24, 12, 0, 68, 69, 5, 9, 0, 0, 69,
		70, 3, 16, 8, 0, 70, 72, 1, 0, 0, 0, 71, 66, 1, 0, 0, 0, 71, 72, 1, 0,
		0, 0, 72, 17, 1, 0, 0, 0, 73, 76, 3, 20, 10, 0, 74, 76, 3, 14, 7, 0, 75,
		73, 1, 0, 0, 0, 75, 74, 1, 0, 0, 0, 76, 19, 1, 0, 0, 0, 77, 78, 5, 7, 0,
		0, 78, 79, 5, 9, 0, 0, 79, 80, 3, 22, 11, 0, 80, 21, 1, 0, 0, 0, 81, 82,
		3, 26, 13, 0, 82, 23, 1, 0, 0, 0, 83, 84, 5, 8, 0, 0, 84, 25, 1, 0, 0,
		0, 85, 86, 5, 11, 0, 0, 86, 27, 1, 0, 0, 0, 4, 35, 47, 71, 75,
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

// LectureParserInit initializes any static state used to implement LectureParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewLectureParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func LectureParserInit() {
	staticData := &LectureParserStaticData
	staticData.once.Do(lectureParserInit)
}

// NewLectureParser produces a new parser instance for the optional input antlr.TokenStream.
func NewLectureParser(input antlr.TokenStream) *LectureParser {
	LectureParserInit()
	this := new(LectureParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &LectureParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "Lecture.g4"

	return this
}

// LectureParser tokens.
const (
	LectureParserEOF                         = antlr.TokenEOF
	LectureParserSTART_CLAUSE                = 1
	LectureParserEND_CLAUSE                  = 2
	LectureParserDECLARATION_STATEMENT_INTRO = 3
	LectureParserPRINT_STATEMENT_INTRO       = 4
	LectureParserTERMINATOR                  = 5
	LectureParserIS                          = 6
	LectureParserLITERALLY                   = 7
	LectureParserPLUS                        = 8
	LectureParserSPACE                       = 9
	LectureParserALPHANUMERICSTRING          = 10
	LectureParserINTEGER                     = 11
	LectureParserWS                          = 12
)

// LectureParser rules.
const (
	LectureParserRULE_lecture              = 0
	LectureParserRULE_program              = 1
	LectureParserRULE_startClause          = 2
	LectureParserRULE_endClause            = 3
	LectureParserRULE_statement            = 4
	LectureParserRULE_declarationStatement = 5
	LectureParserRULE_printStatement       = 6
	LectureParserRULE_identifier           = 7
	LectureParserRULE_valueClause          = 8
	LectureParserRULE_value                = 9
	LectureParserRULE_literalClause        = 10
	LectureParserRULE_literal              = 11
	LectureParserRULE_operator             = 12
	LectureParserRULE_number               = 13
)

// ILectureContext is an interface to support dynamic dispatch.
type ILectureContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Program() IProgramContext
	EOF() antlr.TerminalNode

	// IsLectureContext differentiates from other interfaces.
	IsLectureContext()
}

type LectureContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLectureContext() *LectureContext {
	var p = new(LectureContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LectureParserRULE_lecture
	return p
}

func InitEmptyLectureContext(p *LectureContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LectureParserRULE_lecture
}

func (*LectureContext) IsLectureContext() {}

func NewLectureContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LectureContext {
	var p = new(LectureContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LectureParserRULE_lecture

	return p
}

func (s *LectureContext) GetParser() antlr.Parser { return s.parser }

func (s *LectureContext) Program() IProgramContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IProgramContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IProgramContext)
}

func (s *LectureContext) EOF() antlr.TerminalNode {
	return s.GetToken(LectureParserEOF, 0)
}

func (s *LectureContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LectureContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LectureContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LectureListener); ok {
		listenerT.EnterLecture(s)
	}
}

func (s *LectureContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LectureListener); ok {
		listenerT.ExitLecture(s)
	}
}

func (p *LectureParser) Lecture() (localctx ILectureContext) {
	localctx = NewLectureContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, LectureParserRULE_lecture)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(28)
		p.Program()
	}
	{
		p.SetState(29)
		p.Match(LectureParserEOF)
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

// IProgramContext is an interface to support dynamic dispatch.
type IProgramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	StartClause() IStartClauseContext
	EndClause() IEndClauseContext
	AllStatement() []IStatementContext
	Statement(i int) IStatementContext

	// IsProgramContext differentiates from other interfaces.
	IsProgramContext()
}

type ProgramContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgramContext() *ProgramContext {
	var p = new(ProgramContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LectureParserRULE_program
	return p
}

func InitEmptyProgramContext(p *ProgramContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LectureParserRULE_program
}

func (*ProgramContext) IsProgramContext() {}

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext {
	var p = new(ProgramContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LectureParserRULE_program

	return p
}

func (s *ProgramContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramContext) StartClause() IStartClauseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStartClauseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStartClauseContext)
}

func (s *ProgramContext) EndClause() IEndClauseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEndClauseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEndClauseContext)
}

func (s *ProgramContext) AllStatement() []IStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatementContext); ok {
			len++
		}
	}

	tst := make([]IStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatementContext); ok {
			tst[i] = t.(IStatementContext)
			i++
		}
	}

	return tst
}

func (s *ProgramContext) Statement(i int) IStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
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

	return t.(IStatementContext)
}

func (s *ProgramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LectureListener); ok {
		listenerT.EnterProgram(s)
	}
}

func (s *ProgramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LectureListener); ok {
		listenerT.ExitProgram(s)
	}
}

func (p *LectureParser) Program() (localctx IProgramContext) {
	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, LectureParserRULE_program)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(31)
		p.StartClause()
	}
	p.SetState(33)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == LectureParserDECLARATION_STATEMENT_INTRO || _la == LectureParserPRINT_STATEMENT_INTRO {
		{
			p.SetState(32)
			p.Statement()
		}

		p.SetState(35)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(37)
		p.EndClause()
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

// IStartClauseContext is an interface to support dynamic dispatch.
type IStartClauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	START_CLAUSE() antlr.TerminalNode
	TERMINATOR() antlr.TerminalNode

	// IsStartClauseContext differentiates from other interfaces.
	IsStartClauseContext()
}

type StartClauseContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStartClauseContext() *StartClauseContext {
	var p = new(StartClauseContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LectureParserRULE_startClause
	return p
}

func InitEmptyStartClauseContext(p *StartClauseContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LectureParserRULE_startClause
}

func (*StartClauseContext) IsStartClauseContext() {}

func NewStartClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StartClauseContext {
	var p = new(StartClauseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LectureParserRULE_startClause

	return p
}

func (s *StartClauseContext) GetParser() antlr.Parser { return s.parser }

func (s *StartClauseContext) START_CLAUSE() antlr.TerminalNode {
	return s.GetToken(LectureParserSTART_CLAUSE, 0)
}

func (s *StartClauseContext) TERMINATOR() antlr.TerminalNode {
	return s.GetToken(LectureParserTERMINATOR, 0)
}

func (s *StartClauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StartClauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StartClauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LectureListener); ok {
		listenerT.EnterStartClause(s)
	}
}

func (s *StartClauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LectureListener); ok {
		listenerT.ExitStartClause(s)
	}
}

func (p *LectureParser) StartClause() (localctx IStartClauseContext) {
	localctx = NewStartClauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, LectureParserRULE_startClause)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(39)
		p.Match(LectureParserSTART_CLAUSE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(40)
		p.Match(LectureParserTERMINATOR)
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

// IEndClauseContext is an interface to support dynamic dispatch.
type IEndClauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	END_CLAUSE() antlr.TerminalNode
	TERMINATOR() antlr.TerminalNode

	// IsEndClauseContext differentiates from other interfaces.
	IsEndClauseContext()
}

type EndClauseContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEndClauseContext() *EndClauseContext {
	var p = new(EndClauseContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LectureParserRULE_endClause
	return p
}

func InitEmptyEndClauseContext(p *EndClauseContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LectureParserRULE_endClause
}

func (*EndClauseContext) IsEndClauseContext() {}

func NewEndClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EndClauseContext {
	var p = new(EndClauseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LectureParserRULE_endClause

	return p
}

func (s *EndClauseContext) GetParser() antlr.Parser { return s.parser }

func (s *EndClauseContext) END_CLAUSE() antlr.TerminalNode {
	return s.GetToken(LectureParserEND_CLAUSE, 0)
}

func (s *EndClauseContext) TERMINATOR() antlr.TerminalNode {
	return s.GetToken(LectureParserTERMINATOR, 0)
}

func (s *EndClauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EndClauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EndClauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LectureListener); ok {
		listenerT.EnterEndClause(s)
	}
}

func (s *EndClauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LectureListener); ok {
		listenerT.ExitEndClause(s)
	}
}

func (p *LectureParser) EndClause() (localctx IEndClauseContext) {
	localctx = NewEndClauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, LectureParserRULE_endClause)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(42)
		p.Match(LectureParserEND_CLAUSE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(43)
		p.Match(LectureParserTERMINATOR)
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

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TERMINATOR() antlr.TerminalNode
	DeclarationStatement() IDeclarationStatementContext
	PrintStatement() IPrintStatementContext

	// IsStatementContext differentiates from other interfaces.
	IsStatementContext()
}

type StatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementContext() *StatementContext {
	var p = new(StatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LectureParserRULE_statement
	return p
}

func InitEmptyStatementContext(p *StatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LectureParserRULE_statement
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LectureParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) TERMINATOR() antlr.TerminalNode {
	return s.GetToken(LectureParserTERMINATOR, 0)
}

func (s *StatementContext) DeclarationStatement() IDeclarationStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclarationStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclarationStatementContext)
}

func (s *StatementContext) PrintStatement() IPrintStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrintStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrintStatementContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LectureListener); ok {
		listenerT.EnterStatement(s)
	}
}

func (s *StatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LectureListener); ok {
		listenerT.ExitStatement(s)
	}
}

func (p *LectureParser) Statement() (localctx IStatementContext) {
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, LectureParserRULE_statement)
	p.EnterOuterAlt(localctx, 1)
	p.SetState(47)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case LectureParserDECLARATION_STATEMENT_INTRO:
		{
			p.SetState(45)
			p.DeclarationStatement()
		}

	case LectureParserPRINT_STATEMENT_INTRO:
		{
			p.SetState(46)
			p.PrintStatement()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	{
		p.SetState(49)
		p.Match(LectureParserTERMINATOR)
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

// IDeclarationStatementContext is an interface to support dynamic dispatch.
type IDeclarationStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DECLARATION_STATEMENT_INTRO() antlr.TerminalNode
	AllSPACE() []antlr.TerminalNode
	SPACE(i int) antlr.TerminalNode
	Identifier() IIdentifierContext
	IS() antlr.TerminalNode
	ValueClause() IValueClauseContext

	// IsDeclarationStatementContext differentiates from other interfaces.
	IsDeclarationStatementContext()
}

type DeclarationStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeclarationStatementContext() *DeclarationStatementContext {
	var p = new(DeclarationStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LectureParserRULE_declarationStatement
	return p
}

func InitEmptyDeclarationStatementContext(p *DeclarationStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LectureParserRULE_declarationStatement
}

func (*DeclarationStatementContext) IsDeclarationStatementContext() {}

func NewDeclarationStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeclarationStatementContext {
	var p = new(DeclarationStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LectureParserRULE_declarationStatement

	return p
}

func (s *DeclarationStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *DeclarationStatementContext) DECLARATION_STATEMENT_INTRO() antlr.TerminalNode {
	return s.GetToken(LectureParserDECLARATION_STATEMENT_INTRO, 0)
}

func (s *DeclarationStatementContext) AllSPACE() []antlr.TerminalNode {
	return s.GetTokens(LectureParserSPACE)
}

func (s *DeclarationStatementContext) SPACE(i int) antlr.TerminalNode {
	return s.GetToken(LectureParserSPACE, i)
}

func (s *DeclarationStatementContext) Identifier() IIdentifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *DeclarationStatementContext) IS() antlr.TerminalNode {
	return s.GetToken(LectureParserIS, 0)
}

func (s *DeclarationStatementContext) ValueClause() IValueClauseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueClauseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueClauseContext)
}

func (s *DeclarationStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclarationStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DeclarationStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LectureListener); ok {
		listenerT.EnterDeclarationStatement(s)
	}
}

func (s *DeclarationStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LectureListener); ok {
		listenerT.ExitDeclarationStatement(s)
	}
}

func (p *LectureParser) DeclarationStatement() (localctx IDeclarationStatementContext) {
	localctx = NewDeclarationStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, LectureParserRULE_declarationStatement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(51)
		p.Match(LectureParserDECLARATION_STATEMENT_INTRO)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(52)
		p.Match(LectureParserSPACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(53)
		p.Identifier()
	}
	{
		p.SetState(54)
		p.Match(LectureParserSPACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(55)
		p.Match(LectureParserIS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(56)
		p.Match(LectureParserSPACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(57)
		p.ValueClause()
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

// IPrintStatementContext is an interface to support dynamic dispatch.
type IPrintStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PRINT_STATEMENT_INTRO() antlr.TerminalNode
	SPACE() antlr.TerminalNode
	ValueClause() IValueClauseContext

	// IsPrintStatementContext differentiates from other interfaces.
	IsPrintStatementContext()
}

type PrintStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrintStatementContext() *PrintStatementContext {
	var p = new(PrintStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LectureParserRULE_printStatement
	return p
}

func InitEmptyPrintStatementContext(p *PrintStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LectureParserRULE_printStatement
}

func (*PrintStatementContext) IsPrintStatementContext() {}

func NewPrintStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrintStatementContext {
	var p = new(PrintStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LectureParserRULE_printStatement

	return p
}

func (s *PrintStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *PrintStatementContext) PRINT_STATEMENT_INTRO() antlr.TerminalNode {
	return s.GetToken(LectureParserPRINT_STATEMENT_INTRO, 0)
}

func (s *PrintStatementContext) SPACE() antlr.TerminalNode {
	return s.GetToken(LectureParserSPACE, 0)
}

func (s *PrintStatementContext) ValueClause() IValueClauseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueClauseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueClauseContext)
}

func (s *PrintStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrintStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrintStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LectureListener); ok {
		listenerT.EnterPrintStatement(s)
	}
}

func (s *PrintStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LectureListener); ok {
		listenerT.ExitPrintStatement(s)
	}
}

func (p *LectureParser) PrintStatement() (localctx IPrintStatementContext) {
	localctx = NewPrintStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, LectureParserRULE_printStatement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(59)
		p.Match(LectureParserPRINT_STATEMENT_INTRO)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(60)
		p.Match(LectureParserSPACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(61)
		p.ValueClause()
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

// IIdentifierContext is an interface to support dynamic dispatch.
type IIdentifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ALPHANUMERICSTRING() antlr.TerminalNode

	// IsIdentifierContext differentiates from other interfaces.
	IsIdentifierContext()
}

type IdentifierContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIdentifierContext() *IdentifierContext {
	var p = new(IdentifierContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LectureParserRULE_identifier
	return p
}

func InitEmptyIdentifierContext(p *IdentifierContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LectureParserRULE_identifier
}

func (*IdentifierContext) IsIdentifierContext() {}

func NewIdentifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdentifierContext {
	var p = new(IdentifierContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LectureParserRULE_identifier

	return p
}

func (s *IdentifierContext) GetParser() antlr.Parser { return s.parser }

func (s *IdentifierContext) ALPHANUMERICSTRING() antlr.TerminalNode {
	return s.GetToken(LectureParserALPHANUMERICSTRING, 0)
}

func (s *IdentifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IdentifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LectureListener); ok {
		listenerT.EnterIdentifier(s)
	}
}

func (s *IdentifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LectureListener); ok {
		listenerT.ExitIdentifier(s)
	}
}

func (p *LectureParser) Identifier() (localctx IIdentifierContext) {
	localctx = NewIdentifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, LectureParserRULE_identifier)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(63)
		p.Match(LectureParserALPHANUMERICSTRING)
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

// IValueClauseContext is an interface to support dynamic dispatch.
type IValueClauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Value() IValueContext
	AllSPACE() []antlr.TerminalNode
	SPACE(i int) antlr.TerminalNode
	Operator() IOperatorContext
	ValueClause() IValueClauseContext

	// IsValueClauseContext differentiates from other interfaces.
	IsValueClauseContext()
}

type ValueClauseContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValueClauseContext() *ValueClauseContext {
	var p = new(ValueClauseContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LectureParserRULE_valueClause
	return p
}

func InitEmptyValueClauseContext(p *ValueClauseContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LectureParserRULE_valueClause
}

func (*ValueClauseContext) IsValueClauseContext() {}

func NewValueClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueClauseContext {
	var p = new(ValueClauseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LectureParserRULE_valueClause

	return p
}

func (s *ValueClauseContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueClauseContext) Value() IValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *ValueClauseContext) AllSPACE() []antlr.TerminalNode {
	return s.GetTokens(LectureParserSPACE)
}

func (s *ValueClauseContext) SPACE(i int) antlr.TerminalNode {
	return s.GetToken(LectureParserSPACE, i)
}

func (s *ValueClauseContext) Operator() IOperatorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOperatorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOperatorContext)
}

func (s *ValueClauseContext) ValueClause() IValueClauseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueClauseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueClauseContext)
}

func (s *ValueClauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueClauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValueClauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LectureListener); ok {
		listenerT.EnterValueClause(s)
	}
}

func (s *ValueClauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LectureListener); ok {
		listenerT.ExitValueClause(s)
	}
}

func (p *LectureParser) ValueClause() (localctx IValueClauseContext) {
	localctx = NewValueClauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, LectureParserRULE_valueClause)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(65)
		p.Value()
	}
	p.SetState(71)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == LectureParserSPACE {
		{
			p.SetState(66)
			p.Match(LectureParserSPACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(67)
			p.Operator()
		}
		{
			p.SetState(68)
			p.Match(LectureParserSPACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(69)
			p.ValueClause()
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

// IValueContext is an interface to support dynamic dispatch.
type IValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LiteralClause() ILiteralClauseContext
	Identifier() IIdentifierContext

	// IsValueContext differentiates from other interfaces.
	IsValueContext()
}

type ValueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValueContext() *ValueContext {
	var p = new(ValueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LectureParserRULE_value
	return p
}

func InitEmptyValueContext(p *ValueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LectureParserRULE_value
}

func (*ValueContext) IsValueContext() {}

func NewValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueContext {
	var p = new(ValueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LectureParserRULE_value

	return p
}

func (s *ValueContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueContext) LiteralClause() ILiteralClauseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILiteralClauseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILiteralClauseContext)
}

func (s *ValueContext) Identifier() IIdentifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *ValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LectureListener); ok {
		listenerT.EnterValue(s)
	}
}

func (s *ValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LectureListener); ok {
		listenerT.ExitValue(s)
	}
}

func (p *LectureParser) Value() (localctx IValueContext) {
	localctx = NewValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, LectureParserRULE_value)
	p.SetState(75)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case LectureParserLITERALLY:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(73)
			p.LiteralClause()
		}

	case LectureParserALPHANUMERICSTRING:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(74)
			p.Identifier()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
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

// ILiteralClauseContext is an interface to support dynamic dispatch.
type ILiteralClauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LITERALLY() antlr.TerminalNode
	SPACE() antlr.TerminalNode
	Literal() ILiteralContext

	// IsLiteralClauseContext differentiates from other interfaces.
	IsLiteralClauseContext()
}

type LiteralClauseContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLiteralClauseContext() *LiteralClauseContext {
	var p = new(LiteralClauseContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LectureParserRULE_literalClause
	return p
}

func InitEmptyLiteralClauseContext(p *LiteralClauseContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LectureParserRULE_literalClause
}

func (*LiteralClauseContext) IsLiteralClauseContext() {}

func NewLiteralClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LiteralClauseContext {
	var p = new(LiteralClauseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LectureParserRULE_literalClause

	return p
}

func (s *LiteralClauseContext) GetParser() antlr.Parser { return s.parser }

func (s *LiteralClauseContext) LITERALLY() antlr.TerminalNode {
	return s.GetToken(LectureParserLITERALLY, 0)
}

func (s *LiteralClauseContext) SPACE() antlr.TerminalNode {
	return s.GetToken(LectureParserSPACE, 0)
}

func (s *LiteralClauseContext) Literal() ILiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILiteralContext)
}

func (s *LiteralClauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralClauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LiteralClauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LectureListener); ok {
		listenerT.EnterLiteralClause(s)
	}
}

func (s *LiteralClauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LectureListener); ok {
		listenerT.ExitLiteralClause(s)
	}
}

func (p *LectureParser) LiteralClause() (localctx ILiteralClauseContext) {
	localctx = NewLiteralClauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, LectureParserRULE_literalClause)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(77)
		p.Match(LectureParserLITERALLY)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(78)
		p.Match(LectureParserSPACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(79)
		p.Literal()
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

// ILiteralContext is an interface to support dynamic dispatch.
type ILiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Number() INumberContext

	// IsLiteralContext differentiates from other interfaces.
	IsLiteralContext()
}

type LiteralContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLiteralContext() *LiteralContext {
	var p = new(LiteralContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LectureParserRULE_literal
	return p
}

func InitEmptyLiteralContext(p *LiteralContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LectureParserRULE_literal
}

func (*LiteralContext) IsLiteralContext() {}

func NewLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LiteralContext {
	var p = new(LiteralContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LectureParserRULE_literal

	return p
}

func (s *LiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *LiteralContext) Number() INumberContext {
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

func (s *LiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LectureListener); ok {
		listenerT.EnterLiteral(s)
	}
}

func (s *LiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LectureListener); ok {
		listenerT.ExitLiteral(s)
	}
}

func (p *LectureParser) Literal() (localctx ILiteralContext) {
	localctx = NewLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, LectureParserRULE_literal)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(81)
		p.Number()
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

// IOperatorContext is an interface to support dynamic dispatch.
type IOperatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PLUS() antlr.TerminalNode

	// IsOperatorContext differentiates from other interfaces.
	IsOperatorContext()
}

type OperatorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOperatorContext() *OperatorContext {
	var p = new(OperatorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LectureParserRULE_operator
	return p
}

func InitEmptyOperatorContext(p *OperatorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LectureParserRULE_operator
}

func (*OperatorContext) IsOperatorContext() {}

func NewOperatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OperatorContext {
	var p = new(OperatorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LectureParserRULE_operator

	return p
}

func (s *OperatorContext) GetParser() antlr.Parser { return s.parser }

func (s *OperatorContext) PLUS() antlr.TerminalNode {
	return s.GetToken(LectureParserPLUS, 0)
}

func (s *OperatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OperatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OperatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LectureListener); ok {
		listenerT.EnterOperator(s)
	}
}

func (s *OperatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LectureListener); ok {
		listenerT.ExitOperator(s)
	}
}

func (p *LectureParser) Operator() (localctx IOperatorContext) {
	localctx = NewOperatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, LectureParserRULE_operator)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(83)
		p.Match(LectureParserPLUS)
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
	INTEGER() antlr.TerminalNode

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
	p.RuleIndex = LectureParserRULE_number
	return p
}

func InitEmptyNumberContext(p *NumberContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LectureParserRULE_number
}

func (*NumberContext) IsNumberContext() {}

func NewNumberContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NumberContext {
	var p = new(NumberContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LectureParserRULE_number

	return p
}

func (s *NumberContext) GetParser() antlr.Parser { return s.parser }

func (s *NumberContext) INTEGER() antlr.TerminalNode {
	return s.GetToken(LectureParserINTEGER, 0)
}

func (s *NumberContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumberContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NumberContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LectureListener); ok {
		listenerT.EnterNumber(s)
	}
}

func (s *NumberContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LectureListener); ok {
		listenerT.ExitNumber(s)
	}
}

func (p *LectureParser) Number() (localctx INumberContext) {
	localctx = NewNumberContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, LectureParserRULE_number)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(85)
		p.Match(LectureParserINTEGER)
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
