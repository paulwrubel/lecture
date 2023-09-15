// Code generated from Lecture.g4 by ANTLR 4.13.1. DO NOT EDIT.

package lecture // Lecture
import "github.com/antlr4-go/antlr/v4"

// BaseLectureListener is a complete listener for a parse tree produced by LectureParser.
type BaseLectureListener struct{}

var _ LectureListener = &BaseLectureListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseLectureListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseLectureListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseLectureListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseLectureListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterLecture is called when production lecture is entered.
func (s *BaseLectureListener) EnterLecture(ctx *LectureContext) {}

// ExitLecture is called when production lecture is exited.
func (s *BaseLectureListener) ExitLecture(ctx *LectureContext) {}

// EnterProgram is called when production program is entered.
func (s *BaseLectureListener) EnterProgram(ctx *ProgramContext) {}

// ExitProgram is called when production program is exited.
func (s *BaseLectureListener) ExitProgram(ctx *ProgramContext) {}

// EnterStartClause is called when production startClause is entered.
func (s *BaseLectureListener) EnterStartClause(ctx *StartClauseContext) {}

// ExitStartClause is called when production startClause is exited.
func (s *BaseLectureListener) ExitStartClause(ctx *StartClauseContext) {}

// EnterEndClause is called when production endClause is entered.
func (s *BaseLectureListener) EnterEndClause(ctx *EndClauseContext) {}

// ExitEndClause is called when production endClause is exited.
func (s *BaseLectureListener) ExitEndClause(ctx *EndClauseContext) {}

// EnterStatement is called when production statement is entered.
func (s *BaseLectureListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BaseLectureListener) ExitStatement(ctx *StatementContext) {}

// EnterDeclarationStatement is called when production declarationStatement is entered.
func (s *BaseLectureListener) EnterDeclarationStatement(ctx *DeclarationStatementContext) {}

// ExitDeclarationStatement is called when production declarationStatement is exited.
func (s *BaseLectureListener) ExitDeclarationStatement(ctx *DeclarationStatementContext) {}

// EnterPrintStatement is called when production printStatement is entered.
func (s *BaseLectureListener) EnterPrintStatement(ctx *PrintStatementContext) {}

// ExitPrintStatement is called when production printStatement is exited.
func (s *BaseLectureListener) ExitPrintStatement(ctx *PrintStatementContext) {}

// EnterVariable is called when production variable is entered.
func (s *BaseLectureListener) EnterVariable(ctx *VariableContext) {}

// ExitVariable is called when production variable is exited.
func (s *BaseLectureListener) ExitVariable(ctx *VariableContext) {}

// EnterValue is called when production value is entered.
func (s *BaseLectureListener) EnterValue(ctx *ValueContext) {}

// ExitValue is called when production value is exited.
func (s *BaseLectureListener) ExitValue(ctx *ValueContext) {}

// EnterNumber is called when production number is entered.
func (s *BaseLectureListener) EnterNumber(ctx *NumberContext) {}

// ExitNumber is called when production number is exited.
func (s *BaseLectureListener) ExitNumber(ctx *NumberContext) {}
