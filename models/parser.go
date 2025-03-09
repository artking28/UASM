package models

import (
	"UASM/neander"
)

type (
	MemHeap struct {
		content map[uint16]int16
		last    uint16
	}

	Parser struct {
		Filename       string
		memHep         MemHeap
		ByteCodeLabels map[string]uint16
		tokens         []Token
		output         Ast
		cursor         int
	}
)

func NewParser(filename string, tokens []Token) Parser {
	// Pega o index da ultima constante declarada
	l := GetLastConstant()
	constants := GetBuiltinConstants()
	return Parser{
		Filename:       filename,
		ByteCodeLabels: map[string]uint16{},
		memHep: MemHeap{
			content: constants,
			last:    l,
		},
		output: Ast{},
		tokens: tokens,
		cursor: 0,
	}
}

func (this *Parser) AllocNum(num int16) uint16 {
	this.memHep.last++
	where := this.memHep.last
	this.memHep.content[where] = num
	return where - NeanderPadding + JmpConstantsSize
}

func (this *Parser) WriteProgram() []uint8 {

	var vec [][]uint16
	var stmtSizes int
	resolveLabel := map[int]string{}
	for _, stmt := range this.output.Statements {
		if stmt.GetTitle() == "LabelDeclStmt" {
			labelStmt := stmt.(LabelDeclStmt)
			this.ByteCodeLabels[labelStmt.LabelName] = uint16(stmtSizes)
		}
		bytes := stmt.WriteMemASM()
		if stmt.GetTitle() == "JumpStmt" {
			jmpStmt := stmt.(JumpStmt)
			resolveLabel[stmtSizes+len(bytes)-1] = jmpStmt.TargetLabelName
		}
		stmtSizes += len(bytes)
		vec = append(vec, bytes)
	}
	var program []uint16
	for _, bytes := range vec {
		program = append(program, bytes...)
	}

	// Prefixo do Neander.
	//constants := GetBuiltinConstants()
	constants := this.memHep.content
	constantsCount := uint16(len(constants))
	neanderPrefix := []uint16{1, 1}
	// O padding de constantes n pode ser impar
	PaddingSize := uint16(len(neanderPrefix)) + constantsCount

	for k := range this.ByteCodeLabels {
		this.ByteCodeLabels[k] += PaddingSize
	}

	for k, v := range resolveLabel {
		program[k] = this.ByteCodeLabels[v]
	}

	// Garante q o programa não vai tentar executar o espaço reservado para constantes
	neanderPrefix = append(neanderPrefix, neander.JMP, PaddingSize)
	// Adiciona as constantes e os espaços de memória reversados.
	neanderPrefix = append(neanderPrefix, make([]uint16, constantsCount)...)
	for k, v := range constants {
		neanderPrefix[k] = uint16(uint8(v))
	}

	// Reúne todas as partes.
	neanderPrefix = append(neanderPrefix, program...)
	final := make([]uint8, len(neanderPrefix)*2)
	for i, num := range neanderPrefix {
		final[i*2+1] = uint8(num >> 8)
		final[i*2] = uint8(num)
		//final[i*2] = uint8(num >> 8)
		//final[i*2+1] = uint8(num)
	}

	// Marca o fim do programa
	endAt := len(final)

	// Itera pra ter 516 de tamanho
	if endAt < 516 {
		final = append(final, make([]uint8, 516-endAt)...)
	}

	return final
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
