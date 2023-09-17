// Code generated from Lecture.g4 by ANTLR 4.13.1. DO NOT EDIT.

package lecture // Lecture
import "github.com/antlr4-go/antlr/v4"

// LectureListener is a complete listener for a parse tree produced by LectureParser.
type LectureListener interface {
	antlr.ParseTreeListener

	// EnterLecture is called when entering the lecture production.
	EnterLecture(c *LectureContext)

	// EnterProgram is called when entering the program production.
	EnterProgram(c *ProgramContext)

	// EnterFloatingComment is called when entering the floatingComment production.
	EnterFloatingComment(c *FloatingCommentContext)

	// EnterMainFunction is called when entering the mainFunction production.
	EnterMainFunction(c *MainFunctionContext)

	// EnterMainStartStatement is called when entering the mainStartStatement production.
	EnterMainStartStatement(c *MainStartStatementContext)

	// EnterMainEndStatement is called when entering the mainEndStatement production.
	EnterMainEndStatement(c *MainEndStatementContext)

	// EnterFunction is called when entering the function production.
	EnterFunction(c *FunctionContext)

	// EnterFunctionSignature is called when entering the functionSignature production.
	EnterFunctionSignature(c *FunctionSignatureContext)

	// EnterParametersDeclaration is called when entering the parametersDeclaration production.
	EnterParametersDeclaration(c *ParametersDeclarationContext)

	// EnterParameterDeclarationClause is called when entering the parameterDeclarationClause production.
	EnterParameterDeclarationClause(c *ParameterDeclarationClauseContext)

	// EnterParameterDeclaration is called when entering the parameterDeclaration production.
	EnterParameterDeclaration(c *ParameterDeclarationContext)

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// EnterStatementBlock is called when entering the statementBlock production.
	EnterStatementBlock(c *StatementBlockContext)

	// EnterReturnStatement is called when entering the returnStatement production.
	EnterReturnStatement(c *ReturnStatementContext)

	// EnterAssignmentStatement is called when entering the assignmentStatement production.
	EnterAssignmentStatement(c *AssignmentStatementContext)

	// EnterReassignmentStatement is called when entering the reassignmentStatement production.
	EnterReassignmentStatement(c *ReassignmentStatementContext)

	// EnterPrintStatement is called when entering the printStatement production.
	EnterPrintStatement(c *PrintStatementContext)

	// EnterCommentStatement is called when entering the commentStatement production.
	EnterCommentStatement(c *CommentStatementContext)

	// EnterIfChainStatement is called when entering the ifChainStatement production.
	EnterIfChainStatement(c *IfChainStatementContext)

	// EnterIfStatement is called when entering the ifStatement production.
	EnterIfStatement(c *IfStatementContext)

	// EnterIfSignature is called when entering the ifSignature production.
	EnterIfSignature(c *IfSignatureContext)

	// EnterElseIfStatement is called when entering the elseIfStatement production.
	EnterElseIfStatement(c *ElseIfStatementContext)

	// EnterElseIfSignature is called when entering the elseIfSignature production.
	EnterElseIfSignature(c *ElseIfSignatureContext)

	// EnterElseStatement is called when entering the elseStatement production.
	EnterElseStatement(c *ElseStatementContext)

	// EnterElseSignature is called when entering the elseSignature production.
	EnterElseSignature(c *ElseSignatureContext)

	// EnterIfClosingStatement is called when entering the ifClosingStatement production.
	EnterIfClosingStatement(c *IfClosingStatementContext)

	// EnterConditionClause is called when entering the conditionClause production.
	EnterConditionClause(c *ConditionClauseContext)

	// EnterValueClause is called when entering the valueClause production.
	EnterValueClause(c *ValueClauseContext)

	// EnterValue is called when entering the value production.
	EnterValue(c *ValueContext)

	// EnterLiteralClause is called when entering the literalClause production.
	EnterLiteralClause(c *LiteralClauseContext)

	// EnterFunctionCall is called when entering the functionCall production.
	EnterFunctionCall(c *FunctionCallContext)

	// EnterParametersClause is called when entering the parametersClause production.
	EnterParametersClause(c *ParametersClauseContext)

	// EnterParameter is called when entering the parameter production.
	EnterParameter(c *ParameterContext)

	// EnterType is called when entering the type production.
	EnterType(c *TypeContext)

	// EnterOperator is called when entering the operator production.
	EnterOperator(c *OperatorContext)

	// EnterComparator is called when entering the comparator production.
	EnterComparator(c *ComparatorContext)

	// EnterIdentifier is called when entering the identifier production.
	EnterIdentifier(c *IdentifierContext)

	// EnterLiteral is called when entering the literal production.
	EnterLiteral(c *LiteralContext)

	// EnterString is called when entering the string production.
	EnterString(c *StringContext)

	// EnterNumber is called when entering the number production.
	EnterNumber(c *NumberContext)

	// ExitLecture is called when exiting the lecture production.
	ExitLecture(c *LectureContext)

	// ExitProgram is called when exiting the program production.
	ExitProgram(c *ProgramContext)

	// ExitFloatingComment is called when exiting the floatingComment production.
	ExitFloatingComment(c *FloatingCommentContext)

	// ExitMainFunction is called when exiting the mainFunction production.
	ExitMainFunction(c *MainFunctionContext)

	// ExitMainStartStatement is called when exiting the mainStartStatement production.
	ExitMainStartStatement(c *MainStartStatementContext)

	// ExitMainEndStatement is called when exiting the mainEndStatement production.
	ExitMainEndStatement(c *MainEndStatementContext)

	// ExitFunction is called when exiting the function production.
	ExitFunction(c *FunctionContext)

	// ExitFunctionSignature is called when exiting the functionSignature production.
	ExitFunctionSignature(c *FunctionSignatureContext)

	// ExitParametersDeclaration is called when exiting the parametersDeclaration production.
	ExitParametersDeclaration(c *ParametersDeclarationContext)

	// ExitParameterDeclarationClause is called when exiting the parameterDeclarationClause production.
	ExitParameterDeclarationClause(c *ParameterDeclarationClauseContext)

	// ExitParameterDeclaration is called when exiting the parameterDeclaration production.
	ExitParameterDeclaration(c *ParameterDeclarationContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitStatementBlock is called when exiting the statementBlock production.
	ExitStatementBlock(c *StatementBlockContext)

	// ExitReturnStatement is called when exiting the returnStatement production.
	ExitReturnStatement(c *ReturnStatementContext)

	// ExitAssignmentStatement is called when exiting the assignmentStatement production.
	ExitAssignmentStatement(c *AssignmentStatementContext)

	// ExitReassignmentStatement is called when exiting the reassignmentStatement production.
	ExitReassignmentStatement(c *ReassignmentStatementContext)

	// ExitPrintStatement is called when exiting the printStatement production.
	ExitPrintStatement(c *PrintStatementContext)

	// ExitCommentStatement is called when exiting the commentStatement production.
	ExitCommentStatement(c *CommentStatementContext)

	// ExitIfChainStatement is called when exiting the ifChainStatement production.
	ExitIfChainStatement(c *IfChainStatementContext)

	// ExitIfStatement is called when exiting the ifStatement production.
	ExitIfStatement(c *IfStatementContext)

	// ExitIfSignature is called when exiting the ifSignature production.
	ExitIfSignature(c *IfSignatureContext)

	// ExitElseIfStatement is called when exiting the elseIfStatement production.
	ExitElseIfStatement(c *ElseIfStatementContext)

	// ExitElseIfSignature is called when exiting the elseIfSignature production.
	ExitElseIfSignature(c *ElseIfSignatureContext)

	// ExitElseStatement is called when exiting the elseStatement production.
	ExitElseStatement(c *ElseStatementContext)

	// ExitElseSignature is called when exiting the elseSignature production.
	ExitElseSignature(c *ElseSignatureContext)

	// ExitIfClosingStatement is called when exiting the ifClosingStatement production.
	ExitIfClosingStatement(c *IfClosingStatementContext)

	// ExitConditionClause is called when exiting the conditionClause production.
	ExitConditionClause(c *ConditionClauseContext)

	// ExitValueClause is called when exiting the valueClause production.
	ExitValueClause(c *ValueClauseContext)

	// ExitValue is called when exiting the value production.
	ExitValue(c *ValueContext)

	// ExitLiteralClause is called when exiting the literalClause production.
	ExitLiteralClause(c *LiteralClauseContext)

	// ExitFunctionCall is called when exiting the functionCall production.
	ExitFunctionCall(c *FunctionCallContext)

	// ExitParametersClause is called when exiting the parametersClause production.
	ExitParametersClause(c *ParametersClauseContext)

	// ExitParameter is called when exiting the parameter production.
	ExitParameter(c *ParameterContext)

	// ExitType is called when exiting the type production.
	ExitType(c *TypeContext)

	// ExitOperator is called when exiting the operator production.
	ExitOperator(c *OperatorContext)

	// ExitComparator is called when exiting the comparator production.
	ExitComparator(c *ComparatorContext)

	// ExitIdentifier is called when exiting the identifier production.
	ExitIdentifier(c *IdentifierContext)

	// ExitLiteral is called when exiting the literal production.
	ExitLiteral(c *LiteralContext)

	// ExitString is called when exiting the string production.
	ExitString(c *StringContext)

	// ExitNumber is called when exiting the number production.
	ExitNumber(c *NumberContext)
}
