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
		"", "", "'okay, hear me out'", "'i rest my case'", "'now let's say'",
		"'let's say'", "'then we have'", "'we can use a process known as'",
		"'which needs'", "'to produce a'", "'it proceeds as follows'", "'finally, we get'",
		"'the result of'", "'here's what we need to do'", "'now that we've done that'",
		"'we can move on'", "'by the way'", "'a'", "'called'", "'is'", "'and'",
		"'using'", "'if'", "'otherwise'", "'literally'", "'number'", "'plus'",
		"'minus'", "'quote'", "'unquote'", "','", "' '",
	}
	staticData.SymbolicNames = []string{
		"", "COMMENT", "OKAY_HEAR_ME_OUT", "I_REST_MY_CASE", "NOW_LETS_SAY",
		"LETS_SAY", "THEN_WE_HAVE", "WE_CAN_USE_A_PROCESS_KNOWN_AS", "WHICH_NEEDS",
		"TO_PRODUCE_A", "IT_PROCEEDS_AS_FOLLOWS", "FINALLY_WE_GET", "THE_RESULT_OF",
		"HERES_WHAT_WE_NEED_TO_DO", "NOW_THAT_WEVE_DONE_THAT", "WE_CAN_MOVE_ON",
		"BY_THE_WAY", "A", "CALLED", "IS", "AND", "USING", "IF", "OTHERWISE",
		"LITERALLY", "NUMBER", "PLUS", "MINUS", "QUOTE", "UNQUOTE", "COMMA",
		"SPACE", "TERMINATOR", "IDENTIFIER_STRING", "STRING", "INTEGER", "WS",
	}
	staticData.RuleNames = []string{
		"lecture", "program", "floatingComment", "mainFunction", "mainStartStatement",
		"mainEndStatement", "function", "functionSignature", "parametersDeclaration",
		"parameterDeclarationClause", "parameterDeclaration", "statement", "statementBlock",
		"returnStatement", "assignmentStatement", "reassignmentStatement", "printStatement",
		"commentStatement", "ifChainStatement", "ifStatement", "ifSignature",
		"elseIfStatement", "elseIfSignature", "elseStatement", "elseSignature",
		"ifClosingStatement", "conditionClause", "valueClause", "value", "literalClause",
		"functionCall", "parametersClause", "parameter", "type", "operator",
		"comparator", "identifier", "literal", "string", "number",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 36, 306, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36, 7, 36,
		2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1,
		5, 1, 86, 8, 1, 10, 1, 12, 1, 89, 9, 1, 1, 1, 1, 1, 1, 1, 5, 1, 94, 8,
		1, 10, 1, 12, 1, 97, 9, 1, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1,
		4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 5, 6, 114, 8, 6, 10, 6, 12,
		6, 117, 9, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 3, 7, 125, 8, 7, 1, 7,
		1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8,
		1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 3, 9, 147, 8, 9, 1, 10, 1, 10,
		1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1,
		11, 3, 11, 162, 8, 11, 1, 11, 1, 11, 1, 12, 4, 12, 167, 8, 12, 11, 12,
		12, 12, 168, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 1,
		14, 1, 14, 1, 14, 1, 14, 1, 14, 3, 14, 184, 8, 14, 1, 15, 1, 15, 1, 15,
		1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 3, 15, 194, 8, 15, 1, 16, 1, 16, 1,
		16, 1, 16, 1, 17, 1, 17, 1, 18, 1, 18, 5, 18, 204, 8, 18, 10, 18, 12, 18,
		207, 9, 18, 1, 18, 3, 18, 210, 8, 18, 1, 18, 1, 18, 1, 19, 1, 19, 1, 19,
		1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 21, 1, 21, 1,
		21, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 23, 1, 23, 1, 23, 1, 24, 1, 24,
		1, 24, 1, 24, 1, 24, 1, 24, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 26, 1,
		26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27,
		3, 27, 259, 8, 27, 1, 28, 1, 28, 3, 28, 263, 8, 28, 1, 29, 1, 29, 1, 29,
		1, 29, 1, 30, 1, 30, 1, 30, 1, 30, 3, 30, 273, 8, 30, 1, 30, 1, 30, 1,
		30, 1, 30, 3, 30, 279, 8, 30, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 3, 31,
		286, 8, 31, 1, 32, 1, 32, 1, 33, 1, 33, 1, 34, 1, 34, 1, 35, 1, 35, 1,
		36, 1, 36, 1, 37, 1, 37, 3, 37, 300, 8, 37, 1, 38, 1, 38, 1, 39, 1, 39,
		1, 39, 0, 0, 40, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28,
		30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64,
		66, 68, 70, 72, 74, 76, 78, 0, 1, 1, 0, 26, 27, 287, 0, 80, 1, 0, 0, 0,
		2, 87, 1, 0, 0, 0, 4, 98, 1, 0, 0, 0, 6, 101, 1, 0, 0, 0, 8, 105, 1, 0,
		0, 0, 10, 108, 1, 0, 0, 0, 12, 111, 1, 0, 0, 0, 14, 120, 1, 0, 0, 0, 16,
		134, 1, 0, 0, 0, 18, 141, 1, 0, 0, 0, 20, 148, 1, 0, 0, 0, 22, 161, 1,
		0, 0, 0, 24, 166, 1, 0, 0, 0, 26, 170, 1, 0, 0, 0, 28, 175, 1, 0, 0, 0,
		30, 185, 1, 0, 0, 0, 32, 195, 1, 0, 0, 0, 34, 199, 1, 0, 0, 0, 36, 201,
		1, 0, 0, 0, 38, 213, 1, 0, 0, 0, 40, 216, 1, 0, 0, 0, 42, 224, 1, 0, 0,
		0, 44, 227, 1, 0, 0, 0, 46, 232, 1, 0, 0, 0, 48, 235, 1, 0, 0, 0, 50, 241,
		1, 0, 0, 0, 52, 246, 1, 0, 0, 0, 54, 252, 1, 0, 0, 0, 56, 262, 1, 0, 0,
		0, 58, 264, 1, 0, 0, 0, 60, 268, 1, 0, 0, 0, 62, 280, 1, 0, 0, 0, 64, 287,
		1, 0, 0, 0, 66, 289, 1, 0, 0, 0, 68, 291, 1, 0, 0, 0, 70, 293, 1, 0, 0,
		0, 72, 295, 1, 0, 0, 0, 74, 299, 1, 0, 0, 0, 76, 301, 1, 0, 0, 0, 78, 303,
		1, 0, 0, 0, 80, 81, 3, 2, 1, 0, 81, 82, 5, 0, 0, 1, 82, 1, 1, 0, 0, 0,
		83, 86, 3, 4, 2, 0, 84, 86, 3, 12, 6, 0, 85, 83, 1, 0, 0, 0, 85, 84, 1,
		0, 0, 0, 86, 89, 1, 0, 0, 0, 87, 85, 1, 0, 0, 0, 87, 88, 1, 0, 0, 0, 88,
		90, 1, 0, 0, 0, 89, 87, 1, 0, 0, 0, 90, 95, 3, 6, 3, 0, 91, 94, 3, 4, 2,
		0, 92, 94, 3, 12, 6, 0, 93, 91, 1, 0, 0, 0, 93, 92, 1, 0, 0, 0, 94, 97,
		1, 0, 0, 0, 95, 93, 1, 0, 0, 0, 95, 96, 1, 0, 0, 0, 96, 3, 1, 0, 0, 0,
		97, 95, 1, 0, 0, 0, 98, 99, 3, 34, 17, 0, 99, 100, 5, 32, 0, 0, 100, 5,
		1, 0, 0, 0, 101, 102, 3, 8, 4, 0, 102, 103, 3, 24, 12, 0, 103, 104, 3,
		10, 5, 0, 104, 7, 1, 0, 0, 0, 105, 106, 5, 2, 0, 0, 106, 107, 5, 32, 0,
		0, 107, 9, 1, 0, 0, 0, 108, 109, 5, 3, 0, 0, 109, 110, 5, 32, 0, 0, 110,
		11, 1, 0, 0, 0, 111, 115, 3, 14, 7, 0, 112, 114, 3, 22, 11, 0, 113, 112,
		1, 0, 0, 0, 114, 117, 1, 0, 0, 0, 115, 113, 1, 0, 0, 0, 115, 116, 1, 0,
		0, 0, 116, 118, 1, 0, 0, 0, 117, 115, 1, 0, 0, 0, 118, 119, 3, 26, 13,
		0, 119, 13, 1, 0, 0, 0, 120, 121, 5, 7, 0, 0, 121, 122, 5, 31, 0, 0, 122,
		124, 3, 72, 36, 0, 123, 125, 3, 16, 8, 0, 124, 123, 1, 0, 0, 0, 124, 125,
		1, 0, 0, 0, 125, 126, 1, 0, 0, 0, 126, 127, 5, 31, 0, 0, 127, 128, 5, 9,
		0, 0, 128, 129, 5, 31, 0, 0, 129, 130, 3, 66, 33, 0, 130, 131, 5, 32, 0,
		0, 131, 132, 5, 10, 0, 0, 132, 133, 5, 32, 0, 0, 133, 15, 1, 0, 0, 0, 134,
		135, 5, 30, 0, 0, 135, 136, 5, 31, 0, 0, 136, 137, 5, 8, 0, 0, 137, 138,
		5, 31, 0, 0, 138, 139, 3, 18, 9, 0, 139, 140, 5, 30, 0, 0, 140, 17, 1,
		0, 0, 0, 141, 146, 3, 20, 10, 0, 142, 143, 5, 31, 0, 0, 143, 144, 5, 20,
		0, 0, 144, 145, 5, 31, 0, 0, 145, 147, 3, 18, 9, 0, 146, 142, 1, 0, 0,
		0, 146, 147, 1, 0, 0, 0, 147, 19, 1, 0, 0, 0, 148, 149, 5, 17, 0, 0, 149,
		150, 5, 31, 0, 0, 150, 151, 3, 66, 33, 0, 151, 152, 5, 31, 0, 0, 152, 153,
		5, 18, 0, 0, 153, 154, 5, 31, 0, 0, 154, 155, 3, 72, 36, 0, 155, 21, 1,
		0, 0, 0, 156, 162, 3, 28, 14, 0, 157, 162, 3, 30, 15, 0, 158, 162, 3, 32,
		16, 0, 159, 162, 3, 36, 18, 0, 160, 162, 3, 34, 17, 0, 161, 156, 1, 0,
		0, 0, 161, 157, 1, 0, 0, 0, 161, 158, 1, 0, 0, 0, 161, 159, 1, 0, 0, 0,
		161, 160, 1, 0, 0, 0, 162, 163, 1, 0, 0, 0, 163, 164, 5, 32, 0, 0, 164,
		23, 1, 0, 0, 0, 165, 167, 3, 22, 11, 0, 166, 165, 1, 0, 0, 0, 167, 168,
		1, 0, 0, 0, 168, 166, 1, 0, 0, 0, 168, 169, 1, 0, 0, 0, 169, 25, 1, 0,
		0, 0, 170, 171, 5, 11, 0, 0, 171, 172, 5, 31, 0, 0, 172, 173, 3, 54, 27,
		0, 173, 174, 5, 32, 0, 0, 174, 27, 1, 0, 0, 0, 175, 176, 5, 5, 0, 0, 176,
		177, 5, 31, 0, 0, 177, 178, 3, 72, 36, 0, 178, 179, 5, 31, 0, 0, 179, 180,
		5, 19, 0, 0, 180, 183, 5, 31, 0, 0, 181, 184, 3, 54, 27, 0, 182, 184, 3,
		60, 30, 0, 183, 181, 1, 0, 0, 0, 183, 182, 1, 0, 0, 0, 184, 29, 1, 0, 0,
		0, 185, 186, 5, 4, 0, 0, 186, 187, 5, 31, 0, 0, 187, 188, 3, 72, 36, 0,
		188, 189, 5, 31, 0, 0, 189, 190, 5, 19, 0, 0, 190, 193, 5, 31, 0, 0, 191,
		194, 3, 54, 27, 0, 192, 194, 3, 60, 30, 0, 193, 191, 1, 0, 0, 0, 193, 192,
		1, 0, 0, 0, 194, 31, 1, 0, 0, 0, 195, 196, 5, 6, 0, 0, 196, 197, 5, 31,
		0, 0, 197, 198, 3, 54, 27, 0, 198, 33, 1, 0, 0, 0, 199, 200, 5, 1, 0, 0,
		200, 35, 1, 0, 0, 0, 201, 205, 3, 38, 19, 0, 202, 204, 3, 42, 21, 0, 203,
		202, 1, 0, 0, 0, 204, 207, 1, 0, 0, 0, 205, 203, 1, 0, 0, 0, 205, 206,
		1, 0, 0, 0, 206, 209, 1, 0, 0, 0, 207, 205, 1, 0, 0, 0, 208, 210, 3, 46,
		23, 0, 209, 208, 1, 0, 0, 0, 209, 210, 1, 0, 0, 0, 210, 211, 1, 0, 0, 0,
		211, 212, 3, 50, 25, 0, 212, 37, 1, 0, 0, 0, 213, 214, 3, 40, 20, 0, 214,
		215, 3, 24, 12, 0, 215, 39, 1, 0, 0, 0, 216, 217, 5, 22, 0, 0, 217, 218,
		5, 31, 0, 0, 218, 219, 3, 52, 26, 0, 219, 220, 5, 30, 0, 0, 220, 221, 5,
		31, 0, 0, 221, 222, 5, 13, 0, 0, 222, 223, 5, 32, 0, 0, 223, 41, 1, 0,
		0, 0, 224, 225, 3, 44, 22, 0, 225, 226, 3, 24, 12, 0, 226, 43, 1, 0, 0,
		0, 227, 228, 5, 23, 0, 0, 228, 229, 5, 30, 0, 0, 229, 230, 5, 31, 0, 0,
		230, 231, 3, 40, 20, 0, 231, 45, 1, 0, 0, 0, 232, 233, 3, 48, 24, 0, 233,
		234, 3, 24, 12, 0, 234, 47, 1, 0, 0, 0, 235, 236, 5, 23, 0, 0, 236, 237,
		5, 30, 0, 0, 237, 238, 5, 31, 0, 0, 238, 239, 5, 13, 0, 0, 239, 240, 5,
		32, 0, 0, 240, 49, 1, 0, 0, 0, 241, 242, 5, 14, 0, 0, 242, 243, 5, 30,
		0, 0, 243, 244, 5, 31, 0, 0, 244, 245, 5, 15, 0, 0, 245, 51, 1, 0, 0, 0,
		246, 247, 3, 54, 27, 0, 247, 248, 5, 31, 0, 0, 248, 249, 3, 70, 35, 0,
		249, 250, 5, 31, 0, 0, 250, 251, 3, 54, 27, 0, 251, 53, 1, 0, 0, 0, 252,
		258, 3, 56, 28, 0, 253, 254, 5, 31, 0, 0, 254, 255, 3, 68, 34, 0, 255,
		256, 5, 31, 0, 0, 256, 257, 3, 54, 27, 0, 257, 259, 1, 0, 0, 0, 258, 253,
		1, 0, 0, 0, 258, 259, 1, 0, 0, 0, 259, 55, 1, 0, 0, 0, 260, 263, 3, 58,
		29, 0, 261, 263, 3, 72, 36, 0, 262, 260, 1, 0, 0, 0, 262, 261, 1, 0, 0,
		0, 263, 57, 1, 0, 0, 0, 264, 265, 5, 24, 0, 0, 265, 266, 5, 31, 0, 0, 266,
		267, 3, 74, 37, 0, 267, 59, 1, 0, 0, 0, 268, 269, 5, 12, 0, 0, 269, 270,
		5, 31, 0, 0, 270, 278, 3, 72, 36, 0, 271, 273, 5, 30, 0, 0, 272, 271, 1,
		0, 0, 0, 272, 273, 1, 0, 0, 0, 273, 274, 1, 0, 0, 0, 274, 275, 5, 31, 0,
		0, 275, 276, 5, 21, 0, 0, 276, 277, 5, 31, 0, 0, 277, 279, 3, 62, 31, 0,
		278, 272, 1, 0, 0, 0, 278, 279, 1, 0, 0, 0, 279, 61, 1, 0, 0, 0, 280, 285,
		3, 64, 32, 0, 281, 282, 5, 31, 0, 0, 282, 283, 5, 20, 0, 0, 283, 284, 5,
		31, 0, 0, 284, 286, 3, 62, 31, 0, 285, 281, 1, 0, 0, 0, 285, 286, 1, 0,
		0, 0, 286, 63, 1, 0, 0, 0, 287, 288, 3, 54, 27, 0, 288, 65, 1, 0, 0, 0,
		289, 290, 5, 25, 0, 0, 290, 67, 1, 0, 0, 0, 291, 292, 7, 0, 0, 0, 292,
		69, 1, 0, 0, 0, 293, 294, 5, 19, 0, 0, 294, 71, 1, 0, 0, 0, 295, 296, 5,
		33, 0, 0, 296, 73, 1, 0, 0, 0, 297, 300, 3, 78, 39, 0, 298, 300, 3, 76,
		38, 0, 299, 297, 1, 0, 0, 0, 299, 298, 1, 0, 0, 0, 300, 75, 1, 0, 0, 0,
		301, 302, 5, 34, 0, 0, 302, 77, 1, 0, 0, 0, 303, 304, 5, 35, 0, 0, 304,
		79, 1, 0, 0, 0, 19, 85, 87, 93, 95, 115, 124, 146, 161, 168, 183, 193,
		205, 209, 258, 262, 272, 278, 285, 299,
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
	LectureParserCOMMENT                       = 1
	LectureParserOKAY_HEAR_ME_OUT              = 2
	LectureParserI_REST_MY_CASE                = 3
	LectureParserNOW_LETS_SAY                  = 4
	LectureParserLETS_SAY                      = 5
	LectureParserTHEN_WE_HAVE                  = 6
	LectureParserWE_CAN_USE_A_PROCESS_KNOWN_AS = 7
	LectureParserWHICH_NEEDS                   = 8
	LectureParserTO_PRODUCE_A                  = 9
	LectureParserIT_PROCEEDS_AS_FOLLOWS        = 10
	LectureParserFINALLY_WE_GET                = 11
	LectureParserTHE_RESULT_OF                 = 12
	LectureParserHERES_WHAT_WE_NEED_TO_DO      = 13
	LectureParserNOW_THAT_WEVE_DONE_THAT       = 14
	LectureParserWE_CAN_MOVE_ON                = 15
	LectureParserBY_THE_WAY                    = 16
	LectureParserA                             = 17
	LectureParserCALLED                        = 18
	LectureParserIS                            = 19
	LectureParserAND                           = 20
	LectureParserUSING                         = 21
	LectureParserIF                            = 22
	LectureParserOTHERWISE                     = 23
	LectureParserLITERALLY                     = 24
	LectureParserNUMBER                        = 25
	LectureParserPLUS                          = 26
	LectureParserMINUS                         = 27
	LectureParserQUOTE                         = 28
	LectureParserUNQUOTE                       = 29
	LectureParserCOMMA                         = 30
	LectureParserSPACE                         = 31
	LectureParserTERMINATOR                    = 32
	LectureParserIDENTIFIER_STRING             = 33
	LectureParserSTRING                        = 34
	LectureParserINTEGER                       = 35
	LectureParserWS                            = 36
)

// LectureParser rules.
const (
	LectureParserRULE_lecture                    = 0
	LectureParserRULE_program                    = 1
	LectureParserRULE_floatingComment            = 2
	LectureParserRULE_mainFunction               = 3
	LectureParserRULE_mainStartStatement         = 4
	LectureParserRULE_mainEndStatement           = 5
	LectureParserRULE_function                   = 6
	LectureParserRULE_functionSignature          = 7
	LectureParserRULE_parametersDeclaration      = 8
	LectureParserRULE_parameterDeclarationClause = 9
	LectureParserRULE_parameterDeclaration       = 10
	LectureParserRULE_statement                  = 11
	LectureParserRULE_statementBlock             = 12
	LectureParserRULE_returnStatement            = 13
	LectureParserRULE_assignmentStatement        = 14
	LectureParserRULE_reassignmentStatement      = 15
	LectureParserRULE_printStatement             = 16
	LectureParserRULE_commentStatement           = 17
	LectureParserRULE_ifChainStatement           = 18
	LectureParserRULE_ifStatement                = 19
	LectureParserRULE_ifSignature                = 20
	LectureParserRULE_elseIfStatement            = 21
	LectureParserRULE_elseIfSignature            = 22
	LectureParserRULE_elseStatement              = 23
	LectureParserRULE_elseSignature              = 24
	LectureParserRULE_ifClosingStatement         = 25
	LectureParserRULE_conditionClause            = 26
	LectureParserRULE_valueClause                = 27
	LectureParserRULE_value                      = 28
	LectureParserRULE_literalClause              = 29
	LectureParserRULE_functionCall               = 30
	LectureParserRULE_parametersClause           = 31
	LectureParserRULE_parameter                  = 32
	LectureParserRULE_type                       = 33
	LectureParserRULE_operator                   = 34
	LectureParserRULE_comparator                 = 35
	LectureParserRULE_identifier                 = 36
	LectureParserRULE_literal                    = 37
	LectureParserRULE_string                     = 38
	LectureParserRULE_number                     = 39
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
		p.SetState(80)
		p.Program()
	}
	{
		p.SetState(81)
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
	MainFunction() IMainFunctionContext
	AllFloatingComment() []IFloatingCommentContext
	FloatingComment(i int) IFloatingCommentContext
	AllFunction() []IFunctionContext
	Function(i int) IFunctionContext

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

func (s *ProgramContext) MainFunction() IMainFunctionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMainFunctionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMainFunctionContext)
}

func (s *ProgramContext) AllFloatingComment() []IFloatingCommentContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFloatingCommentContext); ok {
			len++
		}
	}

	tst := make([]IFloatingCommentContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFloatingCommentContext); ok {
			tst[i] = t.(IFloatingCommentContext)
			i++
		}
	}

	return tst
}

func (s *ProgramContext) FloatingComment(i int) IFloatingCommentContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFloatingCommentContext); ok {
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

	return t.(IFloatingCommentContext)
}

func (s *ProgramContext) AllFunction() []IFunctionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFunctionContext); ok {
			len++
		}
	}

	tst := make([]IFunctionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFunctionContext); ok {
			tst[i] = t.(IFunctionContext)
			i++
		}
	}

	return tst
}

func (s *ProgramContext) Function(i int) IFunctionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionContext); ok {
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

	return t.(IFunctionContext)
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
	p.SetState(87)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == LectureParserCOMMENT || _la == LectureParserWE_CAN_USE_A_PROCESS_KNOWN_AS {
		p.SetState(85)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case LectureParserCOMMENT:
			{
				p.SetState(83)
				p.FloatingComment()
			}

		case LectureParserWE_CAN_USE_A_PROCESS_KNOWN_AS:
			{
				p.SetState(84)
				p.Function()
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(89)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(90)
		p.MainFunction()
	}
	p.SetState(95)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == LectureParserCOMMENT || _la == LectureParserWE_CAN_USE_A_PROCESS_KNOWN_AS {
		p.SetState(93)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case LectureParserCOMMENT:
			{
				p.SetState(91)
				p.FloatingComment()
			}

		case LectureParserWE_CAN_USE_A_PROCESS_KNOWN_AS:
			{
				p.SetState(92)
				p.Function()
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(97)
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

// IFloatingCommentContext is an interface to support dynamic dispatch.
type IFloatingCommentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	CommentStatement() ICommentStatementContext
	TERMINATOR() antlr.TerminalNode

	// IsFloatingCommentContext differentiates from other interfaces.
	IsFloatingCommentContext()
}

type FloatingCommentContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFloatingCommentContext() *FloatingCommentContext {
	var p = new(FloatingCommentContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LectureParserRULE_floatingComment
	return p
}

func InitEmptyFloatingCommentContext(p *FloatingCommentContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LectureParserRULE_floatingComment
}

func (*FloatingCommentContext) IsFloatingCommentContext() {}

func NewFloatingCommentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FloatingCommentContext {
	var p = new(FloatingCommentContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LectureParserRULE_floatingComment

	return p
}

func (s *FloatingCommentContext) GetParser() antlr.Parser { return s.parser }

func (s *FloatingCommentContext) CommentStatement() ICommentStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICommentStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICommentStatementContext)
}

func (s *FloatingCommentContext) TERMINATOR() antlr.TerminalNode {
	return s.GetToken(LectureParserTERMINATOR, 0)
}

func (s *FloatingCommentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FloatingCommentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FloatingCommentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LectureListener); ok {
		listenerT.EnterFloatingComment(s)
	}
}

func (s *FloatingCommentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LectureListener); ok {
		listenerT.ExitFloatingComment(s)
	}
}

func (p *LectureParser) FloatingComment() (localctx IFloatingCommentContext) {
	localctx = NewFloatingCommentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, LectureParserRULE_floatingComment)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(98)
		p.CommentStatement()
	}
	{
		p.SetState(99)
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

// IMainFunctionContext is an interface to support dynamic dispatch.
type IMainFunctionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	MainStartStatement() IMainStartStatementContext
	StatementBlock() IStatementBlockContext
	MainEndStatement() IMainEndStatementContext

	// IsMainFunctionContext differentiates from other interfaces.
	IsMainFunctionContext()
}

type MainFunctionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMainFunctionContext() *MainFunctionContext {
	var p = new(MainFunctionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LectureParserRULE_mainFunction
	return p
}

func InitEmptyMainFunctionContext(p *MainFunctionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LectureParserRULE_mainFunction
}

func (*MainFunctionContext) IsMainFunctionContext() {}

func NewMainFunctionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MainFunctionContext {
	var p = new(MainFunctionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LectureParserRULE_mainFunction

	return p
}

func (s *MainFunctionContext) GetParser() antlr.Parser { return s.parser }

func (s *MainFunctionContext) MainStartStatement() IMainStartStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMainStartStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMainStartStatementContext)
}

func (s *MainFunctionContext) StatementBlock() IStatementBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementBlockContext)
}

func (s *MainFunctionContext) MainEndStatement() IMainEndStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMainEndStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMainEndStatementContext)
}

func (s *MainFunctionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MainFunctionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MainFunctionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LectureListener); ok {
		listenerT.EnterMainFunction(s)
	}
}

func (s *MainFunctionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LectureListener); ok {
		listenerT.ExitMainFunction(s)
	}
}

func (p *LectureParser) MainFunction() (localctx IMainFunctionContext) {
	localctx = NewMainFunctionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, LectureParserRULE_mainFunction)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(101)
		p.MainStartStatement()
	}
	{
		p.SetState(102)
		p.StatementBlock()
	}
	{
		p.SetState(103)
		p.MainEndStatement()
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

// IMainStartStatementContext is an interface to support dynamic dispatch.
type IMainStartStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	OKAY_HEAR_ME_OUT() antlr.TerminalNode
	TERMINATOR() antlr.TerminalNode

	// IsMainStartStatementContext differentiates from other interfaces.
	IsMainStartStatementContext()
}

type MainStartStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMainStartStatementContext() *MainStartStatementContext {
	var p = new(MainStartStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LectureParserRULE_mainStartStatement
	return p
}

func InitEmptyMainStartStatementContext(p *MainStartStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LectureParserRULE_mainStartStatement
}

func (*MainStartStatementContext) IsMainStartStatementContext() {}

func NewMainStartStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MainStartStatementContext {
	var p = new(MainStartStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LectureParserRULE_mainStartStatement

	return p
}

func (s *MainStartStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *MainStartStatementContext) OKAY_HEAR_ME_OUT() antlr.TerminalNode {
	return s.GetToken(LectureParserOKAY_HEAR_ME_OUT, 0)
}

func (s *MainStartStatementContext) TERMINATOR() antlr.TerminalNode {
	return s.GetToken(LectureParserTERMINATOR, 0)
}

func (s *MainStartStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MainStartStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MainStartStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LectureListener); ok {
		listenerT.EnterMainStartStatement(s)
	}
}

func (s *MainStartStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LectureListener); ok {
		listenerT.ExitMainStartStatement(s)
	}
}

func (p *LectureParser) MainStartStatement() (localctx IMainStartStatementContext) {
	localctx = NewMainStartStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, LectureParserRULE_mainStartStatement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(105)
		p.Match(LectureParserOKAY_HEAR_ME_OUT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(106)
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

// IMainEndStatementContext is an interface to support dynamic dispatch.
type IMainEndStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	I_REST_MY_CASE() antlr.TerminalNode
	TERMINATOR() antlr.TerminalNode

	// IsMainEndStatementContext differentiates from other interfaces.
	IsMainEndStatementContext()
}

type MainEndStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMainEndStatementContext() *MainEndStatementContext {
	var p = new(MainEndStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LectureParserRULE_mainEndStatement
	return p
}

func InitEmptyMainEndStatementContext(p *MainEndStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LectureParserRULE_mainEndStatement
}

func (*MainEndStatementContext) IsMainEndStatementContext() {}

func NewMainEndStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MainEndStatementContext {
	var p = new(MainEndStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LectureParserRULE_mainEndStatement

	return p
}

func (s *MainEndStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *MainEndStatementContext) I_REST_MY_CASE() antlr.TerminalNode {
	return s.GetToken(LectureParserI_REST_MY_CASE, 0)
}

func (s *MainEndStatementContext) TERMINATOR() antlr.TerminalNode {
	return s.GetToken(LectureParserTERMINATOR, 0)
}

func (s *MainEndStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MainEndStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MainEndStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LectureListener); ok {
		listenerT.EnterMainEndStatement(s)
	}
}

func (s *MainEndStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LectureListener); ok {
		listenerT.ExitMainEndStatement(s)
	}
}

func (p *LectureParser) MainEndStatement() (localctx IMainEndStatementContext) {
	localctx = NewMainEndStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, LectureParserRULE_mainEndStatement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(108)
		p.Match(LectureParserI_REST_MY_CASE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(109)
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
	FunctionSignature() IFunctionSignatureContext
	ReturnStatement() IReturnStatementContext
	AllStatement() []IStatementContext
	Statement(i int) IStatementContext

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

func (s *FunctionContext) FunctionSignature() IFunctionSignatureContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionSignatureContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunctionSignatureContext)
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

func (s *FunctionContext) AllStatement() []IStatementContext {
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

func (s *FunctionContext) Statement(i int) IStatementContext {
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
		p.SetState(111)
		p.FunctionSignature()
	}
	p.SetState(115)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&4194418) != 0 {
		{
			p.SetState(112)
			p.Statement()
		}

		p.SetState(117)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(118)
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

// IFunctionSignatureContext is an interface to support dynamic dispatch.
type IFunctionSignatureContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	WE_CAN_USE_A_PROCESS_KNOWN_AS() antlr.TerminalNode
	AllSPACE() []antlr.TerminalNode
	SPACE(i int) antlr.TerminalNode
	Identifier() IIdentifierContext
	TO_PRODUCE_A() antlr.TerminalNode
	Type_() ITypeContext
	AllTERMINATOR() []antlr.TerminalNode
	TERMINATOR(i int) antlr.TerminalNode
	IT_PROCEEDS_AS_FOLLOWS() antlr.TerminalNode
	ParametersDeclaration() IParametersDeclarationContext

	// IsFunctionSignatureContext differentiates from other interfaces.
	IsFunctionSignatureContext()
}

type FunctionSignatureContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionSignatureContext() *FunctionSignatureContext {
	var p = new(FunctionSignatureContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LectureParserRULE_functionSignature
	return p
}

func InitEmptyFunctionSignatureContext(p *FunctionSignatureContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LectureParserRULE_functionSignature
}

func (*FunctionSignatureContext) IsFunctionSignatureContext() {}

func NewFunctionSignatureContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionSignatureContext {
	var p = new(FunctionSignatureContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LectureParserRULE_functionSignature

	return p
}

func (s *FunctionSignatureContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionSignatureContext) WE_CAN_USE_A_PROCESS_KNOWN_AS() antlr.TerminalNode {
	return s.GetToken(LectureParserWE_CAN_USE_A_PROCESS_KNOWN_AS, 0)
}

func (s *FunctionSignatureContext) AllSPACE() []antlr.TerminalNode {
	return s.GetTokens(LectureParserSPACE)
}

func (s *FunctionSignatureContext) SPACE(i int) antlr.TerminalNode {
	return s.GetToken(LectureParserSPACE, i)
}

func (s *FunctionSignatureContext) Identifier() IIdentifierContext {
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

func (s *FunctionSignatureContext) TO_PRODUCE_A() antlr.TerminalNode {
	return s.GetToken(LectureParserTO_PRODUCE_A, 0)
}

func (s *FunctionSignatureContext) Type_() ITypeContext {
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

func (s *FunctionSignatureContext) AllTERMINATOR() []antlr.TerminalNode {
	return s.GetTokens(LectureParserTERMINATOR)
}

func (s *FunctionSignatureContext) TERMINATOR(i int) antlr.TerminalNode {
	return s.GetToken(LectureParserTERMINATOR, i)
}

func (s *FunctionSignatureContext) IT_PROCEEDS_AS_FOLLOWS() antlr.TerminalNode {
	return s.GetToken(LectureParserIT_PROCEEDS_AS_FOLLOWS, 0)
}

func (s *FunctionSignatureContext) ParametersDeclaration() IParametersDeclarationContext {
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

func (s *FunctionSignatureContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionSignatureContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionSignatureContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LectureListener); ok {
		listenerT.EnterFunctionSignature(s)
	}
}

func (s *FunctionSignatureContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LectureListener); ok {
		listenerT.ExitFunctionSignature(s)
	}
}

func (p *LectureParser) FunctionSignature() (localctx IFunctionSignatureContext) {
	localctx = NewFunctionSignatureContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, LectureParserRULE_functionSignature)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(120)
		p.Match(LectureParserWE_CAN_USE_A_PROCESS_KNOWN_AS)
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
		p.Identifier()
	}
	p.SetState(124)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == LectureParserCOMMA {
		{
			p.SetState(123)
			p.ParametersDeclaration()
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
		p.Match(LectureParserTO_PRODUCE_A)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
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
		p.Type_()
	}
	{
		p.SetState(130)
		p.Match(LectureParserTERMINATOR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(131)
		p.Match(LectureParserIT_PROCEEDS_AS_FOLLOWS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(132)
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
	p.EnterRule(localctx, 16, LectureParserRULE_parametersDeclaration)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(134)
		p.Match(LectureParserCOMMA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
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
		p.Match(LectureParserWHICH_NEEDS)
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
		p.ParameterDeclarationClause()
	}
	{
		p.SetState(139)
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
	p.EnterRule(localctx, 18, LectureParserRULE_parameterDeclarationClause)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(141)
		p.ParameterDeclaration()
	}
	p.SetState(146)
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
			p.Match(LectureParserAND)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
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
	p.EnterRule(localctx, 20, LectureParserRULE_parameterDeclaration)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(148)
		p.Match(LectureParserA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(149)
		p.Match(LectureParserSPACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(150)
		p.Type_()
	}
	{
		p.SetState(151)
		p.Match(LectureParserSPACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(152)
		p.Match(LectureParserCALLED)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(153)
		p.Match(LectureParserSPACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(154)
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

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TERMINATOR() antlr.TerminalNode
	AssignmentStatement() IAssignmentStatementContext
	ReassignmentStatement() IReassignmentStatementContext
	PrintStatement() IPrintStatementContext
	IfChainStatement() IIfChainStatementContext
	CommentStatement() ICommentStatementContext

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

func (s *StatementContext) AssignmentStatement() IAssignmentStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignmentStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssignmentStatementContext)
}

func (s *StatementContext) ReassignmentStatement() IReassignmentStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReassignmentStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReassignmentStatementContext)
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

func (s *StatementContext) IfChainStatement() IIfChainStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIfChainStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIfChainStatementContext)
}

func (s *StatementContext) CommentStatement() ICommentStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICommentStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICommentStatementContext)
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
	p.EnterRule(localctx, 22, LectureParserRULE_statement)
	p.EnterOuterAlt(localctx, 1)
	p.SetState(161)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case LectureParserLETS_SAY:
		{
			p.SetState(156)
			p.AssignmentStatement()
		}

	case LectureParserNOW_LETS_SAY:
		{
			p.SetState(157)
			p.ReassignmentStatement()
		}

	case LectureParserTHEN_WE_HAVE:
		{
			p.SetState(158)
			p.PrintStatement()
		}

	case LectureParserIF:
		{
			p.SetState(159)
			p.IfChainStatement()
		}

	case LectureParserCOMMENT:
		{
			p.SetState(160)
			p.CommentStatement()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	{
		p.SetState(163)
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

// IStatementBlockContext is an interface to support dynamic dispatch.
type IStatementBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllStatement() []IStatementContext
	Statement(i int) IStatementContext

	// IsStatementBlockContext differentiates from other interfaces.
	IsStatementBlockContext()
}

type StatementBlockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementBlockContext() *StatementBlockContext {
	var p = new(StatementBlockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LectureParserRULE_statementBlock
	return p
}

func InitEmptyStatementBlockContext(p *StatementBlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LectureParserRULE_statementBlock
}

func (*StatementBlockContext) IsStatementBlockContext() {}

func NewStatementBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementBlockContext {
	var p = new(StatementBlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LectureParserRULE_statementBlock

	return p
}

func (s *StatementBlockContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementBlockContext) AllStatement() []IStatementContext {
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

func (s *StatementBlockContext) Statement(i int) IStatementContext {
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

func (s *StatementBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementBlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementBlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LectureListener); ok {
		listenerT.EnterStatementBlock(s)
	}
}

func (s *StatementBlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LectureListener); ok {
		listenerT.ExitStatementBlock(s)
	}
}

func (p *LectureParser) StatementBlock() (localctx IStatementBlockContext) {
	localctx = NewStatementBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, LectureParserRULE_statementBlock)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(166)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&4194418) != 0) {
		{
			p.SetState(165)
			p.Statement()
		}

		p.SetState(168)
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
	p.EnterRule(localctx, 26, LectureParserRULE_returnStatement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(170)
		p.Match(LectureParserFINALLY_WE_GET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(171)
		p.Match(LectureParserSPACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(172)
		p.ValueClause()
	}
	{
		p.SetState(173)
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

// IAssignmentStatementContext is an interface to support dynamic dispatch.
type IAssignmentStatementContext interface {
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
	FunctionCall() IFunctionCallContext

	// IsAssignmentStatementContext differentiates from other interfaces.
	IsAssignmentStatementContext()
}

type AssignmentStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignmentStatementContext() *AssignmentStatementContext {
	var p = new(AssignmentStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LectureParserRULE_assignmentStatement
	return p
}

func InitEmptyAssignmentStatementContext(p *AssignmentStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LectureParserRULE_assignmentStatement
}

func (*AssignmentStatementContext) IsAssignmentStatementContext() {}

func NewAssignmentStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignmentStatementContext {
	var p = new(AssignmentStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LectureParserRULE_assignmentStatement

	return p
}

func (s *AssignmentStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignmentStatementContext) LETS_SAY() antlr.TerminalNode {
	return s.GetToken(LectureParserLETS_SAY, 0)
}

func (s *AssignmentStatementContext) AllSPACE() []antlr.TerminalNode {
	return s.GetTokens(LectureParserSPACE)
}

func (s *AssignmentStatementContext) SPACE(i int) antlr.TerminalNode {
	return s.GetToken(LectureParserSPACE, i)
}

func (s *AssignmentStatementContext) Identifier() IIdentifierContext {
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

func (s *AssignmentStatementContext) IS() antlr.TerminalNode {
	return s.GetToken(LectureParserIS, 0)
}

func (s *AssignmentStatementContext) ValueClause() IValueClauseContext {
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

func (s *AssignmentStatementContext) FunctionCall() IFunctionCallContext {
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

func (s *AssignmentStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignmentStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssignmentStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LectureListener); ok {
		listenerT.EnterAssignmentStatement(s)
	}
}

func (s *AssignmentStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LectureListener); ok {
		listenerT.ExitAssignmentStatement(s)
	}
}

func (p *LectureParser) AssignmentStatement() (localctx IAssignmentStatementContext) {
	localctx = NewAssignmentStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, LectureParserRULE_assignmentStatement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(175)
		p.Match(LectureParserLETS_SAY)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(176)
		p.Match(LectureParserSPACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(177)
		p.Identifier()
	}
	{
		p.SetState(178)
		p.Match(LectureParserSPACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(179)
		p.Match(LectureParserIS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(180)
		p.Match(LectureParserSPACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(183)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case LectureParserLITERALLY, LectureParserIDENTIFIER_STRING:
		{
			p.SetState(181)
			p.ValueClause()
		}

	case LectureParserTHE_RESULT_OF:
		{
			p.SetState(182)
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

// IReassignmentStatementContext is an interface to support dynamic dispatch.
type IReassignmentStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NOW_LETS_SAY() antlr.TerminalNode
	AllSPACE() []antlr.TerminalNode
	SPACE(i int) antlr.TerminalNode
	Identifier() IIdentifierContext
	IS() antlr.TerminalNode
	ValueClause() IValueClauseContext
	FunctionCall() IFunctionCallContext

	// IsReassignmentStatementContext differentiates from other interfaces.
	IsReassignmentStatementContext()
}

type ReassignmentStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReassignmentStatementContext() *ReassignmentStatementContext {
	var p = new(ReassignmentStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LectureParserRULE_reassignmentStatement
	return p
}

func InitEmptyReassignmentStatementContext(p *ReassignmentStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LectureParserRULE_reassignmentStatement
}

func (*ReassignmentStatementContext) IsReassignmentStatementContext() {}

func NewReassignmentStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReassignmentStatementContext {
	var p = new(ReassignmentStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LectureParserRULE_reassignmentStatement

	return p
}

func (s *ReassignmentStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *ReassignmentStatementContext) NOW_LETS_SAY() antlr.TerminalNode {
	return s.GetToken(LectureParserNOW_LETS_SAY, 0)
}

func (s *ReassignmentStatementContext) AllSPACE() []antlr.TerminalNode {
	return s.GetTokens(LectureParserSPACE)
}

func (s *ReassignmentStatementContext) SPACE(i int) antlr.TerminalNode {
	return s.GetToken(LectureParserSPACE, i)
}

func (s *ReassignmentStatementContext) Identifier() IIdentifierContext {
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

func (s *ReassignmentStatementContext) IS() antlr.TerminalNode {
	return s.GetToken(LectureParserIS, 0)
}

func (s *ReassignmentStatementContext) ValueClause() IValueClauseContext {
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

func (s *ReassignmentStatementContext) FunctionCall() IFunctionCallContext {
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

func (s *ReassignmentStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReassignmentStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReassignmentStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LectureListener); ok {
		listenerT.EnterReassignmentStatement(s)
	}
}

func (s *ReassignmentStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LectureListener); ok {
		listenerT.ExitReassignmentStatement(s)
	}
}

func (p *LectureParser) ReassignmentStatement() (localctx IReassignmentStatementContext) {
	localctx = NewReassignmentStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, LectureParserRULE_reassignmentStatement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(185)
		p.Match(LectureParserNOW_LETS_SAY)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(186)
		p.Match(LectureParserSPACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(187)
		p.Identifier()
	}
	{
		p.SetState(188)
		p.Match(LectureParserSPACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(189)
		p.Match(LectureParserIS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(190)
		p.Match(LectureParserSPACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(193)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case LectureParserLITERALLY, LectureParserIDENTIFIER_STRING:
		{
			p.SetState(191)
			p.ValueClause()
		}

	case LectureParserTHE_RESULT_OF:
		{
			p.SetState(192)
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
	p.EnterRule(localctx, 32, LectureParserRULE_printStatement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(195)
		p.Match(LectureParserTHEN_WE_HAVE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(196)
		p.Match(LectureParserSPACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(197)
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

// ICommentStatementContext is an interface to support dynamic dispatch.
type ICommentStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	COMMENT() antlr.TerminalNode

	// IsCommentStatementContext differentiates from other interfaces.
	IsCommentStatementContext()
}

type CommentStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCommentStatementContext() *CommentStatementContext {
	var p = new(CommentStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LectureParserRULE_commentStatement
	return p
}

func InitEmptyCommentStatementContext(p *CommentStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LectureParserRULE_commentStatement
}

func (*CommentStatementContext) IsCommentStatementContext() {}

func NewCommentStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CommentStatementContext {
	var p = new(CommentStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LectureParserRULE_commentStatement

	return p
}

func (s *CommentStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *CommentStatementContext) COMMENT() antlr.TerminalNode {
	return s.GetToken(LectureParserCOMMENT, 0)
}

func (s *CommentStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CommentStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CommentStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LectureListener); ok {
		listenerT.EnterCommentStatement(s)
	}
}

func (s *CommentStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LectureListener); ok {
		listenerT.ExitCommentStatement(s)
	}
}

func (p *LectureParser) CommentStatement() (localctx ICommentStatementContext) {
	localctx = NewCommentStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, LectureParserRULE_commentStatement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(199)
		p.Match(LectureParserCOMMENT)
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

// IIfChainStatementContext is an interface to support dynamic dispatch.
type IIfChainStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IfStatement() IIfStatementContext
	IfClosingStatement() IIfClosingStatementContext
	AllElseIfStatement() []IElseIfStatementContext
	ElseIfStatement(i int) IElseIfStatementContext
	ElseStatement() IElseStatementContext

	// IsIfChainStatementContext differentiates from other interfaces.
	IsIfChainStatementContext()
}

type IfChainStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIfChainStatementContext() *IfChainStatementContext {
	var p = new(IfChainStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LectureParserRULE_ifChainStatement
	return p
}

func InitEmptyIfChainStatementContext(p *IfChainStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LectureParserRULE_ifChainStatement
}

func (*IfChainStatementContext) IsIfChainStatementContext() {}

func NewIfChainStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfChainStatementContext {
	var p = new(IfChainStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LectureParserRULE_ifChainStatement

	return p
}

func (s *IfChainStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *IfChainStatementContext) IfStatement() IIfStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIfStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIfStatementContext)
}

func (s *IfChainStatementContext) IfClosingStatement() IIfClosingStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIfClosingStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIfClosingStatementContext)
}

func (s *IfChainStatementContext) AllElseIfStatement() []IElseIfStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IElseIfStatementContext); ok {
			len++
		}
	}

	tst := make([]IElseIfStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IElseIfStatementContext); ok {
			tst[i] = t.(IElseIfStatementContext)
			i++
		}
	}

	return tst
}

func (s *IfChainStatementContext) ElseIfStatement(i int) IElseIfStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IElseIfStatementContext); ok {
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

	return t.(IElseIfStatementContext)
}

func (s *IfChainStatementContext) ElseStatement() IElseStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IElseStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IElseStatementContext)
}

func (s *IfChainStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfChainStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IfChainStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LectureListener); ok {
		listenerT.EnterIfChainStatement(s)
	}
}

func (s *IfChainStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LectureListener); ok {
		listenerT.ExitIfChainStatement(s)
	}
}

func (p *LectureParser) IfChainStatement() (localctx IIfChainStatementContext) {
	localctx = NewIfChainStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, LectureParserRULE_ifChainStatement)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(201)
		p.IfStatement()
	}
	p.SetState(205)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 11, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(202)
				p.ElseIfStatement()
			}

		}
		p.SetState(207)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 11, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(209)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == LectureParserOTHERWISE {
		{
			p.SetState(208)
			p.ElseStatement()
		}

	}
	{
		p.SetState(211)
		p.IfClosingStatement()
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

// IIfStatementContext is an interface to support dynamic dispatch.
type IIfStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IfSignature() IIfSignatureContext
	StatementBlock() IStatementBlockContext

	// IsIfStatementContext differentiates from other interfaces.
	IsIfStatementContext()
}

type IfStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIfStatementContext() *IfStatementContext {
	var p = new(IfStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LectureParserRULE_ifStatement
	return p
}

func InitEmptyIfStatementContext(p *IfStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LectureParserRULE_ifStatement
}

func (*IfStatementContext) IsIfStatementContext() {}

func NewIfStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfStatementContext {
	var p = new(IfStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LectureParserRULE_ifStatement

	return p
}

func (s *IfStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *IfStatementContext) IfSignature() IIfSignatureContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIfSignatureContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIfSignatureContext)
}

func (s *IfStatementContext) StatementBlock() IStatementBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementBlockContext)
}

func (s *IfStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IfStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LectureListener); ok {
		listenerT.EnterIfStatement(s)
	}
}

func (s *IfStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LectureListener); ok {
		listenerT.ExitIfStatement(s)
	}
}

func (p *LectureParser) IfStatement() (localctx IIfStatementContext) {
	localctx = NewIfStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, LectureParserRULE_ifStatement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(213)
		p.IfSignature()
	}
	{
		p.SetState(214)
		p.StatementBlock()
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

// IIfSignatureContext is an interface to support dynamic dispatch.
type IIfSignatureContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IF() antlr.TerminalNode
	AllSPACE() []antlr.TerminalNode
	SPACE(i int) antlr.TerminalNode
	ConditionClause() IConditionClauseContext
	COMMA() antlr.TerminalNode
	HERES_WHAT_WE_NEED_TO_DO() antlr.TerminalNode
	TERMINATOR() antlr.TerminalNode

	// IsIfSignatureContext differentiates from other interfaces.
	IsIfSignatureContext()
}

type IfSignatureContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIfSignatureContext() *IfSignatureContext {
	var p = new(IfSignatureContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LectureParserRULE_ifSignature
	return p
}

func InitEmptyIfSignatureContext(p *IfSignatureContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LectureParserRULE_ifSignature
}

func (*IfSignatureContext) IsIfSignatureContext() {}

func NewIfSignatureContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfSignatureContext {
	var p = new(IfSignatureContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LectureParserRULE_ifSignature

	return p
}

func (s *IfSignatureContext) GetParser() antlr.Parser { return s.parser }

func (s *IfSignatureContext) IF() antlr.TerminalNode {
	return s.GetToken(LectureParserIF, 0)
}

func (s *IfSignatureContext) AllSPACE() []antlr.TerminalNode {
	return s.GetTokens(LectureParserSPACE)
}

func (s *IfSignatureContext) SPACE(i int) antlr.TerminalNode {
	return s.GetToken(LectureParserSPACE, i)
}

func (s *IfSignatureContext) ConditionClause() IConditionClauseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConditionClauseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConditionClauseContext)
}

func (s *IfSignatureContext) COMMA() antlr.TerminalNode {
	return s.GetToken(LectureParserCOMMA, 0)
}

func (s *IfSignatureContext) HERES_WHAT_WE_NEED_TO_DO() antlr.TerminalNode {
	return s.GetToken(LectureParserHERES_WHAT_WE_NEED_TO_DO, 0)
}

func (s *IfSignatureContext) TERMINATOR() antlr.TerminalNode {
	return s.GetToken(LectureParserTERMINATOR, 0)
}

func (s *IfSignatureContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfSignatureContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IfSignatureContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LectureListener); ok {
		listenerT.EnterIfSignature(s)
	}
}

func (s *IfSignatureContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LectureListener); ok {
		listenerT.ExitIfSignature(s)
	}
}

func (p *LectureParser) IfSignature() (localctx IIfSignatureContext) {
	localctx = NewIfSignatureContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, LectureParserRULE_ifSignature)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(216)
		p.Match(LectureParserIF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(217)
		p.Match(LectureParserSPACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(218)
		p.ConditionClause()
	}
	{
		p.SetState(219)
		p.Match(LectureParserCOMMA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(220)
		p.Match(LectureParserSPACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(221)
		p.Match(LectureParserHERES_WHAT_WE_NEED_TO_DO)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(222)
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

// IElseIfStatementContext is an interface to support dynamic dispatch.
type IElseIfStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ElseIfSignature() IElseIfSignatureContext
	StatementBlock() IStatementBlockContext

	// IsElseIfStatementContext differentiates from other interfaces.
	IsElseIfStatementContext()
}

type ElseIfStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyElseIfStatementContext() *ElseIfStatementContext {
	var p = new(ElseIfStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LectureParserRULE_elseIfStatement
	return p
}

func InitEmptyElseIfStatementContext(p *ElseIfStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LectureParserRULE_elseIfStatement
}

func (*ElseIfStatementContext) IsElseIfStatementContext() {}

func NewElseIfStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ElseIfStatementContext {
	var p = new(ElseIfStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LectureParserRULE_elseIfStatement

	return p
}

func (s *ElseIfStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *ElseIfStatementContext) ElseIfSignature() IElseIfSignatureContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IElseIfSignatureContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IElseIfSignatureContext)
}

func (s *ElseIfStatementContext) StatementBlock() IStatementBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementBlockContext)
}

func (s *ElseIfStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ElseIfStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ElseIfStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LectureListener); ok {
		listenerT.EnterElseIfStatement(s)
	}
}

func (s *ElseIfStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LectureListener); ok {
		listenerT.ExitElseIfStatement(s)
	}
}

func (p *LectureParser) ElseIfStatement() (localctx IElseIfStatementContext) {
	localctx = NewElseIfStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, LectureParserRULE_elseIfStatement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(224)
		p.ElseIfSignature()
	}
	{
		p.SetState(225)
		p.StatementBlock()
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

// IElseIfSignatureContext is an interface to support dynamic dispatch.
type IElseIfSignatureContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	OTHERWISE() antlr.TerminalNode
	COMMA() antlr.TerminalNode
	SPACE() antlr.TerminalNode
	IfSignature() IIfSignatureContext

	// IsElseIfSignatureContext differentiates from other interfaces.
	IsElseIfSignatureContext()
}

type ElseIfSignatureContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyElseIfSignatureContext() *ElseIfSignatureContext {
	var p = new(ElseIfSignatureContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LectureParserRULE_elseIfSignature
	return p
}

func InitEmptyElseIfSignatureContext(p *ElseIfSignatureContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LectureParserRULE_elseIfSignature
}

func (*ElseIfSignatureContext) IsElseIfSignatureContext() {}

func NewElseIfSignatureContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ElseIfSignatureContext {
	var p = new(ElseIfSignatureContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LectureParserRULE_elseIfSignature

	return p
}

func (s *ElseIfSignatureContext) GetParser() antlr.Parser { return s.parser }

func (s *ElseIfSignatureContext) OTHERWISE() antlr.TerminalNode {
	return s.GetToken(LectureParserOTHERWISE, 0)
}

func (s *ElseIfSignatureContext) COMMA() antlr.TerminalNode {
	return s.GetToken(LectureParserCOMMA, 0)
}

func (s *ElseIfSignatureContext) SPACE() antlr.TerminalNode {
	return s.GetToken(LectureParserSPACE, 0)
}

func (s *ElseIfSignatureContext) IfSignature() IIfSignatureContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIfSignatureContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIfSignatureContext)
}

func (s *ElseIfSignatureContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ElseIfSignatureContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ElseIfSignatureContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LectureListener); ok {
		listenerT.EnterElseIfSignature(s)
	}
}

func (s *ElseIfSignatureContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LectureListener); ok {
		listenerT.ExitElseIfSignature(s)
	}
}

func (p *LectureParser) ElseIfSignature() (localctx IElseIfSignatureContext) {
	localctx = NewElseIfSignatureContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, LectureParserRULE_elseIfSignature)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(227)
		p.Match(LectureParserOTHERWISE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(228)
		p.Match(LectureParserCOMMA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(229)
		p.Match(LectureParserSPACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(230)
		p.IfSignature()
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

// IElseStatementContext is an interface to support dynamic dispatch.
type IElseStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ElseSignature() IElseSignatureContext
	StatementBlock() IStatementBlockContext

	// IsElseStatementContext differentiates from other interfaces.
	IsElseStatementContext()
}

type ElseStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyElseStatementContext() *ElseStatementContext {
	var p = new(ElseStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LectureParserRULE_elseStatement
	return p
}

func InitEmptyElseStatementContext(p *ElseStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LectureParserRULE_elseStatement
}

func (*ElseStatementContext) IsElseStatementContext() {}

func NewElseStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ElseStatementContext {
	var p = new(ElseStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LectureParserRULE_elseStatement

	return p
}

func (s *ElseStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *ElseStatementContext) ElseSignature() IElseSignatureContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IElseSignatureContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IElseSignatureContext)
}

func (s *ElseStatementContext) StatementBlock() IStatementBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementBlockContext)
}

func (s *ElseStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ElseStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ElseStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LectureListener); ok {
		listenerT.EnterElseStatement(s)
	}
}

func (s *ElseStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LectureListener); ok {
		listenerT.ExitElseStatement(s)
	}
}

func (p *LectureParser) ElseStatement() (localctx IElseStatementContext) {
	localctx = NewElseStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, LectureParserRULE_elseStatement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(232)
		p.ElseSignature()
	}
	{
		p.SetState(233)
		p.StatementBlock()
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

// IElseSignatureContext is an interface to support dynamic dispatch.
type IElseSignatureContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	OTHERWISE() antlr.TerminalNode
	COMMA() antlr.TerminalNode
	SPACE() antlr.TerminalNode
	HERES_WHAT_WE_NEED_TO_DO() antlr.TerminalNode
	TERMINATOR() antlr.TerminalNode

	// IsElseSignatureContext differentiates from other interfaces.
	IsElseSignatureContext()
}

type ElseSignatureContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyElseSignatureContext() *ElseSignatureContext {
	var p = new(ElseSignatureContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LectureParserRULE_elseSignature
	return p
}

func InitEmptyElseSignatureContext(p *ElseSignatureContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LectureParserRULE_elseSignature
}

func (*ElseSignatureContext) IsElseSignatureContext() {}

func NewElseSignatureContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ElseSignatureContext {
	var p = new(ElseSignatureContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LectureParserRULE_elseSignature

	return p
}

func (s *ElseSignatureContext) GetParser() antlr.Parser { return s.parser }

func (s *ElseSignatureContext) OTHERWISE() antlr.TerminalNode {
	return s.GetToken(LectureParserOTHERWISE, 0)
}

func (s *ElseSignatureContext) COMMA() antlr.TerminalNode {
	return s.GetToken(LectureParserCOMMA, 0)
}

func (s *ElseSignatureContext) SPACE() antlr.TerminalNode {
	return s.GetToken(LectureParserSPACE, 0)
}

func (s *ElseSignatureContext) HERES_WHAT_WE_NEED_TO_DO() antlr.TerminalNode {
	return s.GetToken(LectureParserHERES_WHAT_WE_NEED_TO_DO, 0)
}

func (s *ElseSignatureContext) TERMINATOR() antlr.TerminalNode {
	return s.GetToken(LectureParserTERMINATOR, 0)
}

func (s *ElseSignatureContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ElseSignatureContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ElseSignatureContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LectureListener); ok {
		listenerT.EnterElseSignature(s)
	}
}

func (s *ElseSignatureContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LectureListener); ok {
		listenerT.ExitElseSignature(s)
	}
}

func (p *LectureParser) ElseSignature() (localctx IElseSignatureContext) {
	localctx = NewElseSignatureContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, LectureParserRULE_elseSignature)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(235)
		p.Match(LectureParserOTHERWISE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(236)
		p.Match(LectureParserCOMMA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(237)
		p.Match(LectureParserSPACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(238)
		p.Match(LectureParserHERES_WHAT_WE_NEED_TO_DO)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(239)
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

// IIfClosingStatementContext is an interface to support dynamic dispatch.
type IIfClosingStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NOW_THAT_WEVE_DONE_THAT() antlr.TerminalNode
	COMMA() antlr.TerminalNode
	SPACE() antlr.TerminalNode
	WE_CAN_MOVE_ON() antlr.TerminalNode

	// IsIfClosingStatementContext differentiates from other interfaces.
	IsIfClosingStatementContext()
}

type IfClosingStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIfClosingStatementContext() *IfClosingStatementContext {
	var p = new(IfClosingStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LectureParserRULE_ifClosingStatement
	return p
}

func InitEmptyIfClosingStatementContext(p *IfClosingStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LectureParserRULE_ifClosingStatement
}

func (*IfClosingStatementContext) IsIfClosingStatementContext() {}

func NewIfClosingStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfClosingStatementContext {
	var p = new(IfClosingStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LectureParserRULE_ifClosingStatement

	return p
}

func (s *IfClosingStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *IfClosingStatementContext) NOW_THAT_WEVE_DONE_THAT() antlr.TerminalNode {
	return s.GetToken(LectureParserNOW_THAT_WEVE_DONE_THAT, 0)
}

func (s *IfClosingStatementContext) COMMA() antlr.TerminalNode {
	return s.GetToken(LectureParserCOMMA, 0)
}

func (s *IfClosingStatementContext) SPACE() antlr.TerminalNode {
	return s.GetToken(LectureParserSPACE, 0)
}

func (s *IfClosingStatementContext) WE_CAN_MOVE_ON() antlr.TerminalNode {
	return s.GetToken(LectureParserWE_CAN_MOVE_ON, 0)
}

func (s *IfClosingStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfClosingStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IfClosingStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LectureListener); ok {
		listenerT.EnterIfClosingStatement(s)
	}
}

func (s *IfClosingStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LectureListener); ok {
		listenerT.ExitIfClosingStatement(s)
	}
}

func (p *LectureParser) IfClosingStatement() (localctx IIfClosingStatementContext) {
	localctx = NewIfClosingStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, LectureParserRULE_ifClosingStatement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(241)
		p.Match(LectureParserNOW_THAT_WEVE_DONE_THAT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(242)
		p.Match(LectureParserCOMMA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(243)
		p.Match(LectureParserSPACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(244)
		p.Match(LectureParserWE_CAN_MOVE_ON)
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

// IConditionClauseContext is an interface to support dynamic dispatch.
type IConditionClauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllValueClause() []IValueClauseContext
	ValueClause(i int) IValueClauseContext
	AllSPACE() []antlr.TerminalNode
	SPACE(i int) antlr.TerminalNode
	Comparator() IComparatorContext

	// IsConditionClauseContext differentiates from other interfaces.
	IsConditionClauseContext()
}

type ConditionClauseContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConditionClauseContext() *ConditionClauseContext {
	var p = new(ConditionClauseContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LectureParserRULE_conditionClause
	return p
}

func InitEmptyConditionClauseContext(p *ConditionClauseContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LectureParserRULE_conditionClause
}

func (*ConditionClauseContext) IsConditionClauseContext() {}

func NewConditionClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConditionClauseContext {
	var p = new(ConditionClauseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LectureParserRULE_conditionClause

	return p
}

func (s *ConditionClauseContext) GetParser() antlr.Parser { return s.parser }

func (s *ConditionClauseContext) AllValueClause() []IValueClauseContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IValueClauseContext); ok {
			len++
		}
	}

	tst := make([]IValueClauseContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IValueClauseContext); ok {
			tst[i] = t.(IValueClauseContext)
			i++
		}
	}

	return tst
}

func (s *ConditionClauseContext) ValueClause(i int) IValueClauseContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueClauseContext); ok {
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

	return t.(IValueClauseContext)
}

func (s *ConditionClauseContext) AllSPACE() []antlr.TerminalNode {
	return s.GetTokens(LectureParserSPACE)
}

func (s *ConditionClauseContext) SPACE(i int) antlr.TerminalNode {
	return s.GetToken(LectureParserSPACE, i)
}

func (s *ConditionClauseContext) Comparator() IComparatorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IComparatorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IComparatorContext)
}

func (s *ConditionClauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConditionClauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConditionClauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LectureListener); ok {
		listenerT.EnterConditionClause(s)
	}
}

func (s *ConditionClauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LectureListener); ok {
		listenerT.ExitConditionClause(s)
	}
}

func (p *LectureParser) ConditionClause() (localctx IConditionClauseContext) {
	localctx = NewConditionClauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, LectureParserRULE_conditionClause)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(246)
		p.ValueClause()
	}
	{
		p.SetState(247)
		p.Match(LectureParserSPACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(248)
		p.Comparator()
	}
	{
		p.SetState(249)
		p.Match(LectureParserSPACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(250)
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
	p.EnterRule(localctx, 54, LectureParserRULE_valueClause)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(252)
		p.Value()
	}
	p.SetState(258)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 13, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(253)
			p.Match(LectureParserSPACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(254)
			p.Operator()
		}
		{
			p.SetState(255)
			p.Match(LectureParserSPACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(256)
			p.ValueClause()
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
	p.EnterRule(localctx, 56, LectureParserRULE_value)
	p.SetState(262)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case LectureParserLITERALLY:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(260)
			p.LiteralClause()
		}

	case LectureParserIDENTIFIER_STRING:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(261)
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
	p.EnterRule(localctx, 58, LectureParserRULE_literalClause)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(264)
		p.Match(LectureParserLITERALLY)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(265)
		p.Match(LectureParserSPACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(266)
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
	USING() antlr.TerminalNode
	ParametersClause() IParametersClauseContext
	COMMA() antlr.TerminalNode

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

func (s *FunctionCallContext) USING() antlr.TerminalNode {
	return s.GetToken(LectureParserUSING, 0)
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

func (s *FunctionCallContext) COMMA() antlr.TerminalNode {
	return s.GetToken(LectureParserCOMMA, 0)
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
	p.EnterRule(localctx, 60, LectureParserRULE_functionCall)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(268)
		p.Match(LectureParserTHE_RESULT_OF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(269)
		p.Match(LectureParserSPACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(270)
		p.Identifier()
	}
	p.SetState(278)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == LectureParserCOMMA || _la == LectureParserSPACE {
		p.SetState(272)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == LectureParserCOMMA {
			{
				p.SetState(271)
				p.Match(LectureParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(274)
			p.Match(LectureParserSPACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(275)
			p.Match(LectureParserUSING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(276)
			p.Match(LectureParserSPACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(277)
			p.ParametersClause()
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

// IParametersClauseContext is an interface to support dynamic dispatch.
type IParametersClauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Parameter() IParameterContext
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

func (s *ParametersClauseContext) Parameter() IParameterContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParameterContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParameterContext)
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
	p.EnterRule(localctx, 62, LectureParserRULE_parametersClause)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(280)
		p.Parameter()
	}
	p.SetState(285)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == LectureParserSPACE {
		{
			p.SetState(281)
			p.Match(LectureParserSPACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(282)
			p.Match(LectureParserAND)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(283)
			p.Match(LectureParserSPACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(284)
			p.ParametersClause()
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

// IParameterContext is an interface to support dynamic dispatch.
type IParameterContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ValueClause() IValueClauseContext

	// IsParameterContext differentiates from other interfaces.
	IsParameterContext()
}

type ParameterContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParameterContext() *ParameterContext {
	var p = new(ParameterContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LectureParserRULE_parameter
	return p
}

func InitEmptyParameterContext(p *ParameterContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LectureParserRULE_parameter
}

func (*ParameterContext) IsParameterContext() {}

func NewParameterContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParameterContext {
	var p = new(ParameterContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LectureParserRULE_parameter

	return p
}

func (s *ParameterContext) GetParser() antlr.Parser { return s.parser }

func (s *ParameterContext) ValueClause() IValueClauseContext {
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

func (s *ParameterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParameterContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParameterContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LectureListener); ok {
		listenerT.EnterParameter(s)
	}
}

func (s *ParameterContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LectureListener); ok {
		listenerT.ExitParameter(s)
	}
}

func (p *LectureParser) Parameter() (localctx IParameterContext) {
	localctx = NewParameterContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, LectureParserRULE_parameter)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(287)
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
	p.EnterRule(localctx, 66, LectureParserRULE_type)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(289)
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
	MINUS() antlr.TerminalNode

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

func (s *OperatorContext) MINUS() antlr.TerminalNode {
	return s.GetToken(LectureParserMINUS, 0)
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
	p.EnterRule(localctx, 68, LectureParserRULE_operator)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(291)
		_la = p.GetTokenStream().LA(1)

		if !(_la == LectureParserPLUS || _la == LectureParserMINUS) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
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

// IComparatorContext is an interface to support dynamic dispatch.
type IComparatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IS() antlr.TerminalNode

	// IsComparatorContext differentiates from other interfaces.
	IsComparatorContext()
}

type ComparatorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyComparatorContext() *ComparatorContext {
	var p = new(ComparatorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LectureParserRULE_comparator
	return p
}

func InitEmptyComparatorContext(p *ComparatorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LectureParserRULE_comparator
}

func (*ComparatorContext) IsComparatorContext() {}

func NewComparatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ComparatorContext {
	var p = new(ComparatorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LectureParserRULE_comparator

	return p
}

func (s *ComparatorContext) GetParser() antlr.Parser { return s.parser }

func (s *ComparatorContext) IS() antlr.TerminalNode {
	return s.GetToken(LectureParserIS, 0)
}

func (s *ComparatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ComparatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ComparatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LectureListener); ok {
		listenerT.EnterComparator(s)
	}
}

func (s *ComparatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LectureListener); ok {
		listenerT.ExitComparator(s)
	}
}

func (p *LectureParser) Comparator() (localctx IComparatorContext) {
	localctx = NewComparatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, LectureParserRULE_comparator)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(293)
		p.Match(LectureParserIS)
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
	IDENTIFIER_STRING() antlr.TerminalNode

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

func (s *IdentifierContext) IDENTIFIER_STRING() antlr.TerminalNode {
	return s.GetToken(LectureParserIDENTIFIER_STRING, 0)
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
	p.EnterRule(localctx, 72, LectureParserRULE_identifier)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(295)
		p.Match(LectureParserIDENTIFIER_STRING)
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

// ILiteralContext is an interface to support dynamic dispatch.
type ILiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Number() INumberContext
	String_() IStringContext

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

func (s *LiteralContext) String_() IStringContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStringContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStringContext)
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
	p.EnterRule(localctx, 74, LectureParserRULE_literal)
	p.SetState(299)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case LectureParserINTEGER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(297)
			p.Number()
		}

	case LectureParserSTRING:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(298)
			p.String_()
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

// IStringContext is an interface to support dynamic dispatch.
type IStringContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	STRING() antlr.TerminalNode

	// IsStringContext differentiates from other interfaces.
	IsStringContext()
}

type StringContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStringContext() *StringContext {
	var p = new(StringContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LectureParserRULE_string
	return p
}

func InitEmptyStringContext(p *StringContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LectureParserRULE_string
}

func (*StringContext) IsStringContext() {}

func NewStringContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StringContext {
	var p = new(StringContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LectureParserRULE_string

	return p
}

func (s *StringContext) GetParser() antlr.Parser { return s.parser }

func (s *StringContext) STRING() antlr.TerminalNode {
	return s.GetToken(LectureParserSTRING, 0)
}

func (s *StringContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StringContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LectureListener); ok {
		listenerT.EnterString(s)
	}
}

func (s *StringContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LectureListener); ok {
		listenerT.ExitString(s)
	}
}

func (p *LectureParser) String_() (localctx IStringContext) {
	localctx = NewStringContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 76, LectureParserRULE_string)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(301)
		p.Match(LectureParserSTRING)
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
	p.EnterRule(localctx, 78, LectureParserRULE_number)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(303)
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
