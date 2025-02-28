package models

type (
	TokenKindEnum int

	Token struct {
		Kind   TokenKindEnum
		Value  []rune
		Repeat int
	}
)

func NewToken(kind TokenKindEnum, repeat int, value ...rune) Token {
	return Token{kind, value, repeat}
}

func AppendToken(tokens *[]Token, token Token) {
	if tokens == nil {
		tokens = &[]Token{}
	}
	len := len((*tokens))
	if len > 0 && (*tokens)[len-1].Kind == token.Kind && string((*tokens)[len-1].Value) == string(token.Value) {
		(*tokens)[len-1].Repeat = (*tokens)[len-1].Repeat + 1
		return
	}
	*tokens = append(*tokens, token)
}

const (

	// #########################
	//       Normal tokens
	// #########################
	TOKEN_SPACE TokenKindEnum = iota
	TOKEN_BREAK_LINE
	TOKEN_TAB
	TOKEN_ID
	TOKEN_NUMBER
	TOKEN_COMMA
	TOKEN_COLON
	TOKEN_MEM
	TOKEN_SLASH

	// #########################
	//         MINMONICS
	// #########################

	// Memory manipulations
	TOKEN_GET
	TOKEN_SET
	TOKEN_CPY

	// Simple operations
	TOKEN_INC
	TOKEN_DEC
	TOKEN_NEG
	TOKEN_NOT

	// Operations
	TOKEN_ADD
	TOKEN_MUL
	TOKEN_AND
	TOKEN_ORR
	TOKEN_XOR

	// Loops and comparations
	TOKEN_CMP
	TOKEN_JMP
	TOKEN_JIZ
	TOKEN_JNZ

	// Runtime actions
	TOKEN_HLT
)
