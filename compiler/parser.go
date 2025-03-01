package compiler

import (
	"UASM/models"
	"errors"
)

func ParseAll(parser *models.Parser) error {

	tk := parser.Get(0)
	for ; tk != nil && tk.Kind != models.TOKEN_EOF; tk = parser.Get(0) {
		switch tk.Kind {
		case models.TOKEN_SLASH:
			err := ParseComment(parser)
			if err != nil {
				return err
			}
			parser.NextLine()
			break
		case models.TOKEN_INC, models.TOKEN_DEC, models.TOKEN_NEG, models.TOKEN_NOT, models.TOKEN_HLT:
			err := ParsePureInstruction(parser)
			if err != nil {
				return err
			}
			break

		case models.TOKEN_GET, models.TOKEN_SET, models.TOKEN_ADD, models.TOKEN_MUL, models.TOKEN_AND, models.TOKEN_ORR, models.TOKEN_XOR, models.TOKEN_CMP, models.TOKEN_JMP, models.TOKEN_JIZ, models.TOKEN_JNZ:
			err := ParseSingleInstruction(parser)
			if err != nil {
				return err
			}
			break
		case models.TOKEN_CPY:
			err := ParseDoubleInstruction(parser)
			if err != nil {
				return err
			}
			break
		case models.TOKEN_BREAK_LINE:
			parser.NextLine()
			break
		default:
			break
		}
		parser.Consume(1)
	}
	return nil
}

func ParseComment(parser *models.Parser) error {
	pos := parser.Where()
	var comment string
	h0 := parser.Get(0)
	if h0 != nil && h0.Kind == models.TOKEN_SLASH && h0.Repeat >= 2 {
		parser.Consume(2)
	} else {
		return errors.New(string(models.UnexpectedTokenErrLabel))
	}
	for here := parser.Get(0); here != nil && here.Kind != models.TOKEN_BREAK_LINE; here = parser.Get(0) {
		comment += string(here.Value)
		parser.Consume(1)
	}
	parser.Inject(models.NewCommentStmt(comment, pos))
	return nil
}

func ParsePureInstruction(parser *models.Parser) error {
	pos := parser.Where()
	h0 := parser.Get(0)
	if h0 == nil {
		return errors.New(string(models.ExpectedTokenErrLabel))
	}
	switch h0.Kind {
	case models.TOKEN_INC, models.TOKEN_DEC, models.TOKEN_NEG, models.TOKEN_NOT, models.TOKEN_HLT:
		parser.Inject(models.NewPureInstructionStmt(h0.Kind, pos))
		parser.Consume(1)
		break
	default:
		return errors.New(string(models.UnexpectedTokenErrLabel))
	}
	return nil
}

func ParseSingleInstruction(parser *models.Parser) error {
	pos := parser.Where()
	h0 := parser.Get(0)
	if h0 == nil {
		return errors.New(string(models.ExpectedTokenErrLabel))
	}
	return nil
}

func ParseDoubleInstruction(parser *models.Parser) error {
	return nil
}
