package internal

import (
	"github.com/antlr4-go/antlr/v4"
	"github.com/paulwrubel/lecture-lang/parser/go/lecture"
)

type LectureWalker struct {
	Context lecture.ILectureContext
}

func NewLectureWalker(lectureString string) (*LectureWalker, error) {
	// setup the input stream and errorListener
	inputStream := antlr.NewInputStream(lectureString)
	errorListener := NewReasonLangErrorListener()

	// create the Lexer
	lexer := lecture.NewLectureLexer(inputStream)
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(errorListener)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// create the String Parser
	parser := lecture.NewLectureParser(stream)

	// add Custom Error Listener
	parser.RemoveErrorListeners()
	parser.AddErrorListener(errorListener)

	// parse the programString
	programContext := parser.Lecture()

	// checks if SyntaxError Callback Was Invoked
	if errorListener.HadSynError {
		return &LectureWalker{}, errorListener.SynError
	}

	return &LectureWalker{
		Context: programContext,
	}, nil
}

func (w *LectureWalker) Walk(listener antlr.ParseTreeListener) {
	antlr.ParseTreeWalkerDefault.Walk(listener, w.Context)

}
