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
		"'we can use a process known as'", "'which needs'", "'and proceeds as follows'",
		"'finally, we get'", "'the result of'", "", "'a'", "'called'", "'is'",
		"'and'", "'with'", "'literally'", "'number'", "'plus'", "','", "' '",
	}
	staticData.SymbolicNames = []string{
		"", "OKAY_HEAR_ME_OUT", "I_REST_MY_CASE", "LETS_SAY", "THEN_WE_HAVE",
		"WE_CAN_USE_A_PROCESS_KNOWN_AS", "WHICH_NEEDS", "AND_PROCEEDS_AS_FOLLOWS",
		"FINALLY_WE_GET", "THE_RESULT_OF", "TERMINATOR", "A", "CALLED", "IS",
		"AND", "WITH", "LITERALLY", "NUMBER", "PLUS", "COMMA", "SPACE", "ALPHANUMERICSTRING",
		"INTEGER", "WS",
	}
	staticData.RuleNames = []string{
		"lecture", "program", "startClause", "endClause", "statement", "atomicStatement",
		"function", "functionStatement", "returnStatement", "declarationStatement",
		"printStatement", "parametersDeclaration", "parameterDeclarationClause",
		"functionCall", "parametersClause", "valueClause", "literalClause",
		"parameterDeclaration", "value", "literal", "type", "operator", "identifier",
		"number",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 23, 177, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 4,
		1, 54, 8, 1, 11, 1, 12, 1, 55, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3,
		1, 3, 1, 4, 1, 4, 3, 4, 68, 8, 4, 1, 5, 1, 5, 3, 5, 72, 8, 5, 1, 5, 1,
		5, 1, 6, 1, 6, 5, 6, 78, 8, 6, 10, 6, 12, 6, 81, 9, 6, 1, 6, 1, 6, 1, 7,
		1, 7, 1, 7, 1, 7, 3, 7, 89, 8, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1,
		8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 10, 1,
		10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 12,
		1, 12, 1, 12, 1, 12, 1, 12, 3, 12, 124, 8, 12, 1, 13, 1, 13, 1, 13, 1,
		13, 1, 13, 1, 13, 1, 13, 3, 13, 133, 8, 13, 1, 14, 1, 14, 1, 14, 1, 14,
		1, 14, 3, 14, 140, 8, 14, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 3,
		15, 148, 8, 15, 1, 16, 1, 16, 1, 16, 1, 16, 1, 17, 1, 17, 1, 17, 1, 17,
		1, 17, 1, 17, 1, 17, 1, 17, 1, 18, 1, 18, 1, 18, 3, 18, 165, 8, 18, 1,
		19, 1, 19, 1, 20, 1, 20, 1, 21, 1, 21, 1, 22, 1, 22, 1, 23, 1, 23, 1, 23,
		0, 0, 24, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32,
		34, 36, 38, 40, 42, 44, 46, 0, 0, 163, 0, 48, 1, 0, 0, 0, 2, 51, 1, 0,
		0, 0, 4, 59, 1, 0, 0, 0, 6, 62, 1, 0, 0, 0, 8, 67, 1, 0, 0, 0, 10, 71,
		1, 0, 0, 0, 12, 75, 1, 0, 0, 0, 14, 84, 1, 0, 0, 0, 16, 94, 1, 0, 0, 0,
		18, 99, 1, 0, 0, 0, 20, 107, 1, 0, 0, 0, 22, 111, 1, 0, 0, 0, 24, 118,
		1, 0, 0, 0, 26, 125, 1, 0, 0, 0, 28, 134, 1, 0, 0, 0, 30, 141, 1, 0, 0,
		0, 32, 149, 1, 0, 0, 0, 34, 153, 1, 0, 0, 0, 36, 164, 1, 0, 0, 0, 38, 166,
		1, 0, 0, 0, 40, 168, 1, 0, 0, 0, 42, 170, 1, 0, 0, 0, 44, 172, 1, 0, 0,
		0, 46, 174, 1, 0, 0, 0, 48, 49, 3, 2, 1, 0, 49, 50, 5, 0, 0, 1, 50, 1,
		1, 0, 0, 0, 51, 53, 3, 4, 2, 0, 52, 54, 3, 8, 4, 0, 53, 52, 1, 0, 0, 0,
		54, 55, 1, 0, 0, 0, 55, 53, 1, 0, 0, 0, 55, 56, 1, 0, 0, 0, 56, 57, 1,
		0, 0, 0, 57, 58, 3, 6, 3, 0, 58, 3, 1, 0, 0, 0, 59, 60, 5, 1, 0, 0, 60,
		61, 5, 10, 0, 0, 61, 5, 1, 0, 0, 0, 62, 63, 5, 2, 0, 0, 63, 64, 5, 10,
		0, 0, 64, 7, 1, 0, 0, 0, 65, 68, 3, 10, 5, 0, 66, 68, 3, 12, 6, 0, 67,
		65, 1, 0, 0, 0, 67, 66, 1, 0, 0, 0, 68, 9, 1, 0, 0, 0, 69, 72, 3, 18, 9,
		0, 70, 72, 3, 20, 10, 0, 71, 69, 1, 0, 0, 0, 71, 70, 1, 0, 0, 0, 72, 73,
		1, 0, 0, 0, 73, 74, 5, 10, 0, 0, 74, 11, 1, 0, 0, 0, 75, 79, 3, 14, 7,
		0, 76, 78, 3, 10, 5, 0, 77, 76, 1, 0, 0, 0, 78, 81, 1, 0, 0, 0, 79, 77,
		1, 0, 0, 0, 79, 80, 1, 0, 0, 0, 80, 82, 1, 0, 0, 0, 81, 79, 1, 0, 0, 0,
		82, 83, 3, 16, 8, 0, 83, 13, 1, 0, 0, 0, 84, 85, 5, 5, 0, 0, 85, 86, 5,
		20, 0, 0, 86, 88, 3, 44, 22, 0, 87, 89, 3, 22, 11, 0, 88, 87, 1, 0, 0,
		0, 88, 89, 1, 0, 0, 0, 89, 90, 1, 0, 0, 0, 90, 91, 5, 20, 0, 0, 91, 92,
		5, 7, 0, 0, 92, 93, 5, 10, 0, 0, 93, 15, 1, 0, 0, 0, 94, 95, 5, 8, 0, 0,
		95, 96, 5, 20, 0, 0, 96, 97, 3, 30, 15, 0, 97, 98, 5, 10, 0, 0, 98, 17,
		1, 0, 0, 0, 99, 100, 5, 3, 0, 0, 100, 101, 5, 20, 0, 0, 101, 102, 3, 44,
		22, 0, 102, 103, 5, 20, 0, 0, 103, 104, 5, 13, 0, 0, 104, 105, 5, 20, 0,
		0, 105, 106, 3, 30, 15, 0, 106, 19, 1, 0, 0, 0, 107, 108, 5, 4, 0, 0, 108,
		109, 5, 20, 0, 0, 109, 110, 3, 30, 15, 0, 110, 21, 1, 0, 0, 0, 111, 112,
		5, 19, 0, 0, 112, 113, 5, 20, 0, 0, 113, 114, 5, 6, 0, 0, 114, 115, 5,
		20, 0, 0, 115, 116, 3, 24, 12, 0, 116, 117, 5, 19, 0, 0, 117, 23, 1, 0,
		0, 0, 118, 123, 3, 34, 17, 0, 119, 120, 5, 20, 0, 0, 120, 121, 5, 14, 0,
		0, 121, 122, 5, 20, 0, 0, 122, 124, 3, 24, 12, 0, 123, 119, 1, 0, 0, 0,
		123, 124, 1, 0, 0, 0, 124, 25, 1, 0, 0, 0, 125, 126, 5, 9, 0, 0, 126, 127,
		5, 20, 0, 0, 127, 132, 3, 44, 22, 0, 128, 129, 5, 20, 0, 0, 129, 130, 5,
		15, 0, 0, 130, 131, 5, 20, 0, 0, 131, 133, 3, 28, 14, 0, 132, 128, 1, 0,
		0, 0, 132, 133, 1, 0, 0, 0, 133, 27, 1, 0, 0, 0, 134, 139, 3, 36, 18, 0,
		135, 136, 5, 20, 0, 0, 136, 137, 5, 14, 0, 0, 137, 138, 5, 20, 0, 0, 138,
		140, 3, 28, 14, 0, 139, 135, 1, 0, 0, 0, 139, 140, 1, 0, 0, 0, 140, 29,
		1, 0, 0, 0, 141, 147, 3, 36, 18, 0, 142, 143, 5, 20, 0, 0, 143, 144, 3,
		42, 21, 0, 144, 145, 5, 20, 0, 0, 145, 146, 3, 30, 15, 0, 146, 148, 1,
		0, 0, 0, 147, 142, 1, 0, 0, 0, 147, 148, 1, 0, 0, 0, 148, 31, 1, 0, 0,
		0, 149, 150, 5, 16, 0, 0, 150, 151, 5, 20, 0, 0, 151, 152, 3, 38, 19, 0,
		152, 33, 1, 0, 0, 0, 153, 154, 5, 11, 0, 0, 154, 155, 5, 20, 0, 0, 155,
		156, 3, 40, 20, 0, 156, 157, 5, 20, 0, 0, 157, 158, 5, 12, 0, 0, 158, 159,
		5, 20, 0, 0, 159, 160, 3, 44, 22, 0, 160, 35, 1, 0, 0, 0, 161, 165, 3,
		32, 16, 0, 162, 165, 3, 44, 22, 0, 163, 165, 3, 26, 13, 0, 164, 161, 1,
		0, 0, 0, 164, 162, 1, 0, 0, 0, 164, 163, 1, 0, 0, 0, 165, 37, 1, 0, 0,
		0, 166, 167, 3, 46, 23, 0, 167, 39, 1, 0, 0, 0, 168, 169, 5, 17, 0, 0,
		169, 41, 1, 0, 0, 0, 170, 171, 5, 18, 0, 0, 171, 43, 1, 0, 0, 0, 172, 173,
		5, 21, 0, 0, 173, 45, 1, 0, 0, 0, 174, 175, 5, 22, 0, 0, 175, 47, 1, 0,
		0, 0, 10, 55, 67, 71, 79, 88, 123, 132, 139, 147, 164,
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
	LectureParserEOF                           = antlr.TokenEOF
	LectureParserOKAY_HEAR_ME_OUT              = 1
	LectureParserI_REST_MY_CASE                = 2
	LectureParserLETS_SAY                      = 3
	LectureParserTHEN_WE_HAVE                  = 4
	LectureParserWE_CAN_USE_A_PROCESS_KNOWN_AS = 5
	LectureParserWHICH_NEEDS                   = 6
	LectureParserAND_PROCEEDS_AS_FOLLOWS       = 7
	LectureParserFINALLY_WE_GET                = 8
	LectureParserTHE_RESULT_OF                 = 9
	LectureParserTERMINATOR                    = 10
	LectureParserA                             = 11
	LectureParserCALLED                        = 12
	LectureParserIS                            = 13
	LectureParserAND                           = 14
	LectureParserWITH                          = 15
	LectureParserLITERALLY                     = 16
	LectureParserNUMBER                        = 17
	LectureParserPLUS                          = 18
	LectureParserCOMMA                         = 19
	LectureParserSPACE                         = 20
	LectureParserALPHANUMERICSTRING            = 21
	LectureParserINTEGER                       = 22
	LectureParserWS                            = 23
)

// LectureParser rules.
const (
	LectureParserRULE_lecture                    = 0
	LectureParserRULE_program                    = 1
	LectureParserRULE_startClause                = 2
	LectureParserRULE_endClause                  = 3
	LectureParserRULE_statement                  = 4
	LectureParserRULE_atomicStatement            = 5
	LectureParserRULE_function                   = 6
	LectureParserRULE_functionStatement          = 7
	LectureParserRULE_returnStatement            = 8
	LectureParserRULE_declarationStatement       = 9
	LectureParserRULE_printStatement             = 10
	LectureParserRULE_parametersDeclaration      = 11
	LectureParserRULE_parameterDeclarationClause = 12
	LectureParserRULE_functionCall               = 13
	LectureParserRULE_parametersClause           = 14
	LectureParserRULE_valueClause                = 15
	LectureParserRULE_literalClause              = 16
	LectureParserRULE_parameterDeclaration       = 17
	LectureParserRULE_value                      = 18
	LectureParserRULE_literal                    = 19
	LectureParserRULE_type                       = 20
	LectureParserRULE_operator                   = 21
	LectureParserRULE_identifier                 = 22
	LectureParserRULE_number                     = 23
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
		p.SetState(48)
		p.Program()
	}
	{
		p.SetState(49)
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
		p.SetState(51)
		p.StartClause()
	}
	p.SetState(53)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&56) != 0) {
		{
			p.SetState(52)
			p.Statement()
		}

		p.SetState(55)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(57)
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
	OKAY_HEAR_ME_OUT() antlr.TerminalNode
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

func (s *StartClauseContext) OKAY_HEAR_ME_OUT() antlr.TerminalNode {
	return s.GetToken(LectureParserOKAY_HEAR_ME_OUT, 0)
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
		p.SetState(59)
		p.Match(LectureParserOKAY_HEAR_ME_OUT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(60)
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
	I_REST_MY_CASE() antlr.TerminalNode
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

func (s *EndClauseContext) I_REST_MY_CASE() antlr.TerminalNode {
	return s.GetToken(LectureParserI_REST_MY_CASE, 0)
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
		p.SetState(62)
		p.Match(LectureParserI_REST_MY_CASE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(63)
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
	AtomicStatement() IAtomicStatementContext
	Function() IFunctionContext

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

func (s *StatementContext) AtomicStatement() IAtomicStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAtomicStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAtomicStatementContext)
}

func (s *StatementContext) Function() IFunctionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunctionContext)
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
	p.SetState(67)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case LectureParserLETS_SAY, LectureParserTHEN_WE_HAVE:
		{
			p.SetState(65)
			p.AtomicStatement()
		}

	case LectureParserWE_CAN_USE_A_PROCESS_KNOWN_AS:
		{
			p.SetState(66)
			p.Function()
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

// IAtomicStatementContext is an interface to support dynamic dispatch.
type IAtomicStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TERMINATOR() antlr.TerminalNode
	DeclarationStatement() IDeclarationStatementContext
	PrintStatement() IPrintStatementContext

	// IsAtomicStatementContext differentiates from other interfaces.
	IsAtomicStatementContext()
}

type AtomicStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAtomicStatementContext() *AtomicStatementContext {
	var p = new(AtomicStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LectureParserRULE_atomicStatement
	return p
}

func InitEmptyAtomicStatementContext(p *AtomicStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LectureParserRULE_atomicStatement
}

func (*AtomicStatementContext) IsAtomicStatementContext() {}

func NewAtomicStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AtomicStatementContext {
	var p = new(AtomicStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LectureParserRULE_atomicStatement

	return p
}

func (s *AtomicStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *AtomicStatementContext) TERMINATOR() antlr.TerminalNode {
	return s.GetToken(LectureParserTERMINATOR, 0)
}

func (s *AtomicStatementContext) DeclarationStatement() IDeclarationStatementContext {
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

func (s *AtomicStatementContext) PrintStatement() IPrintStatementContext {
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

func (s *AtomicStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AtomicStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AtomicStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LectureListener); ok {
		listenerT.EnterAtomicStatement(s)
	}
}

func (s *AtomicStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LectureListener); ok {
		listenerT.ExitAtomicStatement(s)
	}
}

func (p *LectureParser) AtomicStatement() (localctx IAtomicStatementContext) {
	localctx = NewAtomicStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, LectureParserRULE_atomicStatement)
	p.EnterOuterAlt(localctx, 1)
	p.SetState(71)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case LectureParserLETS_SAY:
		{
			p.SetState(69)
			p.DeclarationStatement()
		}

	case LectureParserTHEN_WE_HAVE:
		{
			p.SetState(70)
			p.PrintStatement()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	{
		p.SetState(73)
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

// IFunctionContext is an interface to support dynamic dispatch.
type IFunctionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FunctionStatement() IFunctionStatementContext
	ReturnStatement() IReturnStatementContext
	AllAtomicStatement() []IAtomicStatementContext
	AtomicStatement(i int) IAtomicStatementContext

	// IsFunctionContext differentiates from other interfaces.
	IsFunctionContext()
}

type FunctionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionContext() *FunctionContext {
	var p = new(FunctionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LectureParserRULE_function
	return p
}

func InitEmptyFunctionContext(p *FunctionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LectureParserRULE_function
}

func (*FunctionContext) IsFunctionContext() {}

func NewFunctionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionContext {
	var p = new(FunctionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LectureParserRULE_function

	return p
}

func (s *FunctionContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionContext) FunctionStatement() IFunctionStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunctionStatementContext)
}

func (s *FunctionContext) ReturnStatement() IReturnStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReturnStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReturnStatementContext)
}

func (s *FunctionContext) AllAtomicStatement() []IAtomicStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IAtomicStatementContext); ok {
			len++
		}
	}

	tst := make([]IAtomicStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IAtomicStatementContext); ok {
			tst[i] = t.(IAtomicStatementContext)
			i++
		}
	}

	return tst
}

func (s *FunctionContext) AtomicStatement(i int) IAtomicStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAtomicStatementContext); ok {
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

	return t.(IAtomicStatementContext)
}

func (s *FunctionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LectureListener); ok {
		listenerT.EnterFunction(s)
	}
}

func (s *FunctionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LectureListener); ok {
		listenerT.ExitFunction(s)
	}
}

func (p *LectureParser) Function() (localctx IFunctionContext) {
	localctx = NewFunctionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, LectureParserRULE_function)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(75)
		p.FunctionStatement()
	}
	p.SetState(79)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == LectureParserLETS_SAY || _la == LectureParserTHEN_WE_HAVE {
		{
			p.SetState(76)
			p.AtomicStatement()
		}

		p.SetState(81)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(82)
		p.ReturnStatement()
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

// IFunctionStatementContext is an interface to support dynamic dispatch.
type IFunctionStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	WE_CAN_USE_A_PROCESS_KNOWN_AS() antlr.TerminalNode
	AllSPACE() []antlr.TerminalNode
	SPACE(i int) antlr.TerminalNode
	Identifier() IIdentifierContext
	AND_PROCEEDS_AS_FOLLOWS() antlr.TerminalNode
	TERMINATOR() antlr.TerminalNode
	ParametersDeclaration() IParametersDeclarationContext

	// IsFunctionStatementContext differentiates from other interfaces.
	IsFunctionStatementContext()
}

type FunctionStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionStatementContext() *FunctionStatementContext {
	var p = new(FunctionStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LectureParserRULE_functionStatement
	return p
}

func InitEmptyFunctionStatementContext(p *FunctionStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LectureParserRULE_functionStatement
}

func (*FunctionStatementContext) IsFunctionStatementContext() {}

func NewFunctionStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionStatementContext {
	var p = new(FunctionStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LectureParserRULE_functionStatement

	return p
}

func (s *FunctionStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionStatementContext) WE_CAN_USE_A_PROCESS_KNOWN_AS() antlr.TerminalNode {
	return s.GetToken(LectureParserWE_CAN_USE_A_PROCESS_KNOWN_AS, 0)
}

func (s *FunctionStatementContext) AllSPACE() []antlr.TerminalNode {
	return s.GetTokens(LectureParserSPACE)
}

func (s *FunctionStatementContext) SPACE(i int) antlr.TerminalNode {
	return s.GetToken(LectureParserSPACE, i)
}

func (s *FunctionStatementContext) Identifier() IIdentifierContext {
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

func (s *FunctionStatementContext) AND_PROCEEDS_AS_FOLLOWS() antlr.TerminalNode {
	return s.GetToken(LectureParserAND_PROCEEDS_AS_FOLLOWS, 0)
}

func (s *FunctionStatementContext) TERMINATOR() antlr.TerminalNode {
	return s.GetToken(LectureParserTERMINATOR, 0)
}

func (s *FunctionStatementContext) ParametersDeclaration() IParametersDeclarationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParametersDeclarationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParametersDeclarationContext)
}

func (s *FunctionStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LectureListener); ok {
		listenerT.EnterFunctionStatement(s)
	}
}

func (s *FunctionStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LectureListener); ok {
		listenerT.ExitFunctionStatement(s)
	}
}

func (p *LectureParser) FunctionStatement() (localctx IFunctionStatementContext) {
	localctx = NewFunctionStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, LectureParserRULE_functionStatement)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(84)
		p.Match(LectureParserWE_CAN_USE_A_PROCESS_KNOWN_AS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(85)
		p.Match(LectureParserSPACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(86)
		p.Identifier()
	}
	p.SetState(88)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == LectureParserCOMMA {
		{
			p.SetState(87)
			p.ParametersDeclaration()
		}

	}
	{
		p.SetState(90)
		p.Match(LectureParserSPACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(91)
		p.Match(LectureParserAND_PROCEEDS_AS_FOLLOWS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(92)
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

// IReturnStatementContext is an interface to support dynamic dispatch.
type IReturnStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FINALLY_WE_GET() antlr.TerminalNode
	SPACE() antlr.TerminalNode
	ValueClause() IValueClauseContext
	TERMINATOR() antlr.TerminalNode

	// IsReturnStatementContext differentiates from other interfaces.
	IsReturnStatementContext()
}

type ReturnStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReturnStatementContext() *ReturnStatementContext {
	var p = new(ReturnStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LectureParserRULE_returnStatement
	return p
}

func InitEmptyReturnStatementContext(p *ReturnStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LectureParserRULE_returnStatement
}

func (*ReturnStatementContext) IsReturnStatementContext() {}

func NewReturnStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReturnStatementContext {
	var p = new(ReturnStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LectureParserRULE_returnStatement

	return p
}

func (s *ReturnStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *ReturnStatementContext) FINALLY_WE_GET() antlr.TerminalNode {
	return s.GetToken(LectureParserFINALLY_WE_GET, 0)
}

func (s *ReturnStatementContext) SPACE() antlr.TerminalNode {
	return s.GetToken(LectureParserSPACE, 0)
}

func (s *ReturnStatementContext) ValueClause() IValueClauseContext {
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

func (s *ReturnStatementContext) TERMINATOR() antlr.TerminalNode {
	return s.GetToken(LectureParserTERMINATOR, 0)
}

func (s *ReturnStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReturnStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReturnStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LectureListener); ok {
		listenerT.EnterReturnStatement(s)
	}
}

func (s *ReturnStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LectureListener); ok {
		listenerT.ExitReturnStatement(s)
	}
}

func (p *LectureParser) ReturnStatement() (localctx IReturnStatementContext) {
	localctx = NewReturnStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, LectureParserRULE_returnStatement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(94)
		p.Match(LectureParserFINALLY_WE_GET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(95)
		p.Match(LectureParserSPACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(96)
		p.ValueClause()
	}
	{
		p.SetState(97)
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
	LETS_SAY() antlr.TerminalNode
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

func (s *DeclarationStatementContext) LETS_SAY() antlr.TerminalNode {
	return s.GetToken(LectureParserLETS_SAY, 0)
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
	p.EnterRule(localctx, 18, LectureParserRULE_declarationStatement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(99)
		p.Match(LectureParserLETS_SAY)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(100)
		p.Match(LectureParserSPACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(101)
		p.Identifier()
	}
	{
		p.SetState(102)
		p.Match(LectureParserSPACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(103)
		p.Match(LectureParserIS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(104)
		p.Match(LectureParserSPACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(105)
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
	THEN_WE_HAVE() antlr.TerminalNode
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

func (s *PrintStatementContext) THEN_WE_HAVE() antlr.TerminalNode {
	return s.GetToken(LectureParserTHEN_WE_HAVE, 0)
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
	p.EnterRule(localctx, 20, LectureParserRULE_printStatement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(107)
		p.Match(LectureParserTHEN_WE_HAVE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(108)
		p.Match(LectureParserSPACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(109)
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

// IParametersDeclarationContext is an interface to support dynamic dispatch.
type IParametersDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode
	AllSPACE() []antlr.TerminalNode
	SPACE(i int) antlr.TerminalNode
	WHICH_NEEDS() antlr.TerminalNode
	ParameterDeclarationClause() IParameterDeclarationClauseContext

	// IsParametersDeclarationContext differentiates from other interfaces.
	IsParametersDeclarationContext()
}

type ParametersDeclarationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParametersDeclarationContext() *ParametersDeclarationContext {
	var p = new(ParametersDeclarationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LectureParserRULE_parametersDeclaration
	return p
}

func InitEmptyParametersDeclarationContext(p *ParametersDeclarationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LectureParserRULE_parametersDeclaration
}

func (*ParametersDeclarationContext) IsParametersDeclarationContext() {}

func NewParametersDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParametersDeclarationContext {
	var p = new(ParametersDeclarationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LectureParserRULE_parametersDeclaration

	return p
}

func (s *ParametersDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *ParametersDeclarationContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(LectureParserCOMMA)
}

func (s *ParametersDeclarationContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(LectureParserCOMMA, i)
}

func (s *ParametersDeclarationContext) AllSPACE() []antlr.TerminalNode {
	return s.GetTokens(LectureParserSPACE)
}

func (s *ParametersDeclarationContext) SPACE(i int) antlr.TerminalNode {
	return s.GetToken(LectureParserSPACE, i)
}

func (s *ParametersDeclarationContext) WHICH_NEEDS() antlr.TerminalNode {
	return s.GetToken(LectureParserWHICH_NEEDS, 0)
}

func (s *ParametersDeclarationContext) ParameterDeclarationClause() IParameterDeclarationClauseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParameterDeclarationClauseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParameterDeclarationClauseContext)
}

func (s *ParametersDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParametersDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParametersDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LectureListener); ok {
		listenerT.EnterParametersDeclaration(s)
	}
}

func (s *ParametersDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LectureListener); ok {
		listenerT.ExitParametersDeclaration(s)
	}
}

func (p *LectureParser) ParametersDeclaration() (localctx IParametersDeclarationContext) {
	localctx = NewParametersDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, LectureParserRULE_parametersDeclaration)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(111)
		p.Match(LectureParserCOMMA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(112)
		p.Match(LectureParserSPACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(113)
		p.Match(LectureParserWHICH_NEEDS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(114)
		p.Match(LectureParserSPACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(115)
		p.ParameterDeclarationClause()
	}
	{
		p.SetState(116)
		p.Match(LectureParserCOMMA)
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

// IParameterDeclarationClauseContext is an interface to support dynamic dispatch.
type IParameterDeclarationClauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ParameterDeclaration() IParameterDeclarationContext
	AllSPACE() []antlr.TerminalNode
	SPACE(i int) antlr.TerminalNode
	AND() antlr.TerminalNode
	ParameterDeclarationClause() IParameterDeclarationClauseContext

	// IsParameterDeclarationClauseContext differentiates from other interfaces.
	IsParameterDeclarationClauseContext()
}

type ParameterDeclarationClauseContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParameterDeclarationClauseContext() *ParameterDeclarationClauseContext {
	var p = new(ParameterDeclarationClauseContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LectureParserRULE_parameterDeclarationClause
	return p
}

func InitEmptyParameterDeclarationClauseContext(p *ParameterDeclarationClauseContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LectureParserRULE_parameterDeclarationClause
}

func (*ParameterDeclarationClauseContext) IsParameterDeclarationClauseContext() {}

func NewParameterDeclarationClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParameterDeclarationClauseContext {
	var p = new(ParameterDeclarationClauseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LectureParserRULE_parameterDeclarationClause

	return p
}

func (s *ParameterDeclarationClauseContext) GetParser() antlr.Parser { return s.parser }

func (s *ParameterDeclarationClauseContext) ParameterDeclaration() IParameterDeclarationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParameterDeclarationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParameterDeclarationContext)
}

func (s *ParameterDeclarationClauseContext) AllSPACE() []antlr.TerminalNode {
	return s.GetTokens(LectureParserSPACE)
}

func (s *ParameterDeclarationClauseContext) SPACE(i int) antlr.TerminalNode {
	return s.GetToken(LectureParserSPACE, i)
}

func (s *ParameterDeclarationClauseContext) AND() antlr.TerminalNode {
	return s.GetToken(LectureParserAND, 0)
}

func (s *ParameterDeclarationClauseContext) ParameterDeclarationClause() IParameterDeclarationClauseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParameterDeclarationClauseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParameterDeclarationClauseContext)
}

func (s *ParameterDeclarationClauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParameterDeclarationClauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParameterDeclarationClauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LectureListener); ok {
		listenerT.EnterParameterDeclarationClause(s)
	}
}

func (s *ParameterDeclarationClauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LectureListener); ok {
		listenerT.ExitParameterDeclarationClause(s)
	}
}

func (p *LectureParser) ParameterDeclarationClause() (localctx IParameterDeclarationClauseContext) {
	localctx = NewParameterDeclarationClauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, LectureParserRULE_parameterDeclarationClause)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(118)
		p.ParameterDeclaration()
	}
	p.SetState(123)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == LectureParserSPACE {
		{
			p.SetState(119)
			p.Match(LectureParserSPACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(120)
			p.Match(LectureParserAND)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(121)
			p.Match(LectureParserSPACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(122)
			p.ParameterDeclarationClause()
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

// IFunctionCallContext is an interface to support dynamic dispatch.
type IFunctionCallContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	THE_RESULT_OF() antlr.TerminalNode
	AllSPACE() []antlr.TerminalNode
	SPACE(i int) antlr.TerminalNode
	Identifier() IIdentifierContext
	WITH() antlr.TerminalNode
	ParametersClause() IParametersClauseContext

	// IsFunctionCallContext differentiates from other interfaces.
	IsFunctionCallContext()
}

type FunctionCallContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionCallContext() *FunctionCallContext {
	var p = new(FunctionCallContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LectureParserRULE_functionCall
	return p
}

func InitEmptyFunctionCallContext(p *FunctionCallContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LectureParserRULE_functionCall
}

func (*FunctionCallContext) IsFunctionCallContext() {}

func NewFunctionCallContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionCallContext {
	var p = new(FunctionCallContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LectureParserRULE_functionCall

	return p
}

func (s *FunctionCallContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionCallContext) THE_RESULT_OF() antlr.TerminalNode {
	return s.GetToken(LectureParserTHE_RESULT_OF, 0)
}

func (s *FunctionCallContext) AllSPACE() []antlr.TerminalNode {
	return s.GetTokens(LectureParserSPACE)
}

func (s *FunctionCallContext) SPACE(i int) antlr.TerminalNode {
	return s.GetToken(LectureParserSPACE, i)
}

func (s *FunctionCallContext) Identifier() IIdentifierContext {
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

func (s *FunctionCallContext) WITH() antlr.TerminalNode {
	return s.GetToken(LectureParserWITH, 0)
}

func (s *FunctionCallContext) ParametersClause() IParametersClauseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParametersClauseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParametersClauseContext)
}

func (s *FunctionCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionCallContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionCallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LectureListener); ok {
		listenerT.EnterFunctionCall(s)
	}
}

func (s *FunctionCallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LectureListener); ok {
		listenerT.ExitFunctionCall(s)
	}
}

func (p *LectureParser) FunctionCall() (localctx IFunctionCallContext) {
	localctx = NewFunctionCallContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, LectureParserRULE_functionCall)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(125)
		p.Match(LectureParserTHE_RESULT_OF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(126)
		p.Match(LectureParserSPACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(127)
		p.Identifier()
	}
	p.SetState(132)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(128)
			p.Match(LectureParserSPACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(129)
			p.Match(LectureParserWITH)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(130)
			p.Match(LectureParserSPACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(131)
			p.ParametersClause()
		}

	} else if p.HasError() { // JIM
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

// IParametersClauseContext is an interface to support dynamic dispatch.
type IParametersClauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Value() IValueContext
	AllSPACE() []antlr.TerminalNode
	SPACE(i int) antlr.TerminalNode
	AND() antlr.TerminalNode
	ParametersClause() IParametersClauseContext

	// IsParametersClauseContext differentiates from other interfaces.
	IsParametersClauseContext()
}

type ParametersClauseContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParametersClauseContext() *ParametersClauseContext {
	var p = new(ParametersClauseContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LectureParserRULE_parametersClause
	return p
}

func InitEmptyParametersClauseContext(p *ParametersClauseContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LectureParserRULE_parametersClause
}

func (*ParametersClauseContext) IsParametersClauseContext() {}

func NewParametersClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParametersClauseContext {
	var p = new(ParametersClauseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LectureParserRULE_parametersClause

	return p
}

func (s *ParametersClauseContext) GetParser() antlr.Parser { return s.parser }

func (s *ParametersClauseContext) Value() IValueContext {
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

func (s *ParametersClauseContext) AllSPACE() []antlr.TerminalNode {
	return s.GetTokens(LectureParserSPACE)
}

func (s *ParametersClauseContext) SPACE(i int) antlr.TerminalNode {
	return s.GetToken(LectureParserSPACE, i)
}

func (s *ParametersClauseContext) AND() antlr.TerminalNode {
	return s.GetToken(LectureParserAND, 0)
}

func (s *ParametersClauseContext) ParametersClause() IParametersClauseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParametersClauseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParametersClauseContext)
}

func (s *ParametersClauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParametersClauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParametersClauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LectureListener); ok {
		listenerT.EnterParametersClause(s)
	}
}

func (s *ParametersClauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LectureListener); ok {
		listenerT.ExitParametersClause(s)
	}
}

func (p *LectureParser) ParametersClause() (localctx IParametersClauseContext) {
	localctx = NewParametersClauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, LectureParserRULE_parametersClause)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(134)
		p.Value()
	}
	p.SetState(139)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 7, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(135)
			p.Match(LectureParserSPACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(136)
			p.Match(LectureParserAND)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(137)
			p.Match(LectureParserSPACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(138)
			p.ParametersClause()
		}

	} else if p.HasError() { // JIM
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
	p.EnterRule(localctx, 30, LectureParserRULE_valueClause)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(141)
		p.Value()
	}
	p.SetState(147)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == LectureParserSPACE {
		{
			p.SetState(142)
			p.Match(LectureParserSPACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(143)
			p.Operator()
		}
		{
			p.SetState(144)
			p.Match(LectureParserSPACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(145)
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
	p.EnterRule(localctx, 32, LectureParserRULE_literalClause)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(149)
		p.Match(LectureParserLITERALLY)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(150)
		p.Match(LectureParserSPACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(151)
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

// IParameterDeclarationContext is an interface to support dynamic dispatch.
type IParameterDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	A() antlr.TerminalNode
	AllSPACE() []antlr.TerminalNode
	SPACE(i int) antlr.TerminalNode
	Type_() ITypeContext
	CALLED() antlr.TerminalNode
	Identifier() IIdentifierContext

	// IsParameterDeclarationContext differentiates from other interfaces.
	IsParameterDeclarationContext()
}

type ParameterDeclarationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParameterDeclarationContext() *ParameterDeclarationContext {
	var p = new(ParameterDeclarationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LectureParserRULE_parameterDeclaration
	return p
}

func InitEmptyParameterDeclarationContext(p *ParameterDeclarationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LectureParserRULE_parameterDeclaration
}

func (*ParameterDeclarationContext) IsParameterDeclarationContext() {}

func NewParameterDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParameterDeclarationContext {
	var p = new(ParameterDeclarationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LectureParserRULE_parameterDeclaration

	return p
}

func (s *ParameterDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *ParameterDeclarationContext) A() antlr.TerminalNode {
	return s.GetToken(LectureParserA, 0)
}

func (s *ParameterDeclarationContext) AllSPACE() []antlr.TerminalNode {
	return s.GetTokens(LectureParserSPACE)
}

func (s *ParameterDeclarationContext) SPACE(i int) antlr.TerminalNode {
	return s.GetToken(LectureParserSPACE, i)
}

func (s *ParameterDeclarationContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *ParameterDeclarationContext) CALLED() antlr.TerminalNode {
	return s.GetToken(LectureParserCALLED, 0)
}

func (s *ParameterDeclarationContext) Identifier() IIdentifierContext {
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

func (s *ParameterDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParameterDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParameterDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LectureListener); ok {
		listenerT.EnterParameterDeclaration(s)
	}
}

func (s *ParameterDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LectureListener); ok {
		listenerT.ExitParameterDeclaration(s)
	}
}

func (p *LectureParser) ParameterDeclaration() (localctx IParameterDeclarationContext) {
	localctx = NewParameterDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, LectureParserRULE_parameterDeclaration)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(153)
		p.Match(LectureParserA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(154)
		p.Match(LectureParserSPACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(155)
		p.Type_()
	}
	{
		p.SetState(156)
		p.Match(LectureParserSPACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(157)
		p.Match(LectureParserCALLED)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(158)
		p.Match(LectureParserSPACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(159)
		p.Identifier()
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
	FunctionCall() IFunctionCallContext

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

func (s *ValueContext) FunctionCall() IFunctionCallContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionCallContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunctionCallContext)
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
	p.EnterRule(localctx, 36, LectureParserRULE_value)
	p.SetState(164)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case LectureParserLITERALLY:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(161)
			p.LiteralClause()
		}

	case LectureParserALPHANUMERICSTRING:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(162)
			p.Identifier()
		}

	case LectureParserTHE_RESULT_OF:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(163)
			p.FunctionCall()
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
	p.EnterRule(localctx, 38, LectureParserRULE_literal)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(166)
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

// ITypeContext is an interface to support dynamic dispatch.
type ITypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NUMBER() antlr.TerminalNode

	// IsTypeContext differentiates from other interfaces.
	IsTypeContext()
}

type TypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeContext() *TypeContext {
	var p = new(TypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LectureParserRULE_type
	return p
}

func InitEmptyTypeContext(p *TypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LectureParserRULE_type
}

func (*TypeContext) IsTypeContext() {}

func NewTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeContext {
	var p = new(TypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LectureParserRULE_type

	return p
}

func (s *TypeContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(LectureParserNUMBER, 0)
}

func (s *TypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LectureListener); ok {
		listenerT.EnterType(s)
	}
}

func (s *TypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LectureListener); ok {
		listenerT.ExitType(s)
	}
}

func (p *LectureParser) Type_() (localctx ITypeContext) {
	localctx = NewTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, LectureParserRULE_type)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(168)
		p.Match(LectureParserNUMBER)
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
	p.EnterRule(localctx, 42, LectureParserRULE_operator)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(170)
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
	p.EnterRule(localctx, 44, LectureParserRULE_identifier)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(172)
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
	p.EnterRule(localctx, 46, LectureParserRULE_number)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(174)
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
