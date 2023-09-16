package internal

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/dave/jennifer/jen"
	"github.com/paulwrubel/lecture-lang/parser/go/lecture"
)

// GolangLectureListener is a listener for a parse tree produced by LectureParser.
//
// It will transpile a Lecture into a Golang script
type GolangLectureListener struct {
	*lecture.BaseLectureListener

	OutputFile *os.File
	Errors     []error

	jenFile *jen.File

	// jennifer constructs
	blockStack     *blockStack
	statementStack *statementStack
}

func (l *GolangLectureListener) EnterLecture(ctx *lecture.LectureContext) {
	l.jenFile = jen.NewFile("main")
	l.statementStack = &statementStack{}
	l.blockStack = &blockStack{}
}

func (l *GolangLectureListener) ExitLecture(ctx *lecture.LectureContext) {
	err := l.jenFile.Render(l.OutputFile)
	if err != nil {
		l.Errors = append(l.Errors, err)
	}
}

func (l *GolangLectureListener) EnterMainFunction(ctx *lecture.MainFunctionContext) {
	l.blockStack.Push([]jen.Code{})
	l.statementStack.Push(jen.Func().Id("main").Params())
}

func (l *GolangLectureListener) ExitMainFunction(ctx *lecture.MainFunctionContext) {
	funcStatement := l.statementStack.Pop()
	funcBlock := l.blockStack.Pop()

	funcStatement.Add(jen.Block(funcBlock...))
	l.jenFile.Add(funcStatement)
}

func (l *GolangLectureListener) EnterFunction(ctx *lecture.FunctionContext) {
	l.statementStack.Push(jen.Func())
	l.blockStack.Push([]jen.Code{})
}

func (l *GolangLectureListener) ExitFunction(ctx *lecture.FunctionContext) {
	funcStatement := l.statementStack.Pop()
	funcBlock := l.blockStack.Pop()

	funcStatement.Add(jen.Block(funcBlock...))
	l.jenFile.Add(funcStatement)
}

func (l *GolangLectureListener) EnterFunctionSignature(ctx *lecture.FunctionSignatureContext) {
	funcStatement := l.statementStack.Pop()
	id := ctx.Identifier().GetText()
	funcStatement.Add(jen.Id(id))
	l.statementStack.Push(funcStatement)

	// holds parameter statements
	l.blockStack.Push([]jen.Code{})
}

func (l *GolangLectureListener) ExitFunctionSignature(ctx *lecture.FunctionSignatureContext) {
	funcStatement := l.statementStack.Pop()
	paramsBlock := l.blockStack.Pop()

	funcStatement.Add(jen.Params(paramsBlock...))

	typeString := ctx.Type_().GetText()
	switch strings.ToLower(typeString) {
	case "number":
		funcStatement.Add(jen.Int64())
	default:
		l.Errors = append(l.Errors, fmt.Errorf("unknown type: %s", typeString))
		return
	}

	l.statementStack.Push(funcStatement)
}

func (l *GolangLectureListener) EnterParameterDeclaration(ctx *lecture.ParameterDeclarationContext) {
	funcParamsBlock := l.blockStack.Pop()
	id := ctx.Identifier().GetText()

	typeString := ctx.Type_().GetText()
	switch strings.ToLower(typeString) {
	case "number":
		funcParamsBlock = append(funcParamsBlock, jen.Id(id).Int64())
	default:
		l.Errors = append(l.Errors, fmt.Errorf("unknown type: %s", typeString))
		return
	}

	l.blockStack.Push(funcParamsBlock)
}

func (l *GolangLectureListener) EnterReturnStatement(ctx *lecture.ReturnStatementContext) {
	resultStatement := jen.Null()

	l.statementStack.Push(resultStatement)
}

func (l *GolangLectureListener) ExitReturnStatement(ctx *lecture.ReturnStatementContext) {
	resultStatement := l.statementStack.Pop()
	funcBlock := l.blockStack.Pop()

	funcBlock = append(funcBlock, jen.Return(resultStatement))

	l.blockStack.Push(funcBlock)
}

func (l *GolangLectureListener) EnterAssignmentStatement(ctx *lecture.AssignmentStatementContext) {
	// initialize statement
	varName := ctx.Identifier().GetText()
	statement := jen.Id(varName).Op(":=")

	l.statementStack.Push(statement)
}

func (l *GolangLectureListener) ExitAssignmentStatement(ctx *lecture.AssignmentStatementContext) {
	assignmentStatement := l.statementStack.Pop()
	funcBlock := l.blockStack.Pop()
	funcBlock = append(funcBlock, assignmentStatement)

	l.blockStack.Push(funcBlock)
}

func (l *GolangLectureListener) EnterReassignmentStatement(ctx *lecture.ReassignmentStatementContext) {
	// initialize statement
	varName := ctx.Identifier().GetText()
	statement := jen.Id(varName).Op("=")

	l.statementStack.Push(statement)
}

func (l *GolangLectureListener) ExitReassignmentStatement(ctx *lecture.ReassignmentStatementContext) {
	reassignmentStatement := l.statementStack.Pop()
	funcBlock := l.blockStack.Pop()
	funcBlock = append(funcBlock, reassignmentStatement)

	l.blockStack.Push(funcBlock)
}

func (l *GolangLectureListener) EnterPrintStatement(ctx *lecture.PrintStatementContext) {
	statement := jen.Null()

	l.statementStack.Push(statement)
}

func (l *GolangLectureListener) ExitPrintStatement(ctx *lecture.PrintStatementContext) {
	printArgStatement := l.statementStack.Pop()
	printStatement := jen.Qual("fmt", "Println").Call(printArgStatement)
	funcBlock := l.blockStack.Pop()

	funcBlock = append(funcBlock, printStatement)

	l.blockStack.Push(funcBlock)
}

func (l *GolangLectureListener) EnterIfChainStatement(ctx *lecture.IfChainStatementContext) {
	ifChainStatement := jen.Null()

	l.statementStack.Push(ifChainStatement)
}

func (l *GolangLectureListener) ExitIfChainStatement(ctx *lecture.IfChainStatementContext) {
	ifChainStatement := l.statementStack.Pop()
	funcBlock := l.blockStack.Pop()

	funcBlock = append(funcBlock, ifChainStatement)

	l.blockStack.Push(funcBlock)
}

func (l *GolangLectureListener) EnterIfStatement(ctx *lecture.IfStatementContext) {
	ifBlock := []jen.Code{}

	l.blockStack.Push(ifBlock)
}

func (l *GolangLectureListener) ExitIfStatement(ctx *lecture.IfStatementContext) {
	ifChainStatement := l.statementStack.Pop()
	ifBlock := l.blockStack.Pop()

	ifChainStatement.Add(jen.Block(ifBlock...))

	l.statementStack.Push(ifChainStatement)
}

func (l *GolangLectureListener) EnterIfSignature(ctx *lecture.IfSignatureContext) {
	ifConditionStatement := jen.Null()

	l.statementStack.Push(ifConditionStatement)
}

func (l *GolangLectureListener) ExitIfSignature(ctx *lecture.IfSignatureContext) {
	ifConditionStatement := l.statementStack.Pop()
	ifChainStatement := l.statementStack.Pop()

	ifChainStatement.Add(jen.If(ifConditionStatement))

	l.statementStack.Push(ifChainStatement)
}

func (l *GolangLectureListener) EnterElseIfStatement(ctx *lecture.ElseIfStatementContext) {
	elseIfBlock := []jen.Code{}

	l.blockStack.Push(elseIfBlock)
}

func (l *GolangLectureListener) ExitElseIfStatement(ctx *lecture.ElseIfStatementContext) {
	ifChainStatement := l.statementStack.Pop()
	elseIfBlock := l.blockStack.Pop()

	ifChainStatement.Add(jen.Block(elseIfBlock...))

	l.statementStack.Push(ifChainStatement)
}

func (l *GolangLectureListener) EnterElseIfSignature(ctx *lecture.ElseIfSignatureContext) {
	ifChainStatement := l.statementStack.Pop()
	ifChainStatement.Add(jen.Else())
	l.statementStack.Push(ifChainStatement)
}

func (l *GolangLectureListener) EnterElseStatement(ctx *lecture.ElseStatementContext) {
	elseBlock := []jen.Code{}

	l.blockStack.Push(elseBlock)
}

func (l *GolangLectureListener) ExitElseStatement(ctx *lecture.ElseStatementContext) {
	ifChainStatement := l.statementStack.Pop()
	elseBlock := l.blockStack.Pop()

	ifChainStatement.Add(jen.Block(elseBlock...))

	l.statementStack.Push(ifChainStatement)
}

func (l *GolangLectureListener) ExitElseSignature(ctx *lecture.ElseSignatureContext) {
	ifChainStatement := l.statementStack.Pop()

	ifChainStatement.Add(jen.Else())

	l.statementStack.Push(ifChainStatement)
}

func (l *GolangLectureListener) EnterValue(ctx *lecture.ValueContext) {
	valuesStatement := l.statementStack.Pop()

	if ctx.Identifier() != nil {
		// we are an identifier
		valuesStatement.Add(jen.Id(ctx.GetText()))
	} else if ctx.LiteralClause() != nil {
		// we are a literal
		literalInt64, err := strconv.ParseInt(ctx.LiteralClause().Literal().GetText(), 10, 64)
		if err != nil {
			l.Errors = append(l.Errors, err)
			return
		}
		valuesStatement.Add(jen.Lit(literalInt64))
	} else {
		l.Errors = append(l.Errors, errors.New("value is not identifier or literal (shouldn't be possible)"))
	}

	l.statementStack.Push(valuesStatement)
}

func (l *GolangLectureListener) EnterFunctionCall(ctx *lecture.FunctionCallContext) {
	paramsBlock := []jen.Code{}

	l.blockStack.Push(paramsBlock)
}

func (l *GolangLectureListener) ExitFunctionCall(ctx *lecture.FunctionCallContext) {
	assignmentStatement := l.statementStack.Pop()
	paramsBlock := l.blockStack.Pop()

	id := ctx.Identifier().GetText()
	funcCallStatement := jen.Id(id).Call(paramsBlock...)

	assignmentStatement.Add(funcCallStatement)

	l.statementStack.Push(assignmentStatement)
}
func (l *GolangLectureListener) EnterParameter(ctx *lecture.ParameterContext) {
	paramStatement := jen.Null()

	l.statementStack.Push(paramStatement)
}

func (l *GolangLectureListener) ExitParameter(ctx *lecture.ParameterContext) {
	paramStatement := l.statementStack.Pop()
	paramsBlock := l.blockStack.Pop()

	paramsBlock = append(paramsBlock, paramStatement)

	l.blockStack.Push(paramsBlock)
}

func (l *GolangLectureListener) EnterOperator(ctx *lecture.OperatorContext) {
	statement := l.statementStack.Pop()

	operatorString := ctx.GetText()
	var op string
	switch strings.ToLower(operatorString) {
	case "plus":
		op = "+"
	case "minus":
		op = "-"
	default:
		l.Errors = append(l.Errors, fmt.Errorf("unknown operator: %s", operatorString))
		return
	}
	statement.Add(jen.Op(op))

	l.statementStack.Push(statement)
}

func (l *GolangLectureListener) EnterComparator(ctx *lecture.ComparatorContext) {
	statement := l.statementStack.Pop()

	comparatorString := ctx.GetText()
	var comp string
	switch strings.ToLower(comparatorString) {
	case "is":
		comp = "=="
	default:
		l.Errors = append(l.Errors, fmt.Errorf("unknown comparator: %s", comparatorString))
		return
	}
	statement.Add(jen.Op(comp))

	l.statementStack.Push(statement)
}
