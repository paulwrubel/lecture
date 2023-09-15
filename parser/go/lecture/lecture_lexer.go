// Code generated from Lecture.g4 by ANTLR 4.13.1. DO NOT EDIT.

package lecture

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

type LectureLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var LectureLexerLexerStaticData struct {
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

func lecturelexerLexerInit() {
	staticData := &LectureLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
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
		"NON_QUOTE", "ESCAPED_QUOTE", "ALPHANUMERIC", "ALPHA", "BASICSYMBOL",
		"NONQUOTEORESCAPED", "DOT", "DIGIT", "START_CLAUSE", "END_CLAUSE", "DECLARATION_STATEMENT_INTRO",
		"PRINT_STATEMENT_INTRO", "TERMINATOR", "IS", "LITERALLY", "PLUS", "SPACE",
		"ALPHANUMERICSTRING", "INTEGER", "WS",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 12, 157, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 1, 0, 1,
		0, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 5, 1, 5, 3,
		5, 55, 8, 5, 1, 6, 1, 6, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8,
		1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8,
		1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9,
		1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10,
		1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1,
		11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 13, 1, 13, 1, 13, 1, 14,
		1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 15, 1,
		15, 1, 15, 1, 15, 1, 15, 1, 16, 1, 16, 1, 17, 1, 17, 5, 17, 141, 8, 17,
		10, 17, 12, 17, 144, 9, 17, 1, 18, 4, 18, 147, 8, 18, 11, 18, 12, 18, 148,
		1, 19, 4, 19, 152, 8, 19, 11, 19, 12, 19, 153, 1, 19, 1, 19, 0, 0, 20,
		1, 0, 3, 0, 5, 0, 7, 0, 9, 0, 11, 0, 13, 0, 15, 0, 17, 1, 19, 2, 21, 3,
		23, 4, 25, 5, 27, 6, 29, 7, 31, 8, 33, 9, 35, 10, 37, 11, 39, 12, 1, 0,
		6, 1, 0, 34, 34, 3, 0, 48, 57, 65, 90, 97, 122, 2, 0, 65, 90, 97, 122,
		3, 0, 45, 46, 64, 64, 95, 95, 1, 0, 48, 57, 3, 0, 9, 10, 13, 13, 32, 32,
		152, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1,
		0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31,
		1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0,
		39, 1, 0, 0, 0, 1, 41, 1, 0, 0, 0, 3, 43, 1, 0, 0, 0, 5, 46, 1, 0, 0, 0,
		7, 48, 1, 0, 0, 0, 9, 50, 1, 0, 0, 0, 11, 54, 1, 0, 0, 0, 13, 56, 1, 0,
		0, 0, 15, 58, 1, 0, 0, 0, 17, 60, 1, 0, 0, 0, 19, 78, 1, 0, 0, 0, 21, 93,
		1, 0, 0, 0, 23, 103, 1, 0, 0, 0, 25, 116, 1, 0, 0, 0, 27, 118, 1, 0, 0,
		0, 29, 121, 1, 0, 0, 0, 31, 131, 1, 0, 0, 0, 33, 136, 1, 0, 0, 0, 35, 138,
		1, 0, 0, 0, 37, 146, 1, 0, 0, 0, 39, 151, 1, 0, 0, 0, 41, 42, 8, 0, 0,
		0, 42, 2, 1, 0, 0, 0, 43, 44, 5, 92, 0, 0, 44, 45, 5, 34, 0, 0, 45, 4,
		1, 0, 0, 0, 46, 47, 7, 1, 0, 0, 47, 6, 1, 0, 0, 0, 48, 49, 7, 2, 0, 0,
		49, 8, 1, 0, 0, 0, 50, 51, 7, 3, 0, 0, 51, 10, 1, 0, 0, 0, 52, 55, 3, 1,
		0, 0, 53, 55, 3, 3, 1, 0, 54, 52, 1, 0, 0, 0, 54, 53, 1, 0, 0, 0, 55, 12,
		1, 0, 0, 0, 56, 57, 5, 46, 0, 0, 57, 14, 1, 0, 0, 0, 58, 59, 7, 4, 0, 0,
		59, 16, 1, 0, 0, 0, 60, 61, 5, 111, 0, 0, 61, 62, 5, 107, 0, 0, 62, 63,
		5, 97, 0, 0, 63, 64, 5, 121, 0, 0, 64, 65, 5, 44, 0, 0, 65, 66, 5, 32,
		0, 0, 66, 67, 5, 104, 0, 0, 67, 68, 5, 101, 0, 0, 68, 69, 5, 97, 0, 0,
		69, 70, 5, 114, 0, 0, 70, 71, 5, 32, 0, 0, 71, 72, 5, 109, 0, 0, 72, 73,
		5, 101, 0, 0, 73, 74, 5, 32, 0, 0, 74, 75, 5, 111, 0, 0, 75, 76, 5, 117,
		0, 0, 76, 77, 5, 116, 0, 0, 77, 18, 1, 0, 0, 0, 78, 79, 5, 105, 0, 0, 79,
		80, 5, 32, 0, 0, 80, 81, 5, 114, 0, 0, 81, 82, 5, 101, 0, 0, 82, 83, 5,
		115, 0, 0, 83, 84, 5, 116, 0, 0, 84, 85, 5, 32, 0, 0, 85, 86, 5, 109, 0,
		0, 86, 87, 5, 121, 0, 0, 87, 88, 5, 32, 0, 0, 88, 89, 5, 99, 0, 0, 89,
		90, 5, 97, 0, 0, 90, 91, 5, 115, 0, 0, 91, 92, 5, 101, 0, 0, 92, 20, 1,
		0, 0, 0, 93, 94, 5, 108, 0, 0, 94, 95, 5, 101, 0, 0, 95, 96, 5, 116, 0,
		0, 96, 97, 5, 39, 0, 0, 97, 98, 5, 115, 0, 0, 98, 99, 5, 32, 0, 0, 99,
		100, 5, 115, 0, 0, 100, 101, 5, 97, 0, 0, 101, 102, 5, 121, 0, 0, 102,
		22, 1, 0, 0, 0, 103, 104, 5, 116, 0, 0, 104, 105, 5, 104, 0, 0, 105, 106,
		5, 101, 0, 0, 106, 107, 5, 110, 0, 0, 107, 108, 5, 32, 0, 0, 108, 109,
		5, 119, 0, 0, 109, 110, 5, 101, 0, 0, 110, 111, 5, 32, 0, 0, 111, 112,
		5, 104, 0, 0, 112, 113, 5, 97, 0, 0, 113, 114, 5, 118, 0, 0, 114, 115,
		5, 101, 0, 0, 115, 24, 1, 0, 0, 0, 116, 117, 3, 13, 6, 0, 117, 26, 1, 0,
		0, 0, 118, 119, 5, 105, 0, 0, 119, 120, 5, 115, 0, 0, 120, 28, 1, 0, 0,
		0, 121, 122, 5, 108, 0, 0, 122, 123, 5, 105, 0, 0, 123, 124, 5, 116, 0,
		0, 124, 125, 5, 101, 0, 0, 125, 126, 5, 114, 0, 0, 126, 127, 5, 97, 0,
		0, 127, 128, 5, 108, 0, 0, 128, 129, 5, 108, 0, 0, 129, 130, 5, 121, 0,
		0, 130, 30, 1, 0, 0, 0, 131, 132, 5, 112, 0, 0, 132, 133, 5, 108, 0, 0,
		133, 134, 5, 117, 0, 0, 134, 135, 5, 115, 0, 0, 135, 32, 1, 0, 0, 0, 136,
		137, 5, 32, 0, 0, 137, 34, 1, 0, 0, 0, 138, 142, 3, 7, 3, 0, 139, 141,
		3, 5, 2, 0, 140, 139, 1, 0, 0, 0, 141, 144, 1, 0, 0, 0, 142, 140, 1, 0,
		0, 0, 142, 143, 1, 0, 0, 0, 143, 36, 1, 0, 0, 0, 144, 142, 1, 0, 0, 0,
		145, 147, 3, 15, 7, 0, 146, 145, 1, 0, 0, 0, 147, 148, 1, 0, 0, 0, 148,
		146, 1, 0, 0, 0, 148, 149, 1, 0, 0, 0, 149, 38, 1, 0, 0, 0, 150, 152, 7,
		5, 0, 0, 151, 150, 1, 0, 0, 0, 152, 153, 1, 0, 0, 0, 153, 151, 1, 0, 0,
		0, 153, 154, 1, 0, 0, 0, 154, 155, 1, 0, 0, 0, 155, 156, 6, 19, 0, 0, 156,
		40, 1, 0, 0, 0, 5, 0, 54, 142, 148, 153, 1, 6, 0, 0,
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

// LectureLexerInit initializes any static state used to implement LectureLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewLectureLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func LectureLexerInit() {
	staticData := &LectureLexerLexerStaticData
	staticData.once.Do(lecturelexerLexerInit)
}

// NewLectureLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewLectureLexer(input antlr.CharStream) *LectureLexer {
	LectureLexerInit()
	l := new(LectureLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &LectureLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "Lecture.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// LectureLexer tokens.
const (
	LectureLexerSTART_CLAUSE                = 1
	LectureLexerEND_CLAUSE                  = 2
	LectureLexerDECLARATION_STATEMENT_INTRO = 3
	LectureLexerPRINT_STATEMENT_INTRO       = 4
	LectureLexerTERMINATOR                  = 5
	LectureLexerIS                          = 6
	LectureLexerLITERALLY                   = 7
	LectureLexerPLUS                        = 8
	LectureLexerSPACE                       = 9
	LectureLexerALPHANUMERICSTRING          = 10
	LectureLexerINTEGER                     = 11
	LectureLexerWS                          = 12
)
