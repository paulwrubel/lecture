package internal

import "github.com/dave/jennifer/jen"

type statementStack struct {
	statements []*jen.Statement
}

func (s *statementStack) Pop() *jen.Statement {
	if len(s.statements) == 0 {
		return nil
	}
	stmt := s.statements[len(s.statements)-1]
	s.statements = s.statements[:len(s.statements)-1]
	return stmt
}

func (s *statementStack) Push(stmt *jen.Statement) {
	s.statements = append(s.statements, stmt)
}

type blockStack struct {
	blocks [][]jen.Code
}

func (s *blockStack) Pop() []jen.Code {
	if len(s.blocks) == 0 {
		return nil
	}
	block := s.blocks[len(s.blocks)-1]
	s.blocks = s.blocks[:len(s.blocks)-1]
	return block
}

func (s *blockStack) Push(blck []jen.Code) {
	s.blocks = append(s.blocks, blck)
}
