package models

type Parser struct {
	Filename string
	labels   map[string]int
	tokens   []Token
	output   Ast
	cursor   int
}

func NewParser(filename string, tokens []Token) Parser {
	return Parser{
		Filename: filename,
		output:   Ast{},
		tokens:   tokens,
		cursor:   0,
	}
}

func (this *Parser) WriteProgram() (ret []byte) {
	//vec := mgu.VecMap(this.output.Statements, func(stmt Stmt) []uint16 {
	//    return stmt.WriteMemASM()
	//})
	//for _, bytes := range vec {
	//ret = append(ret, bytes...)
	//}
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
	this.cursor += n
}

const (
	NoSpaceMode = iota
	OptionalSpaceMode
	MandatorySpaceMode
)

func (this *Parser) HasNextConsume(spaceMode int, kinds ...TokenKindEnum) *Token {
	if spaceMode < NoSpaceMode || spaceMode > MandatorySpaceMode {
		panic("invalid argument in function 'HasNextConsume'")
	}
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
