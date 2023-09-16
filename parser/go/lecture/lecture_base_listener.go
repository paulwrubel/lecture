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

// EnterMainFunction is called when production mainFunction is entered.
func (s *BaseLectureListener) EnterMainFunction(ctx *MainFunctionContext) {}

// ExitMainFunction is called when production mainFunction is exited.
func (s *BaseLectureListener) ExitMainFunction(ctx *MainFunctionContext) {}

// EnterMainStartStatement is called when production mainStartStatement is entered.
func (s *BaseLectureListener) EnterMainStartStatement(ctx *MainStartStatementContext) {}

// ExitMainStartStatement is called when production mainStartStatement is exited.
func (s *BaseLectureListener) ExitMainStartStatement(ctx *MainStartStatementContext) {}

// EnterMainEndStatement is called when production mainEndStatement is entered.
func (s *BaseLectureListener) EnterMainEndStatement(ctx *MainEndStatementContext) {}

// ExitMainEndStatement is called when production mainEndStatement is exited.
func (s *BaseLectureListener) ExitMainEndStatement(ctx *MainEndStatementContext) {}

// EnterFunction is called when production function is entered.
func (s *BaseLectureListener) EnterFunction(ctx *FunctionContext) {}

// ExitFunction is called when production function is exited.
func (s *BaseLectureListener) ExitFunction(ctx *FunctionContext) {}

// EnterFunctionSignature is called when production functionSignature is entered.
func (s *BaseLectureListener) EnterFunctionSignature(ctx *FunctionSignatureContext) {}

// ExitFunctionSignature is called when production functionSignature is exited.
func (s *BaseLectureListener) ExitFunctionSignature(ctx *FunctionSignatureContext) {}

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

// EnterParameterDeclaration is called when production parameterDeclaration is entered.
func (s *BaseLectureListener) EnterParameterDeclaration(ctx *ParameterDeclarationContext) {}

// ExitParameterDeclaration is called when production parameterDeclaration is exited.
func (s *BaseLectureListener) ExitParameterDeclaration(ctx *ParameterDeclarationContext) {}

// EnterStatement is called when production statement is entered.
func (s *BaseLectureListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BaseLectureListener) ExitStatement(ctx *StatementContext) {}

// EnterStatementBlock is called when production statementBlock is entered.
func (s *BaseLectureListener) EnterStatementBlock(ctx *StatementBlockContext) {}

// ExitStatementBlock is called when production statementBlock is exited.
func (s *BaseLectureListener) ExitStatementBlock(ctx *StatementBlockContext) {}

// EnterReturnStatement is called when production returnStatement is entered.
func (s *BaseLectureListener) EnterReturnStatement(ctx *ReturnStatementContext) {}

// ExitReturnStatement is called when production returnStatement is exited.
func (s *BaseLectureListener) ExitReturnStatement(ctx *ReturnStatementContext) {}

// EnterAssignmentStatement is called when production assignmentStatement is entered.
func (s *BaseLectureListener) EnterAssignmentStatement(ctx *AssignmentStatementContext) {}

// ExitAssignmentStatement is called when production assignmentStatement is exited.
func (s *BaseLectureListener) ExitAssignmentStatement(ctx *AssignmentStatementContext) {}

// EnterReassignmentStatement is called when production reassignmentStatement is entered.
func (s *BaseLectureListener) EnterReassignmentStatement(ctx *ReassignmentStatementContext) {}

// ExitReassignmentStatement is called when production reassignmentStatement is exited.
func (s *BaseLectureListener) ExitReassignmentStatement(ctx *ReassignmentStatementContext) {}

// EnterPrintStatement is called when production printStatement is entered.
func (s *BaseLectureListener) EnterPrintStatement(ctx *PrintStatementContext) {}

// ExitPrintStatement is called when production printStatement is exited.
func (s *BaseLectureListener) ExitPrintStatement(ctx *PrintStatementContext) {}

// EnterIfChainStatement is called when production ifChainStatement is entered.
func (s *BaseLectureListener) EnterIfChainStatement(ctx *IfChainStatementContext) {}

// ExitIfChainStatement is called when production ifChainStatement is exited.
func (s *BaseLectureListener) ExitIfChainStatement(ctx *IfChainStatementContext) {}

// EnterIfStatement is called when production ifStatement is entered.
func (s *BaseLectureListener) EnterIfStatement(ctx *IfStatementContext) {}

// ExitIfStatement is called when production ifStatement is exited.
func (s *BaseLectureListener) ExitIfStatement(ctx *IfStatementContext) {}

// EnterIfSignature is called when production ifSignature is entered.
func (s *BaseLectureListener) EnterIfSignature(ctx *IfSignatureContext) {}

// ExitIfSignature is called when production ifSignature is exited.
func (s *BaseLectureListener) ExitIfSignature(ctx *IfSignatureContext) {}

// EnterElseIfStatement is called when production elseIfStatement is entered.
func (s *BaseLectureListener) EnterElseIfStatement(ctx *ElseIfStatementContext) {}

// ExitElseIfStatement is called when production elseIfStatement is exited.
func (s *BaseLectureListener) ExitElseIfStatement(ctx *ElseIfStatementContext) {}

// EnterElseIfSignature is called when production elseIfSignature is entered.
func (s *BaseLectureListener) EnterElseIfSignature(ctx *ElseIfSignatureContext) {}

// ExitElseIfSignature is called when production elseIfSignature is exited.
func (s *BaseLectureListener) ExitElseIfSignature(ctx *ElseIfSignatureContext) {}

// EnterElseStatement is called when production elseStatement is entered.
func (s *BaseLectureListener) EnterElseStatement(ctx *ElseStatementContext) {}

// ExitElseStatement is called when production elseStatement is exited.
func (s *BaseLectureListener) ExitElseStatement(ctx *ElseStatementContext) {}

// EnterElseSignature is called when production elseSignature is entered.
func (s *BaseLectureListener) EnterElseSignature(ctx *ElseSignatureContext) {}

// ExitElseSignature is called when production elseSignature is exited.
func (s *BaseLectureListener) ExitElseSignature(ctx *ElseSignatureContext) {}

// EnterIfClosingStatement is called when production ifClosingStatement is entered.
func (s *BaseLectureListener) EnterIfClosingStatement(ctx *IfClosingStatementContext) {}

// ExitIfClosingStatement is called when production ifClosingStatement is exited.
func (s *BaseLectureListener) ExitIfClosingStatement(ctx *IfClosingStatementContext) {}

// EnterConditionClause is called when production conditionClause is entered.
func (s *BaseLectureListener) EnterConditionClause(ctx *ConditionClauseContext) {}

// ExitConditionClause is called when production conditionClause is exited.
func (s *BaseLectureListener) ExitConditionClause(ctx *ConditionClauseContext) {}

// EnterValueClause is called when production valueClause is entered.
func (s *BaseLectureListener) EnterValueClause(ctx *ValueClauseContext) {}

// ExitValueClause is called when production valueClause is exited.
func (s *BaseLectureListener) ExitValueClause(ctx *ValueClauseContext) {}

// EnterValue is called when production value is entered.
func (s *BaseLectureListener) EnterValue(ctx *ValueContext) {}

// ExitValue is called when production value is exited.
func (s *BaseLectureListener) ExitValue(ctx *ValueContext) {}

// EnterLiteralClause is called when production literalClause is entered.
func (s *BaseLectureListener) EnterLiteralClause(ctx *LiteralClauseContext) {}

// ExitLiteralClause is called when production literalClause is exited.
func (s *BaseLectureListener) ExitLiteralClause(ctx *LiteralClauseContext) {}

// EnterFunctionCall is called when production functionCall is entered.
func (s *BaseLectureListener) EnterFunctionCall(ctx *FunctionCallContext) {}

// ExitFunctionCall is called when production functionCall is exited.
func (s *BaseLectureListener) ExitFunctionCall(ctx *FunctionCallContext) {}

// EnterParametersClause is called when production parametersClause is entered.
func (s *BaseLectureListener) EnterParametersClause(ctx *ParametersClauseContext) {}

// ExitParametersClause is called when production parametersClause is exited.
func (s *BaseLectureListener) ExitParametersClause(ctx *ParametersClauseContext) {}

// EnterParameter is called when production parameter is entered.
func (s *BaseLectureListener) EnterParameter(ctx *ParameterContext) {}

// ExitParameter is called when production parameter is exited.
func (s *BaseLectureListener) ExitParameter(ctx *ParameterContext) {}

// EnterType is called when production type is entered.
func (s *BaseLectureListener) EnterType(ctx *TypeContext) {}

// ExitType is called when production type is exited.
func (s *BaseLectureListener) ExitType(ctx *TypeContext) {}

// EnterOperator is called when production operator is entered.
func (s *BaseLectureListener) EnterOperator(ctx *OperatorContext) {}

// ExitOperator is called when production operator is exited.
func (s *BaseLectureListener) ExitOperator(ctx *OperatorContext) {}

// EnterComparator is called when production comparator is entered.
func (s *BaseLectureListener) EnterComparator(ctx *ComparatorContext) {}

// ExitComparator is called when production comparator is exited.
func (s *BaseLectureListener) ExitComparator(ctx *ComparatorContext) {}

// EnterIdentifier is called when production identifier is entered.
func (s *BaseLectureListener) EnterIdentifier(ctx *IdentifierContext) {}

// ExitIdentifier is called when production identifier is exited.
func (s *BaseLectureListener) ExitIdentifier(ctx *IdentifierContext) {}

// EnterLiteral is called when production literal is entered.
func (s *BaseLectureListener) EnterLiteral(ctx *LiteralContext) {}

// ExitLiteral is called when production literal is exited.
func (s *BaseLectureListener) ExitLiteral(ctx *LiteralContext) {}

// EnterString is called when production string is entered.
func (s *BaseLectureListener) EnterString(ctx *StringContext) {}

// ExitString is called when production string is exited.
func (s *BaseLectureListener) ExitString(ctx *StringContext) {}

// EnterNumber is called when production number is entered.
func (s *BaseLectureListener) EnterNumber(ctx *NumberContext) {}

// ExitNumber is called when production number is exited.
func (s *BaseLectureListener) ExitNumber(ctx *NumberContext) {}
