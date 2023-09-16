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

// EnterAtomicStatement is called when production atomicStatement is entered.
func (s *BaseLectureListener) EnterAtomicStatement(ctx *AtomicStatementContext) {}

// ExitAtomicStatement is called when production atomicStatement is exited.
func (s *BaseLectureListener) ExitAtomicStatement(ctx *AtomicStatementContext) {}

// EnterFunction is called when production function is entered.
func (s *BaseLectureListener) EnterFunction(ctx *FunctionContext) {}

// ExitFunction is called when production function is exited.
func (s *BaseLectureListener) ExitFunction(ctx *FunctionContext) {}

// EnterFunctionStatement is called when production functionStatement is entered.
func (s *BaseLectureListener) EnterFunctionStatement(ctx *FunctionStatementContext) {}

// ExitFunctionStatement is called when production functionStatement is exited.
func (s *BaseLectureListener) ExitFunctionStatement(ctx *FunctionStatementContext) {}

// EnterReturnStatement is called when production returnStatement is entered.
func (s *BaseLectureListener) EnterReturnStatement(ctx *ReturnStatementContext) {}

// ExitReturnStatement is called when production returnStatement is exited.
func (s *BaseLectureListener) ExitReturnStatement(ctx *ReturnStatementContext) {}

// EnterDeclarationStatement is called when production declarationStatement is entered.
func (s *BaseLectureListener) EnterDeclarationStatement(ctx *DeclarationStatementContext) {}

// ExitDeclarationStatement is called when production declarationStatement is exited.
func (s *BaseLectureListener) ExitDeclarationStatement(ctx *DeclarationStatementContext) {}

// EnterPrintStatement is called when production printStatement is entered.
func (s *BaseLectureListener) EnterPrintStatement(ctx *PrintStatementContext) {}

// ExitPrintStatement is called when production printStatement is exited.
func (s *BaseLectureListener) ExitPrintStatement(ctx *PrintStatementContext) {}

// EnterParametersDeclaration is called when production parametersDeclaration is entered.
func (s *BaseLectureListener) EnterParametersDeclaration(ctx *ParametersDeclarationContext) {}

// ExitParametersDeclaration is called when production parametersDeclaration is exited.
func (s *BaseLectureListener) ExitParametersDeclaration(ctx *ParametersDeclarationContext) {}

// EnterParameterDeclarationClause is called when production parameterDeclarationClause is entered.
func (s *BaseLectureListener) EnterParameterDeclarationClause(ctx *ParameterDeclarationClauseContext) {
}

// ExitParameterDeclarationClause is called when production parameterDeclarationClause is exited.
func (s *BaseLectureListener) ExitParameterDeclarationClause(ctx *ParameterDeclarationClauseContext) {
}

// EnterFunctionCall is called when production functionCall is entered.
func (s *BaseLectureListener) EnterFunctionCall(ctx *FunctionCallContext) {}

// ExitFunctionCall is called when production functionCall is exited.
func (s *BaseLectureListener) ExitFunctionCall(ctx *FunctionCallContext) {}

// EnterParametersClause is called when production parametersClause is entered.
func (s *BaseLectureListener) EnterParametersClause(ctx *ParametersClauseContext) {}

// ExitParametersClause is called when production parametersClause is exited.
func (s *BaseLectureListener) ExitParametersClause(ctx *ParametersClauseContext) {}

// EnterValueClause is called when production valueClause is entered.
func (s *BaseLectureListener) EnterValueClause(ctx *ValueClauseContext) {}

// ExitValueClause is called when production valueClause is exited.
func (s *BaseLectureListener) ExitValueClause(ctx *ValueClauseContext) {}

// EnterLiteralClause is called when production literalClause is entered.
func (s *BaseLectureListener) EnterLiteralClause(ctx *LiteralClauseContext) {}

// ExitLiteralClause is called when production literalClause is exited.
func (s *BaseLectureListener) ExitLiteralClause(ctx *LiteralClauseContext) {}

// EnterParameterDeclaration is called when production parameterDeclaration is entered.
func (s *BaseLectureListener) EnterParameterDeclaration(ctx *ParameterDeclarationContext) {}

// ExitParameterDeclaration is called when production parameterDeclaration is exited.
func (s *BaseLectureListener) ExitParameterDeclaration(ctx *ParameterDeclarationContext) {}

// EnterValue is called when production value is entered.
func (s *BaseLectureListener) EnterValue(ctx *ValueContext) {}

// ExitValue is called when production value is exited.
func (s *BaseLectureListener) ExitValue(ctx *ValueContext) {}

// EnterLiteral is called when production literal is entered.
func (s *BaseLectureListener) EnterLiteral(ctx *LiteralContext) {}

// ExitLiteral is called when production literal is exited.
func (s *BaseLectureListener) ExitLiteral(ctx *LiteralContext) {}

// EnterType is called when production type is entered.
func (s *BaseLectureListener) EnterType(ctx *TypeContext) {}

// ExitType is called when production type is exited.
func (s *BaseLectureListener) ExitType(ctx *TypeContext) {}

// EnterOperator is called when production operator is entered.
func (s *BaseLectureListener) EnterOperator(ctx *OperatorContext) {}

// ExitOperator is called when production operator is exited.
func (s *BaseLectureListener) ExitOperator(ctx *OperatorContext) {}

// EnterIdentifier is called when production identifier is entered.
func (s *BaseLectureListener) EnterIdentifier(ctx *IdentifierContext) {}

// ExitIdentifier is called when production identifier is exited.
func (s *BaseLectureListener) ExitIdentifier(ctx *IdentifierContext) {}

// EnterNumber is called when production number is entered.
func (s *BaseLectureListener) EnterNumber(ctx *NumberContext) {}

// ExitNumber is called when production number is exited.
func (s *BaseLectureListener) ExitNumber(ctx *NumberContext) {}
