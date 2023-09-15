package main

import (
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
	currentFuncName   string
	currentStatements []jen.Code
	currentStatement  *jen.Statement
}

func (l *GolangLectureListener) EnterLecture(ctx *lecture.LectureContext) {
	l.jenFile = jen.NewFile("main")
}

func (l *GolangLectureListener) ExitLecture(ctx *lecture.LectureContext) {
	err := l.jenFile.Render(l.OutputFile)
	if err != nil {
		l.Errors = append(l.Errors, err)
	}
}

func (l GolangLectureListener) EnterStartClause(ctx *lecture.StartClauseContext) {
	l.currentFuncName = "main"
	l.currentStatements = []jen.Code{}
}

func (l GolangLectureListener) EnterEndClause(ctx *lecture.EndClauseContext) {
	l.jenFile.Func().Id("main").Params().Block(l.currentStatements...)
}

func (l *GolangLectureListener) EnterDeclarationStatement(ctx *lecture.DeclarationStatementContext) {
	// initialize statement
	varName := ctx.Identifier().GetText()
	statement := jen.Id(varName).Op(":=")

	l.currentStatement = statement
}

func (l *GolangLectureListener) ExitDeclarationStatement(ctx *lecture.DeclarationStatementContext) {
	l.currentStatements = append(l.currentStatements, l.currentStatement)
	l.currentStatement = nil
}

func (l *GolangLectureListener) EnterPrintStatement(ctx *lecture.PrintStatementContext) {
	l.currentStatement = jen.Null()
}

func (l *GolangLectureListener) ExitPrintStatement(ctx *lecture.PrintStatementContext) {
	statement := jen.Qual("fmt", "Println").Call(l.currentStatement)

	l.currentStatements = append(l.currentStatements, statement)
	l.currentStatement = nil
}

func (l *GolangLectureListener) EnterValue(ctx *lecture.ValueContext) {
	if ctx.Identifier() != nil {
		// we are an identifier
		l.currentStatement.Add(jen.Id(ctx.GetText()))
	} else {
		// we are a literal
		literalInt64, err := strconv.ParseInt(ctx.LiteralClause().Literal().GetText(), 10, 64)
		if err != nil {
			l.Errors = append(l.Errors, err)
			return
		}
		l.currentStatement.Add(jen.Lit(literalInt64))
	}
}

func (l *GolangLectureListener) EnterOperator(ctx *lecture.OperatorContext) {
	operatorString := ctx.GetText()
	var op string
	switch strings.ToLower(operatorString) {
	case "plus":
		op = "+"
	default:
		l.Errors = append(l.Errors, fmt.Errorf("unknown operator: %s", operatorString))
		return
	}
	l.currentStatement.Add(jen.Op(op))
}
