package models

import (
	mgu "github.com/artking28/myGoUtils"
)

type Parser struct {
	output []Stmt
	tokens []Token
	cursor int
	column int
	line   int
}

func NewParser(tokens []Token) Parser {
	return Parser{
		tokens: tokens,
		cursor: 0,
		column: 0,
		line:   1,
	}
}

func (this *Parser) WriteProgram() (ret []byte) {
	vec := mgu.VecMap(this.output, func(stmt Stmt) []byte {
		return stmt.WriteMemASM()
	})
	for _, bytes := range vec {
		ret = append(ret, bytes...)
	}
	return ret
}

func (this *Parser) Inject(stmts ...Stmt) {
	this.output = append(this.output, stmts...)
}

func (this *Parser) Get(n int) *Token {
	if this.cursor+n >= len(this.tokens) {
		return nil
	}
	return &this.tokens[this.cursor+n]
}

func (this *Parser) Where() Pos {
	return Pos{
		Column: int64(this.column),
		Line:   int64(this.line),
	}
}

func (this *Parser) NextLine() {
	this.column = 0
	this.line += 1
}

func (this *Parser) Consume(n int) {
	if this.cursor >= len(this.tokens) {
		return
	}
	this.column += n
	this.cursor += n
}
