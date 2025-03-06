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
	h0, h1 := parser.Get(0), parser.Get(2)
	if h0 == nil || h1 == nil {
		return errors.New(string(models.ExpectedTokenErrLabel))
	}
	switch h0.Kind {
	case models.TOKEN_GET:
		if h1.Kind != models.TOKEN_MEM && h1.Kind != models.TOKEN_NUMBER {
			return errors.New(string(models.UnexpectedTokenErrLabel))
		}
		break
	case models.TOKEN_SET, models.TOKEN_ADD, models.TOKEN_MUL,
		models.TOKEN_AND, models.TOKEN_ORR, models.TOKEN_XOR,
		models.TOKEN_CMP, models.TOKEN_JMP, models.TOKEN_JIZ, models.TOKEN_JNZ:
		if h1.Kind != models.TOKEN_MEM {
			return errors.New(string(models.UnexpectedTokenErrLabel))
		}
		break
	default:
		return errors.New(string(models.UnexpectedTokenErrLabel))
	}
	parser.Inject(models.NewSingleInstructionStmt(h0.Kind, pos, *h1))
	parser.Consume(2)
	return nil
}

// ParseDoubleInstruction parses: Instruction space memAddress space comma space (memAddress|number)
func ParseDoubleInstruction(parser *models.Parser) error {
	pos := parser.Where()
	h0 := parser.Get(0)
	if h0 == nil {
		return errors.New(string(models.ExpectedTokenErrLabel))
	}
	if h0.Kind != models.TOKEN_CPY {
		return errors.New(string(models.UnexpectedTokenErrLabel))
	}
	parser.Consume(1)
	h1 := parser.HasNextConsume(models.OptionalSpaceMode, models.TOKEN_MEM)
	if h1 == nil {
		return errors.New(string(models.ExpectedTokenErrLabel))
	}
	if k := parser.HasNextConsume(models.OptionalSpaceMode, models.TOKEN_COMMA); k == nil {
		return errors.New(string(models.ExpectedTokenErrLabel))
	}
	h2 := parser.HasNextConsume(models.OptionalSpaceMode, models.TOKEN_MEM, models.TOKEN_NUMBER)
	if h2 == nil {
		return errors.New(string(models.ExpectedTokenErrLabel))
	}

	parser.Inject(models.NewDoubleInstructionStmt(h0.Kind, pos, *h1, *h2))
	return nil
}
