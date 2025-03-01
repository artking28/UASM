package models

import (
	"fmt"
	"strconv"
	"strings"
)

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
	count := len(*tokens)
	if count > 0 && (*tokens)[count-1].Kind == token.Kind && string((*tokens)[count-1].Value) == string(token.Value) {
		(*tokens)[count-1].Repeat = (*tokens)[count-1].Repeat + 1
		return
	}
	*tokens = append(*tokens, token)
}

func ResolveTokenId(token Token) Token {
	if token.Kind != TOKEN_ID {
		return token
	}
	value := string(token.Value)
	count := len(value)
	if value[0] == 'm' {
		mem, err := strconv.ParseInt(value[1:], 10, 64)
		if err == nil {
			return NewToken(TOKEN_MEM, 1, rune(mem))
		}
	} else if count > 2 && (value[:2] == "0b" || value[:2] == "0o" || value[:2] == "0x") {
		num, err := strconv.ParseInt(value[:count], 0, 64)
		if err == nil {
			return NewToken(TOKEN_NUMBER, 1, rune(num))
		}
	} else if strings.ToUpper(value) == ("GET") {
		return NewToken(TOKEN_GET, 1, []rune(value)...)
	} else if strings.ToUpper(value) == ("SET") {
		return NewToken(TOKEN_SET, 1, []rune(value)...)
	} else if strings.ToUpper(value) == ("CPY") {
		return NewToken(TOKEN_CPY, 1, []rune(value)...)
	} else if strings.ToUpper(value) == ("INC") {
		return NewToken(TOKEN_INC, 1, []rune(value)...)
	} else if strings.ToUpper(value) == ("DEC") {
		return NewToken(TOKEN_DEC, 1, []rune(value)...)
	} else if strings.ToUpper(value) == ("NEG") {
		return NewToken(TOKEN_NEG, 1, []rune(value)...)
	} else if strings.ToUpper(value) == ("NOT") {
		return NewToken(TOKEN_NOT, 1, []rune(value)...)
	} else if strings.ToUpper(value) == ("ADD") {
		return NewToken(TOKEN_ADD, 1, []rune(value)...)
	} else if strings.ToUpper(value) == ("MUL") {
		return NewToken(TOKEN_MUL, 1, []rune(value)...)
	} else if strings.ToUpper(value) == ("AND") {
		return NewToken(TOKEN_AND, 1, []rune(value)...)
	} else if strings.ToUpper(value) == ("ORR") {
		return NewToken(TOKEN_ORR, 1, []rune(value)...)
	} else if strings.ToUpper(value) == ("XOR") {
		return NewToken(TOKEN_XOR, 1, []rune(value)...)
	} else if strings.ToUpper(value) == ("CMP") {
		return NewToken(TOKEN_CMP, 1, []rune(value)...)
	} else if strings.ToUpper(value) == ("JMP") {
		return NewToken(TOKEN_JMP, 1, []rune(value)...)
	} else if strings.ToUpper(value) == ("JIZ") {
		return NewToken(TOKEN_JIZ, 1, []rune(value)...)
	} else if strings.ToUpper(value) == ("JNZ") {
		return NewToken(TOKEN_JNZ, 1, []rune(value)...)
	} else if strings.ToUpper(value) == ("HLT") {
		return NewToken(TOKEN_HLT, 1, []rune(value)...)
	}

	return token
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
	TOKEN_EOF

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

func (this Token) String() string {
	var s string
	switch this.Kind {
	case TOKEN_SPACE:
		s = "TOKEN_SPACE"
		break
	case TOKEN_BREAK_LINE:
		s = "TOKEN_BREAK_LINE"
		break
	case TOKEN_TAB:
		s = "TOKEN_TAB"
		break
	case TOKEN_ID:
		s = "TOKEN_ID"
		break
	case TOKEN_NUMBER:
		s = "TOKEN_NUMBER"
		break
	case TOKEN_COMMA:
		s = "TOKEN_COMMA"
		break
	case TOKEN_COLON:
		s = "TOKEN_COLON"
		break
	case TOKEN_MEM:
		s = "TOKEN_MEM"
		break
	case TOKEN_SLASH:
		s = "TOKEN_SLASH"
		break
	case TOKEN_EOF:
		s = "TOKEN_EOF"
		break
	case TOKEN_GET:
		s = "TOKEN_GET"
		break
	case TOKEN_SET:
		s = "TOKEN_SET"
		break
	case TOKEN_CPY:
		s = "TOKEN_CPY"
		break
	case TOKEN_INC:
		s = "TOKEN_INC"
		break
	case TOKEN_DEC:
		s = "TOKEN_DEC"
		break
	case TOKEN_NEG:
		s = "TOKEN_NEG"
		break
	case TOKEN_NOT:
		s = "TOKEN_NOT"
		break
	case TOKEN_ADD:
		s = "TOKEN_ADD"
		break
	case TOKEN_MUL:
		s = "TOKEN_MUL"
		break
	case TOKEN_AND:
		s = "TOKEN_AND"
		break
	case TOKEN_ORR:
		s = "TOKEN_ORR"
		break
	case TOKEN_XOR:
		s = "TOKEN_XOR"
		break
	case TOKEN_CMP:
		s = "TOKEN_CMP"
		break
	case TOKEN_JMP:
		s = "TOKEN_JMP"
		break
	case TOKEN_JIZ:
		s = "TOKEN_JIZ"
		break
	case TOKEN_JNZ:
		s = "TOKEN_JNZ"
		break
	case TOKEN_HLT:
		s = "TOKEN_HLT"
		break
	}
	v := string(this.Value)
	if this.Kind == TOKEN_BREAK_LINE {
		v = "\\n"
	} else if this.Kind == TOKEN_TAB {
		v = "\\t"
	} else if this.Kind == TOKEN_EOF {
		v = "\\0"
	} else if this.Kind == TOKEN_MEM || this.Kind == TOKEN_NUMBER {
		v = strconv.Itoa(int(this.Value[0]))
	}
	return fmt.Sprintf("Token{%s, \"%s\", %.2d}\n", s, v, this.Repeat)

}
