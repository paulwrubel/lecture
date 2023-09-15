package main

import (
	"os"
	"strconv"

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
	varName := ctx.Variable().GetText()
	varValueString := ctx.Value().Number().GetText()

	varValue, err := strconv.ParseInt(varValueString, 10, 64)
	if err != nil {
		l.Errors = append(l.Errors, err)
		return
	}

	l.currentStatements = append(l.currentStatements, jen.Id(varName).Op(":=").Lit(varValue))
}

func (l *GolangLectureListener) EnterPrintStatement(ctx *lecture.PrintStatementContext) {
	varName := ctx.Variable().GetText()

	l.currentStatements = append(l.currentStatements, jen.Qual("fmt", "Println").Call(jen.Id(varName)))
}
