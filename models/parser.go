package models

import (
	mgu "github.com/artking28/myGoUtils"
)

type Parser struct {
	output Ast
	tokens []Token
	cursor int
	column int
	line   int
}

func NewParser(tokens []Token) Parser {
	return Parser{
		output: Ast{},
		tokens: tokens,
		cursor: 0,
	}
}

func (this *Parser) WriteProgram() (ret []byte) {
	vec := mgu.VecMap(this.output.Statements, func(stmt Stmt) []byte {
		return stmt.WriteMemASM()
	})
	for _, bytes := range vec {
		ret = append(ret, bytes...)
	}
	return ret
}

func (this *Parser) Inject(stmts ...Stmt) {
	this.output.Statements = append(this.output.Statements, stmts...)
}

func (this *Parser) Inspect() {
	this.output.Inspect()
}

func (this *Parser) Get(n int) *Token {
	if this.cursor+n >= len(this.tokens) {
		return nil
	}
	return &this.tokens[this.cursor+n]
}

func (this *Parser) Consume(n int) {
	if this.cursor >= len(this.tokens) {
		return
	}
	for i := 0; i < n; i++ {
		tk := this.tokens[this.cursor+i]
		if tk.Kind == TOKEN_BREAK_LINE {
			this.column = 1
			this.line += 1
		}
		this.column += len(tk.Value)
		if tk.Kind == TOKEN_MEM {
			this.column++
		}
	}
	this.cursor += n
}

const (
	NoSpaceMode = iota
	OptionalSpaceMode
	MandatorySpaceMode
)

func (this *Parser) HasNextConsume(spaceMode int, kinds ...TokenKindEnum) *Token {
	for findSpace := false; ; {
		token := this.Get(0)
		if token == nil {
			// Fim dos tokens sem encontrar um tipo esperado
			return nil
		}

		for _, kind := range kinds {
			if token.Kind == kind {
				// Se espaços eram obrigatórios mas não foram encontrados, falha
				if spaceMode == MandatorySpaceMode && !findSpace {
					return nil
				}
				this.Consume(1)
				return token
			}
		}

		if token.Kind == TOKEN_SPACE {
			findSpace = true
			this.Consume(1)
			continue
		}

		// Se espaços não eram permitidos ou eram obrigatórios e encontrou outro token, falha
		if spaceMode == NoSpaceMode || spaceMode == MandatorySpaceMode {
			return nil
		}

		return nil // Qualquer outro caso não esperado falha
	}
}
