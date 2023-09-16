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

	// EnterStartClause is called when entering the startClause production.
	EnterStartClause(c *StartClauseContext)

	// EnterEndClause is called when entering the endClause production.
	EnterEndClause(c *EndClauseContext)

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// EnterAtomicStatement is called when entering the atomicStatement production.
	EnterAtomicStatement(c *AtomicStatementContext)

	// EnterFunction is called when entering the function production.
	EnterFunction(c *FunctionContext)

	// EnterFunctionStatement is called when entering the functionStatement production.
	EnterFunctionStatement(c *FunctionStatementContext)

	// EnterReturnStatement is called when entering the returnStatement production.
	EnterReturnStatement(c *ReturnStatementContext)

	// EnterDeclarationStatement is called when entering the declarationStatement production.
	EnterDeclarationStatement(c *DeclarationStatementContext)

	// EnterPrintStatement is called when entering the printStatement production.
	EnterPrintStatement(c *PrintStatementContext)

	// EnterParametersDeclaration is called when entering the parametersDeclaration production.
	EnterParametersDeclaration(c *ParametersDeclarationContext)

	// EnterParameterDeclarationClause is called when entering the parameterDeclarationClause production.
	EnterParameterDeclarationClause(c *ParameterDeclarationClauseContext)

	// EnterFunctionCall is called when entering the functionCall production.
	EnterFunctionCall(c *FunctionCallContext)

	// EnterParametersClause is called when entering the parametersClause production.
	EnterParametersClause(c *ParametersClauseContext)

	// EnterValueClause is called when entering the valueClause production.
	EnterValueClause(c *ValueClauseContext)

	// EnterLiteralClause is called when entering the literalClause production.
	EnterLiteralClause(c *LiteralClauseContext)

	// EnterParameterDeclaration is called when entering the parameterDeclaration production.
	EnterParameterDeclaration(c *ParameterDeclarationContext)

	// EnterValue is called when entering the value production.
	EnterValue(c *ValueContext)

	// EnterLiteral is called when entering the literal production.
	EnterLiteral(c *LiteralContext)

	// EnterType is called when entering the type production.
	EnterType(c *TypeContext)

	// EnterOperator is called when entering the operator production.
	EnterOperator(c *OperatorContext)

	// EnterIdentifier is called when entering the identifier production.
	EnterIdentifier(c *IdentifierContext)

	// EnterNumber is called when entering the number production.
	EnterNumber(c *NumberContext)

	// ExitLecture is called when exiting the lecture production.
	ExitLecture(c *LectureContext)

	// ExitProgram is called when exiting the program production.
	ExitProgram(c *ProgramContext)

	// ExitStartClause is called when exiting the startClause production.
	ExitStartClause(c *StartClauseContext)

	// ExitEndClause is called when exiting the endClause production.
	ExitEndClause(c *EndClauseContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitAtomicStatement is called when exiting the atomicStatement production.
	ExitAtomicStatement(c *AtomicStatementContext)

	// ExitFunction is called when exiting the function production.
	ExitFunction(c *FunctionContext)

	// ExitFunctionStatement is called when exiting the functionStatement production.
	ExitFunctionStatement(c *FunctionStatementContext)

	// ExitReturnStatement is called when exiting the returnStatement production.
	ExitReturnStatement(c *ReturnStatementContext)

	// ExitDeclarationStatement is called when exiting the declarationStatement production.
	ExitDeclarationStatement(c *DeclarationStatementContext)

	// ExitPrintStatement is called when exiting the printStatement production.
	ExitPrintStatement(c *PrintStatementContext)

	// ExitParametersDeclaration is called when exiting the parametersDeclaration production.
	ExitParametersDeclaration(c *ParametersDeclarationContext)

	// ExitParameterDeclarationClause is called when exiting the parameterDeclarationClause production.
	ExitParameterDeclarationClause(c *ParameterDeclarationClauseContext)

	// ExitFunctionCall is called when exiting the functionCall production.
	ExitFunctionCall(c *FunctionCallContext)

	// ExitParametersClause is called when exiting the parametersClause production.
	ExitParametersClause(c *ParametersClauseContext)

	// ExitValueClause is called when exiting the valueClause production.
	ExitValueClause(c *ValueClauseContext)

	// ExitLiteralClause is called when exiting the literalClause production.
	ExitLiteralClause(c *LiteralClauseContext)

	// ExitParameterDeclaration is called when exiting the parameterDeclaration production.
	ExitParameterDeclaration(c *ParameterDeclarationContext)

	// ExitValue is called when exiting the value production.
	ExitValue(c *ValueContext)

	// ExitLiteral is called when exiting the literal production.
	ExitLiteral(c *LiteralContext)

	// ExitType is called when exiting the type production.
	ExitType(c *TypeContext)

	// ExitOperator is called when exiting the operator production.
	ExitOperator(c *OperatorContext)

	// ExitIdentifier is called when exiting the identifier production.
	ExitIdentifier(c *IdentifierContext)

	// ExitNumber is called when exiting the number production.
	ExitNumber(c *NumberContext)
}
